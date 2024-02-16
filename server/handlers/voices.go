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

	if err := utils.HttpMethod("GET", req); err != nil {
		return err
	}

	val, err := s.database.CheckVoicesValues()
	if err != nil {
		return err
	}

	if val > 0 {
		return utils.WriteJson(res, http.StatusBadRequest, map[string]string{
			"ERROR": "Unable to perform Action",
		})
	}

	url := "https://api.play.ht/api/v2/voices"

	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
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

	if err := json.Unmarshal(body, &voicesArr); err != nil {
		return err
	}

	for _, voices := range voicesArr {
		if err := s.database.InsertVoice(&voices); err != nil {
			return err
		}
	}

	return nil
}

func (s *ApiServer) FetchVoicesHandler(res http.ResponseWriter, req *http.Request) error {
	startTime := time.Now()

	if err := utils.HttpMethod(http.MethodGet, req); err != nil {
		return err
	}


	search_voice := mux.Vars(req)["search_voice"]
	var voiceCached []*types.FetchVoiceTypes

	if cacheErr := s.client.CacheGet(search_voice, &voiceCached); cacheErr == nil {
		executionTime := time.Since(startTime)

		resMap := map[string]any{
			"voices":               voiceCached,
			"executionTime cached": executionTime.String(),
		}

		return utils.WriteJson(res, http.StatusOK, resMap)
	}

	
	voices, err := s.database.Voices(search_voice)
	if err != nil {
		return err
	}


	if err := s.client.CacheSet(voices, search_voice); err != nil {
		return err
	}

	executionTime := time.Since(startTime)
	resMap := map[string]any{
		"voices":                 voices,
		"executionTime uncached": executionTime.String(),
	}

	return utils.WriteJson(res, http.StatusOK, resMap)
	

}

func (s *ApiServer) VoiceCloneHandler(res http.ResponseWriter, req *http.Request) error {

	if err := utils.HttpMethod(http.MethodPost, req); err != nil {
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

		resp, err := s.VoiceCloneRequest(voice_name, sample_file)
		if err != nil {
			errorChan <- err
		}

		respChan <- resp
	}()

	select {

	case err := <-errorChan:
		return err

	case responseBody := <-respChan:

		if err := s.client.CacheDelete("voiceCloneData"); err != nil {
			return err
		}

		if err := json.Unmarshal(responseBody, &voiceClone); err != nil {
			return err
		}

		return utils.WriteJson(res, http.StatusOK, voiceClone)
	}
}

func (s *ApiServer) VoiceCloneRequest(voice_name string, sample_file multipart.File) ([]byte, error) {

	var requestBody bytes.Buffer
	formWriter := multipart.NewWriter(&requestBody)

	if err := formWriter.SetBoundary("---011000010111000001101001"); err != nil {
		return nil, err
	}

	if err := formWriter.WriteField("voice_name", voice_name); err != nil {
		return nil, err
	}

	fileWriter, err := formWriter.CreatePart(textproto.MIMEHeader{
		"Content-Disposition": []string{`form-data; name="sample_file"; filename="sample_file.mp3"`},
		"Content-Type":        []string{"audio/mpeg"},
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

	if err := utils.HttpMethod(http.MethodGet, req); err != nil {
		return err
	}

	var voiceClone []*types.VoiceCloneType
	resBodyChan := make(chan *bytes.Buffer, 1)
	errorChan := make(chan error, 1)

	if cacheErr := s.client.CacheGet("voiceCloneData", &voiceClone); cacheErr == nil {
		return utils.WriteJson(res, http.StatusOK, voiceClone)
	}

	url := "https://api.play.ht/api/v2/cloned-voices"

	headers := make(map[string]string, 3)
	headers["Accept"] = "application/json"
	headers["AUTHORIZATION"] = os.Getenv("AUTHORIZATION_API_KEY")
	headers["X-USER-ID"] = os.Getenv("USER_API_KEY")

	jsonBody, jsonErr := json.Marshal(voiceClone)
	if jsonErr != nil {
		return jsonErr
	}

	go func() {
		defer close(errorChan)
		defer close(resBodyChan)

		resBody, err := s.ExternalApiRequest("GET", url, bytes.NewBuffer(jsonBody), headers)
		if err != nil {
			errorChan <- err
		}

		resBodyChan <- resBody
	}()

	select {

	    case err := <-errorChan:
		    return err

	    case responseBody := <-resBodyChan:

		    if err := json.NewDecoder(responseBody).Decode(&voiceClone); err != nil {
			    return err
		    }

		    if err := s.client.CacheSet(voiceClone, "voiceCloneData"); err != nil {
			    return err
		    }

		    return utils.WriteJson(res, http.StatusOK, voiceClone)
	}

}

type VoiceID struct {
	VoiceID string `json:"voice_id"`
}

func (s *ApiServer) DeleteVoiceClone(res http.ResponseWriter, req *http.Request) error {

	if err := utils.HttpMethod(http.MethodPost, req); err != nil {
		return err
	}

	var clone *VoiceID
	resBodyChan := make(chan *bytes.Buffer, 1)
	errorChan := make(chan error, 1)

	if err := json.NewDecoder(req.Body).Decode(&clone); err != nil {
		return err
	}

	jsonBody, err := json.Marshal(clone)
	if err != nil {
		return err
	}

	url := "https://api.play.ht/api/v2/cloned-voices/"

	headers := make(map[string]string, 4)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["AUTHORIZATION"] = os.Getenv("AUTHORIZATION_API_KEY")
	headers["X-USER-ID"] = os.Getenv("USER_API_KEY")

	go func() {
		defer close(errorChan)
		defer close(resBodyChan)

		resBody, err := s.ExternalApiRequest("DELETE", url, bytes.NewBuffer(jsonBody), headers)
		if err != nil {
			errorChan <- err
		}

		resBodyChan <- resBody
	}()

	select {

	case err := <-errorChan:
		return err

	case responseBody := <-resBodyChan:

		if err := s.client.CacheDelete("voiceCloneData"); err != nil {
			return err
		}

		res.Header().Set("Content-Type", "text/plain")
		res.Write([]byte(responseBody.String()))
	}

	return nil

}

func (s *ApiServer) ExternalApiRequest(method string, url string, reqBody *bytes.Buffer, headers map[string]string) (*bytes.Buffer, error) {

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var bodyBuffer bytes.Buffer
	
	_, err = io.Copy(&bodyBuffer, res.Body)
	if err != nil{
		return nil, err
	}

	return &bodyBuffer, nil
}
