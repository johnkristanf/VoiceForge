package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
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
	httpReq.Header.Set("AUTHORIZATION", os.Getenv("AUTHORIZATION_API_KEY"))
	httpReq.Header.Set("X-USER-ID", os.Getenv("USER_API_KEY"))

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

	var voiceClone *types.VoiceCloneType
	errorChan := make(chan error, 1)
	respChan := make(chan []byte, 1)

	err := req.ParseMultipartForm(10 << 20) 
    if err != nil {
        return err
    }


    voice_name := req.FormValue("voice_name")
	sample_file, _, err := req.FormFile("sample_file")
	if err != nil {
		return err
	}
	defer sample_file.Close()


	go func() {
		defer close(errorChan)
		defer close(respChan)

		resp, err := s.VoiceCloneRequest(voice_name, sample_file); 
		if err != nil{
			errorChan <- err
		}

		respChan <- resp
	}()

	select {

	    case err := <-errorChan:
		    return err

	    case responseBody := <-respChan:
		
		    if err := s.client.CacheSet(responseBody, "voiceCloneData"); err != nil {
			    return err
		    }

		    if err := json.Unmarshal(responseBody, &voiceClone); err != nil{
			    return err
		    }

		    return utils.WriteJson(res, http.StatusOK, voiceClone)
	}
}


func (s *ApiServer) VoiceCloneRequest(voice_name string, sample_file multipart.File) ([]byte, error) {

	var requestBody bytes.Buffer
	formWriter := multipart.NewWriter(&requestBody)

	if err := formWriter.SetBoundary("---011000010111000001101001"); err != nil{
		return nil, err
	}

	if err := formWriter.WriteField("voice_name", voice_name); err != nil {
		return nil, err
	}

	fileWriter, err := formWriter.CreatePart(textproto.MIMEHeader{
		"Content-Disposition":   []string{`form-data; name="sample_file"; filename="sample_file.mp3"`},
		"Content-Type":          []string{"audio/mpeg"},
	})

	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, sample_file)
	if err != nil {
		return nil, err
	}
	
	formWriter.Close()
	
	url := "https://api.play.ht/api/v2/cloned-voices/instant"
	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", formWriter.FormDataContentType())
	req.Header.Set("AUTHORIZATION", os.Getenv("AUTHORIZATION_API_KEY"))
	req.Header.Set("X-USER-ID", os.Getenv("USER_API_KEY"))


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


func (s *ApiServer) FetchVoiceClone(res http.ResponseWriter, req *http.Request) error {

	if err := utils.HttpMethod(http.MethodGet, req); err != nil{
		return err
	}

	var voiceClone []*types.VoiceCloneType

	if cacheErr := s.client.CacheGet("voiceCloneData", &voiceClone); cacheErr == nil {
		return utils.WriteJson(res, http.StatusOK, voiceClone)
	}

	url := "https://api.play.ht/api/v2/cloned-voices"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("AUTHORIZATION", os.Getenv("AUTHORIZATION_API_KEY"))
	req.Header.Add("X-USER-ID", os.Getenv("USER_API_KEY"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		return err
	}
	defer resp.Body.Close()


	if err := json.NewDecoder(resp.Body).Decode(&voiceClone); err != nil{
		return err
	}

	if err := s.client.CacheSet(voiceClone, "voiceCloneData"); err != nil{
		return err
	}

	return utils.WriteJson(res, http.StatusOK, voiceClone)
}


type VoiceID struct{
	VoiceID string `json:"voice_id"`
}

func (s *ApiServer) DeleteVoiceClone(res http.ResponseWriter, req *http.Request) error {

	if err := utils.HttpMethod(http.MethodPost, req); err != nil{
		return err
	}

	var clone *VoiceID

	body, err := io.ReadAll(req.Body)
	if err != nil{
		return err
	}

	if err := json.Unmarshal(body, &clone); err != nil{
		return err
	}


	jsonBody, err := json.Marshal(clone);
	if err != nil{
		return err
	}

	url := "https://api.play.ht/api/v2/cloned-voices/"

	delrequest, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonBody))
	if err != nil{
		return  err
	}
	
	delrequest.Header.Add("Accept", "application/json")
	delrequest.Header.Add("Content-Type", "application/json")
	delrequest.Header.Add("AUTHORIZATION", os.Getenv("AUTHORIZATION_API_KEY"))
	delrequest.Header.Add("X-USER-ID", os.Getenv("USER_API_KEY"))

	
	resp, err := http.DefaultClient.Do(delrequest)
	if err != nil{
		return err
	}
	defer resp.Body.Close()


	readBody, err := io.ReadAll(resp.Body)
	if err != nil{
		return  err
	}

	if err := s.client.CacheDelete("voiceCloneData"); err != nil{
		return err
	}

	return utils.WriteJson(res, http.StatusOK, string(readBody))
}


