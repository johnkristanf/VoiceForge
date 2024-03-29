package handlers

import (
	"log"
	"net/http"
	

	"github.com/gorilla/mux"
	"github.com/johnkristanf/VoiceForge/server/auth"
	"github.com/johnkristanf/VoiceForge/server/config"
	"github.com/johnkristanf/VoiceForge/server/database"
	"github.com/rs/cors"
)

type APIError struct {
	ERROR string
}

type ApiServer struct {
	listenAddr string
	database   database.Method
	cors       *cors.Cors
	client     config.RedisMethod
	smtpClient auth.SmtpClientMethod
	ERROR error
}

type APIFunction func(res http.ResponseWriter, req *http.Request) error

func (s *ApiServer) makeHTTPHandlerFunc(handlerFunc APIFunction) http.HandlerFunc {

    return func(res http.ResponseWriter, req *http.Request) {
       
        if err := handlerFunc(res, req); err != nil {
            log.Println("HANDLER ERROR:", &APIError{ERROR: err.Error()})
			s.ERROR = err
        }
    }
}


func NewAPIServer(listenAddr string, db database.Method, cors *cors.Cors, client config.RedisMethod, smtpClient auth.SmtpClientMethod) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
		database:   db,
		cors:       cors,
		client:     client,
		smtpClient: smtpClient,
	}

}

func (s *ApiServer) Run() error {

	router := mux.NewRouter()

	requestHandler := s.cors.Handler(router)

	// GET HANDLER
	router.HandleFunc("/api/audio/data", auth.AuthenticationMiddleWare(s.makeHTTPHandlerFunc(s.FetchAudioDataHandler)))
	router.HandleFunc("/api/voices/{search_voice}", auth.AuthenticationMiddleWare(s.makeHTTPHandlerFunc(s.FetchVoicesHandler)))
	router.HandleFunc("/api/get/voice/clone", auth.AuthenticationMiddleWare(s.makeHTTPHandlerFunc(s.FetchVoiceClone)))
	router.HandleFunc("/api/insert/voice", auth.AuthenticationMiddleWare(s.makeHTTPHandlerFunc(s.FetchAndInsertVoicesInDBHandler)))

	router.HandleFunc("/user/data", auth.AuthenticationMiddleWare(s.makeHTTPHandlerFunc(s.FetchUserDatahandler)))

	// POST HANDLER
	router.HandleFunc("/api/stream/voices", auth.AuthenticationMiddleWare(s.makeHTTPHandlerFunc(s.StreamAudioHandler)))
	router.HandleFunc("/api/voice/clone", auth.AuthenticationMiddleWare(s.makeHTTPHandlerFunc(s.VoiceCloneHandler)))
	router.HandleFunc("/voice/clone/delete", auth.AuthenticationMiddleWare(s.makeHTTPHandlerFunc(s.DeleteVoiceClone)))

	router.HandleFunc("/auth/signup", s.makeHTTPHandlerFunc(s.SignUpHandler))
	router.HandleFunc("/auth/login", s.makeHTTPHandlerFunc(s.LoginHandler))
	router.HandleFunc("/auth/verification", s.makeHTTPHandlerFunc(s.VerifyUserHandler))

	router.HandleFunc("/logout", auth.AuthenticationMiddleWare(s.makeHTTPHandlerFunc(s.LogoutHandler)))

	// DELETE HANDLER
	router.HandleFunc("/api/audio/delete/{audio_id}", auth.AuthenticationMiddleWare(s.makeHTTPHandlerFunc(s.DeleteAudioDataHandler)))

	// TOKEN HANDLERS
	router.HandleFunc("/token/refresh", s.RefreshTokenHandler)

	if err := http.ListenAndServe(s.listenAddr, requestHandler); err != nil {
		return err
	}

	return nil
}
