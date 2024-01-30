package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/johnkristanf/VoiceForge/server/types"
	"github.com/johnkristanf/VoiceForge/server/utils"
)


func (s *ApiServer) FetchAndInsertVoicesInDBHandler(res http.ResponseWriter, req *http.Request) error {

	if err := utils.HttpMethod("GET", req); err != nil{
		return err
	}

	url := "https://api.play.ht/api/v2/voices"

	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil{
		return err
	}

	httpReq.Header.Set("accept", "application/json")
	httpReq.Header.Set("AUTHORIZATION", "dc23bdb0088e43d0ae92155f682d658b")
	httpReq.Header.Set("X-USER-ID", "zXUVGgbbxFM42MjWQG3foHTHnLT2")

	client := &http.Client{}
    httpRes, err := client.Do(httpReq)
    if err != nil {
        return err
    }

	defer httpRes.Body.Close()

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}

	var voicesArr []types.VoiceStruct

	if err := json.Unmarshal(body, &voicesArr); err != nil{
		return err
	}

	for _, voices := range voicesArr {
		if err := s.database.InsertVoice(&voices); err != nil{
			return err
		}
	}
	

	return nil
}


func (s *ApiServer) FetchVoicesHandler(res http.ResponseWriter, req *http.Request) error {
	startTime := time.Now()

	if err := utils.HttpMethod(http.MethodGet, req); err != nil{
		return err
	}

	search_voice := mux.Vars(req)["search_voice"] 
	var voiceCached []*types.VoiceStruct

	if cacheErr := s.client.CacheGet(search_voice, &voiceCached); cacheErr == nil{
		executionTime := time.Since(startTime)

	    resMap := map[string]any{
		    "voices": voiceCached, 
		    "executionTime cached": executionTime.String(),
	    }

	    utils.WriteJson(res, http.StatusOK, resMap)
		return nil
	}

	voices, err := s.database.Voices(search_voice) 
	if err != nil{
		return err
	}
	

	if err := s.client.CacheSet(voices, search_voice); err != nil{
		return err
	}

	executionTime := time.Since(startTime)
	resMap := map[string]any{
		"voices": voices, 
		"executionTime uncached": executionTime.String(),
	}

	utils.WriteJson(res, http.StatusOK, resMap)

	return nil
}


func (s *ApiServer) VoiceCloneHandler(res http.ResponseWriter, req *http.Request) error{

	if err := utils.HttpMethod(http.MethodPost, req); err != nil{
		return err
	}

	errorChan := make(chan error, 1)
	respChan := make(chan []byte, 1)

	err := req.ParseMultipartForm(50 << 20) 
    if err != nil {
        return err
    }

    voice_name := req.FormValue("voice_name")
	sample_file, _, err := req.FormFile("sample_file")
	if err != nil {
		return err
	}
	defer sample_file.Close()



	wg.Add(1)

	    go func() {
		    defer wg.Done()

		    resp, err := s.VoiceCloneRequest(voice_name, sample_file); 
			if err != nil{
			    errorChan <- err
		    }

			fmt.Println("Response from VoiceCloneRequest:", string(resp))
			respChan <- resp
	    }()


	wg.Wait()  
	close(errorChan)
	close(respChan)

	if err := <- errorChan; err != nil{
		return err
	}

	return nil
}

func (s *ApiServer) VoiceCloneRequest(voice_name string, sample_file multipart.File) ([]byte, error) {

	var requestBody bytes.Buffer
	formWriter := multipart.NewWriter(&requestBody)

	fmt.Println("Voice Name sa request", voice_name)
	fmt.Println("Voice File  sa request", sample_file)


	if err := formWriter.WriteField("voice_name", voice_name); err != nil {
		return nil, err
	}
	

	fileWriter, err := formWriter.CreateFormFile("sample_file", "sample_file")
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, sample_file)
	if err != nil {
		return nil, err
	}

	formWriter.Close()

	contentType := formWriter.FormDataContentType()
    fmt.Println("Content-Type:", contentType)
	
	url := "https://api.play.ht/api/v2/cloned-voices/instant"
	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return nil, err
	}


	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "multipart/form-data; boundary=---011000010111000001101001")
	req.Header.Set("AUTHORIZATION", "e1f2dd6ceaa54658a0741be57e927cb6")
	req.Header.Set("X-USER-ID", "5zqbxykOY0byMItNgL7YEjPsTNz1")


	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()


	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
