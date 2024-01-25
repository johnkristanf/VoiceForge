package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/johnkristanf/VoiceForge/server/config"
	"github.com/johnkristanf/VoiceForge/server/database"
	"github.com/johnkristanf/VoiceForge/server/handlers"
	"github.com/rs/cors"
)

func main() {

	db, err := database.VoiceForgeDB()
	if err != nil {
		log.Fatalln("DATABASE ERROR", err.Error())
	}

	if err := db.DBInit(); err != nil {
		log.Fatalln("DATABASE TABLES ERROR", err.Error())
	}

	fmt.Printf("%v+\n", db)


	client, err := config.RedisConfig()
	if err != nil{
		log.Fatalln("REDIS CONFIG ERROR", err)
	}


	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:500"},

		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},

		AllowCredentials: true,

		AllowedHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Origin",
			"Cookie",
		},
	})

	
	server := handlers.NewAPIServer(":800", db, cors, client)

	if err := server.Run(); err != nil {
		log.Fatalln("SERVER ERROR", err)
	}
}
