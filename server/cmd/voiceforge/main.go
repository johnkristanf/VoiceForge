package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"

	"server/internal/routes"
)

func main(){

	router := gin.Default()

	routes.Voices(router)

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{})})

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

	handler := cors.Handler(router)

	if err := http.ListenAndServe(":800", handler); err != nil{
		panic(err)
	}
	
}