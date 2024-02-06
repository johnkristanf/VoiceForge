package main

import (
	"fmt"
	"log"
	"os"

	"github.com/johnkristanf/VoiceForge/server/auth"
	"github.com/johnkristanf/VoiceForge/server/config"
	"github.com/johnkristanf/VoiceForge/server/database"
	"github.com/johnkristanf/VoiceForge/server/handlers"
	"github.com/joho/godotenv"
)

func main() {


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


	fmt.Println("bag o dooo babye cors")


	server := handlers.NewAPIServer(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), db, client, smtpClient)

	fmt.Println("port", os.Getenv("SERVER_PORT"))
	
	if err := server.Run(); err != nil {
		log.Fatalln("SERVER ERROR", err)
	}
}
