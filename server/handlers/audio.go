package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/johnkristanf/VoiceForge/server/types"
	"github.com/johnkristanf/VoiceForge/server/utils"
)


var wg sync.WaitGroup


func (s *ApiServer) StreamAudioHandler(res http.ResponseWriter, req *http.Request) error {

	if err := utils.HttpMethod(http.MethodPost, req); err != nil{
		return err
	}

	var StreamBody types.StreamBody
	var buffer bytes.Buffer
	errChan := make(chan error, 1)	

	body, err := io.ReadAll(req.Body)
	if err != nil{
		return err
	}

	if err := json.Unmarshal(body, &StreamBody); err != nil{
		return err
	}

	wg.Add(1)

	    go func(){
		    defer wg.Done()

		    if streamErr := s.sendStreamRequestAndWriteBuffer(&buffer, StreamBody); streamErr != nil{
			  errChan <- streamErr
		    }

	    }()
	
	wg.Wait()
	close(errChan)

	if err := <- errChan; err != nil{
		return err
	}

	if err := s.database.InsertAudioStream(StreamBody.Text, &buffer); err != nil{
		return err
	}


	if err := s.client.CacheDelete("audio"); err != nil{
		return err
	}
	
	res.Header().Set("Content-Type", "audio/mpeg")
	res.Write(buffer.Bytes())

	return nil
}

func (s *ApiServer) sendStreamRequestAndWriteBuffer(buffer *bytes.Buffer, streamBody types.StreamBody) error {

	jsonBody, jsonErr := json.Marshal(streamBody)
	if jsonErr != nil{
		return jsonErr
	}

	req, reqErr := http.NewRequest("POST", "https://api.play.ht/api/v2/tts/stream", bytes.NewBuffer(jsonBody))
    if reqErr != nil {
        return reqErr
    }

	headers := map[string]string{
        "Accept":         "audio/mpeg",
        "Content-type":   "application/json",
        "AUTHORIZATION":  "e1f2dd6ceaa54658a0741be57e927cb6",
        "X-USER-ID":      "5zqbxykOY0byMItNgL7YEjPsTNz1",
    }

	for key, value := range headers{
		req.Header.Set(key, value)
	}

	client := &http.Client{}
    resp, resErr := client.Do(req)
    if resErr != nil {
        return resErr
    }

	defer resp.Body.Close()

	_, copyErr := io.Copy(buffer, resp.Body)
	if copyErr != nil{
		return copyErr
	}

	return nil

}


func (s *ApiServer) FetchAudioDataHandler(res http.ResponseWriter, req *http.Request) error {
	startTime := time.Now()

	if err := utils.HttpMethod(http.MethodGet, req); err != nil{
		return err
	}

	var audioCached []*types.AudioStruct

	if cacheErr := s.client.CacheGet("audio", &audioCached); cacheErr == nil {
		executionTime := time.Since(startTime)
		fmt.Println("Execution Time cached", executionTime.String())
		utils.WriteJson(res, http.StatusOK, map[string][]*types.AudioStruct{"audioDataArray": audioCached})
		return nil
	}
	

	audio, err := s.database.FetchAudioStream()
	if err != nil{
		return err
	}

	if err := s.client.CacheSet(audio, "audio"); err != nil{
		return err
	}

	executionTime := time.Since(startTime)
	fmt.Println("Execution Time uncached", executionTime.String())
	utils.WriteJson(res, http.StatusOK, map[string][]*types.AudioStruct{"audioDataArray": audio})
	
	return nil
}


func (s *ApiServer) DeleteAudioDataHandler(res http.ResponseWriter, req *http.Request) error {

	if err := utils.HttpMethod(http.MethodDelete, req); err != nil{
		return err
	}

	audio_idStr := mux.Vars(req)["audio_id"] 
	audio_id, err := strconv.ParseInt(audio_idStr, 10, 0)


	deletedID, err := s.database.DeleteAudioData(audio_id)
	if err != nil{
		return err
	}


	if err := s.client.CacheDelete("audio"); err != nil{
		return err
	}

	utils.WriteJson(res, http.StatusOK, map[string]int64{"DELETED": deletedID})

	return nil
}


