package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/johnkristanf/VoiceForge/server/config"
	"github.com/johnkristanf/VoiceForge/server/database"
	"github.com/rs/cors"
)

type APIError struct{
	ERROR string
}

type ApiServer struct{
	listenAddr string
	database   database.Method
	cors *cors.Cors
	client config.RedisMethod
}


type APIFunction func (res http.ResponseWriter, req *http.Request) error

func makeHTTPHandlerFunc(handlerFunc APIFunction) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		if err := handlerFunc(res, req); err != nil{
			log.Fatalln("HANDLER ERROR:", &APIError{ERROR: err.Error()})
		}
	}
}


func NewAPIServer(listenAddr string, db database.Method, cors *cors.Cors, client config.RedisMethod) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
		database: db,
		cors: cors,
		client: client,
	}
}


func (s *ApiServer) Run() error {

	router := mux.NewRouter()

	requestHandler := s.cors.Handler(router) 

	// GET HANDLER
	router.HandleFunc("/api/audio/data", makeHTTPHandlerFunc(s.FetchAudioDataHandler))
	router.HandleFunc("/api/voices", makeHTTPHandlerFunc(s.FetchVoicesHandler))

	// POST HANDLER
	router.HandleFunc("/api/stream/voices", makeHTTPHandlerFunc(s.StreamAudioHandler))
	router.HandleFunc("/api/voice/clone", makeHTTPHandlerFunc(s.VoiceCloneHandler))
	router.HandleFunc("/auth/signup", makeHTTPHandlerFunc(s.SignUpHandler))

	// DELETE HANDLER
	router.HandleFunc("/api/audio/delete/{audio_id}", makeHTTPHandlerFunc(s.DeleteAudioDataHandler))



	if err := http.ListenAndServe(s.listenAddr, requestHandler); err != nil{
		return err
	}

	log.Println("Server Running on Port", s.listenAddr)

	return nil
}

