package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/johnkristanf/VoiceForge/server/auth"
	"github.com/johnkristanf/VoiceForge/server/config"
	"github.com/johnkristanf/VoiceForge/server/database"
	"github.com/johnkristanf/VoiceForge/server/handlers"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	fmt.Println("MING GANA NA ANG DOCKER liboga uyy")

	fmt.Println("DATABASE_URL", os.Getenv("DATABASE_URL"))
	fmt.Println("REDIS_URL", os.Getenv("REDIS_URL"))




	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}


	db, err := database.VoiceForgeDB()
	if err != nil {
		log.Fatalln("DATABASE ERROR", err.Error())
	}


	if err := db.DBInit(); err != nil {
		log.Fatalln("DATABASE TABLES ERROR", err.Error())
	}


	client, err := config.RedisConfig(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatalln("REDIS CONFIG ERROR", err)
	}


	smtpClient, err := auth.NewSmtpClient()
	if err != nil {
		log.Fatalln("SMTP CONFIG ERROR", err)
	}

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"https://voice-forge-client.vercel.app"},

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

	server := handlers.NewAPIServer(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), db, cors, client, smtpClient)

	fmt.Println("port", os.Getenv("SERVER_PORT"))
	
	if err := server.Run(); err != nil {
		log.Fatalln("SERVER ERROR", err)
	}
}
