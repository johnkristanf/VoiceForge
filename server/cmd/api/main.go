package main

import (
	"log"
	"net/http"

	"github.com/johnkristanf/VoiceForge/server/auth"
	"github.com/johnkristanf/VoiceForge/server/config"
	"github.com/johnkristanf/VoiceForge/server/database"
	"github.com/johnkristanf/VoiceForge/server/handlers"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)




func main() {


	if err := godotenv.Load("../../.env"); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
    

	db, err := database.VoiceForgeDB()
	if err != nil {
		log.Fatalln("DATABASE ERROR", err.Error())
	}

	if err := db.DBInit(); err != nil {
		log.Fatalln("DATABASE TABLES ERROR", err.Error())
	}


	client, err := config.RedisConfig()
	if err != nil{
		log.Fatalln("REDIS CONFIG ERROR", err)
	}


	smtpClient, err := auth.NewSmtpClient()
	if err != nil{
		log.Fatalln("SMTP CONFIG ERROR", err)
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

	
	server := handlers.NewAPIServer(":800", db, cors, client, smtpClient)

	if err := server.Run(); err != nil {
		log.Fatalln("SERVER ERROR", err)
	}
}
