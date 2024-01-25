package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/johnkristanf/VoiceForge/server/types"
	"github.com/johnkristanf/VoiceForge/server/utils"
)

func (s *ApiServer) SignUpHandler(res http.ResponseWriter, req *http.Request) error {

	startTime := time.Now()

	if err := utils.HttpMethod(http.MethodPost, req); err != nil{
		return err
	}

	var signUpCredentials *types.SignupCredentials

	body, err := io.ReadAll(req.Body)
	if err != nil{
		return err
	}

	if err := json.Unmarshal(body, &signUpCredentials); err != nil{
		return err
	}

	
    if err := s.database.SignUp(signUpCredentials); err != nil {
        return err
	} 
    
	
	utils.WriteJson(res, http.StatusOK, map[string]bool{"Signup": true})

	executionTime := time.Since(startTime)
	fmt.Println("Exec Signup", executionTime.String())
	return nil
	
}