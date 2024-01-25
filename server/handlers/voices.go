package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
	"fmt"

	"github.com/johnkristanf/VoiceForge/server/types"
	"github.com/johnkristanf/VoiceForge/server/utils"
)


func (s *ApiServer) FetchAndInsertVoicesInDB(res http.ResponseWriter, req *http.Request) error {

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


func (s *ApiServer) FetchVoices(res http.ResponseWriter, req *http.Request) error {
	startTime := time.Now()

	if err := utils.HttpMethod(http.MethodGet, req); err != nil{
		return err
	}

	var voiceCached []*types.VoiceStruct

	if cacheErr := s.client.CacheGet("voices", &voiceCached); cacheErr == nil{
		executionTime := time.Since(startTime)

	    resMap := map[string]any{
		    "voices": voiceCached, 
		    "executionTime cached": executionTime.String(),
	    }

	    utils.WriteJson(res, http.StatusOK, resMap)
		return nil
	}

	voices, err := s.database.Voices() 
	if err != nil{
		return err
	}
	

	if err := s.client.CacheSet(voices, "voices"); err != nil{
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


func (s *ApiServer) VoiceClone(res http.ResponseWriter, req *http.Request) error{

	if err := utils.HttpMethod(http.MethodPost, req); err != nil{
		return err
	}

	var voiceCloneData *types.VoiceCloneTypes
	errorChan := make(chan error, 1)

	body, err := io.ReadAll(req.Body)
	if err != nil{
		return err
	}

	if err := json.Unmarshal(body, &voiceCloneData); err != nil{
		return err
	}


	wg.Add(1)

	    go func() {
		  defer wg.Done()

		    if err := s.VoiceCloneRequest(voiceCloneData); err != nil{
			    errorChan <- err
		    }

	    }()

	wg.Wait()  
	close(errorChan)

	if err := <- errorChan; err != nil{
		return err
	}

	return nil
}

func (s *ApiServer) VoiceCloneRequest(voiceCloneData *types.VoiceCloneTypes) error {

	jsonBody, jsonErr := json.Marshal(voiceCloneData)
	if jsonErr != nil{
		return jsonErr
	}


	url := "https://api.play.ht/api/v2/cloned-voices/instant"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))

	boundary := "---011000010111000001101001"

    headers := map[string]string{
        "Accept":         "audio/mpeg",
        "Content-Type":   "multipart/form-data; boundary=" + boundary,
        "AUTHORIZATION":  "e1f2dd6ceaa54658a0741be57e927cb6",
        "X-USER-ID":      "5zqbxykOY0byMItNgL7YEjPsTNz1",
    }

	for key, value := range headers{
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		return err
	}

	defer resp.Body.Close()
	readBody, err := io.ReadAll(resp.Body)
	if err != nil{
		return err
	}

	fmt.Println(string(readBody))

	return nil

}