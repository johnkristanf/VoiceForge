package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/johnkristanf/VoiceForge/server/auth"
	"github.com/johnkristanf/VoiceForge/server/config"
	"github.com/johnkristanf/VoiceForge/server/database"
)

type APIError struct{
	ERROR string
}

type ApiServer struct{
	listenAddr string
	database   database.Method
	client config.RedisMethod
	smtpClient auth.SmtpClientMethod
}


type APIFunction func (res http.ResponseWriter, req *http.Request) error


func makeHTTPHandlerFunc(handlerFunc APIFunction) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		if err := handlerFunc(res, req); err != nil{
			log.Fatalln("HANDLER ERROR:", &APIError{ERROR: err.Error()})
		}
	}
}


func NewAPIServer(listenAddr string, db database.Method, client config.RedisMethod, smtpClient auth.SmtpClientMethod) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
		database: db,
		client: client,
		smtpClient: smtpClient,
	}
}


func (s *ApiServer) Run() error {

	router := mux.NewRouter()

	// GET HANDLER
	router.HandleFunc("/api/audio/data", auth.AuthenticationMiddleWare(makeHTTPHandlerFunc(s.FetchAudioDataHandler)))
	router.HandleFunc("/api/voices/{search_voice}", auth.AuthenticationMiddleWare(makeHTTPHandlerFunc(s.FetchVoicesHandler)))
	router.HandleFunc("/api/get/voice/clone", auth.AuthenticationMiddleWare(makeHTTPHandlerFunc(s.FetchVoiceClone)))


	router.HandleFunc("/user/data", auth.AuthenticationMiddleWare(makeHTTPHandlerFunc(s.FetchUserDatahandler)))


	// POST HANDLER
	router.HandleFunc("/api/stream/voices", auth.AuthenticationMiddleWare(makeHTTPHandlerFunc(s.StreamAudioHandler)))
	router.HandleFunc("/api/voice/clone", auth.AuthenticationMiddleWare(makeHTTPHandlerFunc(s.VoiceCloneHandler)))
	router.HandleFunc("/voice/clone/delete", auth.AuthenticationMiddleWare(makeHTTPHandlerFunc(s.DeleteVoiceClone)))
	
	router.HandleFunc("/auth/signup", makeHTTPHandlerFunc(s.SignUpHandler))
	router.HandleFunc("/auth/login", makeHTTPHandlerFunc(s.LoginHandler))
	router.HandleFunc("/auth/verification", makeHTTPHandlerFunc(s.VerifyUserHandler))

	router.HandleFunc("/logout", auth.AuthenticationMiddleWare(makeHTTPHandlerFunc(s.LogoutHandler)))



	// DELETE HANDLER
	router.HandleFunc("/api/audio/delete/{audio_id}", auth.AuthenticationMiddleWare(makeHTTPHandlerFunc(s.DeleteAudioDataHandler)))

	// TOKEN HANDLERS
	router.HandleFunc("/token/refresh", s.RefreshTokenHandler)



	if err := http.ListenAndServe(s.listenAddr, router); err != nil{
		return err
	}

	return nil
}

