package handlers

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"

	"io"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/johnkristanf/VoiceForge/server/types"
	"github.com/johnkristanf/VoiceForge/server/utils"
)

var wg sync.WaitGroup

func (s *ApiServer) StreamAudioHandler(res http.ResponseWriter, req *http.Request) error {

	if err := utils.HttpMethod(http.MethodPost, req); err != nil {
		return err
	}

	var StreamBody types.StreamBody
	streamBufferChan := make(chan []byte, 1)
	idChan := make(chan int64, 1)
	errChan := make(chan error, 2)

	if err := json.NewDecoder(req.Body).Decode(&StreamBody); err != nil {
		return err
	}

	go func() {
		defer close(streamBufferChan)

		streamData, streamErr := s.sendStreamRequestAndWriteBuffer(StreamBody)
		if streamErr != nil {
			errChan <- streamErr
		}

		streamBufferChan <- streamData

	}()

	select {

	    case streamBuffer := <-streamBufferChan:

			wg.Add(1)

		        go func() {
			        defer wg.Done()

			        insertedID, err := s.database.InsertAudioStream(StreamBody.Text, streamBuffer); 
			        if err != nil{
				        errChan <- err
			        }

			        idChan <- insertedID

		        }()

			wg.Wait()
			close(errChan)
			close(idChan)

			if err := s.client.CacheDelete("audio"); err != nil {
				return err
			}

			lastInsertedID := <- idChan
	
			res.Header().Set("Content-Type", "text/plain")
			res.Write([]byte(fmt.Sprint(lastInsertedID)))


		case err := <- errChan:
			return err
			
	}

	return nil

}

func (s *ApiServer) compressData(data []byte) ([]byte, error) {

	var buff bytes.Buffer

	gz := gzip.NewWriter(&buff)
	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	}

	if err := gz.Close(); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func (s *ApiServer) sendStreamRequestAndWriteBuffer(streamBody types.StreamBody) ([]byte, error) {

	jsonBody, jsonErr := json.Marshal(streamBody)
	if jsonErr != nil {
		return nil, jsonErr
	}

	req, reqErr := http.NewRequest("POST", "https://api.play.ht/api/v2/tts/stream", bytes.NewBuffer(jsonBody))
	if reqErr != nil {
		return nil, reqErr
	}

	req.Header.Set("Accept", "audio/mpeg")
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("AUTHORIZATION", os.Getenv("AUTHORIZATION_API_KEY"))
	req.Header.Set("X-USER-ID", os.Getenv("USER_API_KEY"))

	client := &http.Client{}
	resp, resErr := client.Do(req)
	if resErr != nil {
		return nil, resErr
	}

	defer resp.Body.Close()

	readData, copyErr := io.ReadAll(resp.Body)
	if copyErr != nil {
		return nil, copyErr
	}

	compressedData, err := s.compressData(readData)
	if err != nil {
		return nil, err
	}

	return compressedData, nil

}

func (s *ApiServer) FetchAudioDataHandler(res http.ResponseWriter, req *http.Request) error {

	if err := utils.HttpMethod(http.MethodGet, req); err != nil {
		return err
	}

	var audioCached []*types.AudioStruct

	if cacheErr := s.client.CacheGet("audio", &audioCached); cacheErr == nil {
		return utils.WriteJson(res, http.StatusOK, audioCached)
	}

	audio, err := s.database.FetchAudioStream()
	if err != nil {
		return err
	}

	if err := s.client.CacheSet(audio, "audio"); err != nil {
		return err
	}

	return utils.WriteJson(res, http.StatusOK, audio)

}

func (s *ApiServer) DeleteAudioDataHandler(res http.ResponseWriter, req *http.Request) error {

	if err := utils.HttpMethod(http.MethodDelete, req); err != nil {
		return err
	}

	errorChan := make(chan error, 1)
	idChan := make(chan int64, 1)

	audio_idStr := mux.Vars(req)["audio_id"]
	audio_id, err := strconv.ParseInt(audio_idStr, 10, 0)
	if err != nil{
		return err
	}

	go func ()  {
		defer close(errorChan)
		defer close(idChan)

		deletedID, err := s.database.DeleteAudioData(audio_id)
	    if err != nil {
		    errorChan <- err
	    }

		idChan <- deletedID

	}()


	select{

	    case deletedID := <- idChan:

			if err := s.client.CacheDelete("audio"); err != nil {
				return err
			}
		
			res.Header().Set("Content-Type", "text/plain")
			res.Write([]byte(fmt.Sprint(deletedID)))
		
			return nil
	}
}
