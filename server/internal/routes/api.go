package routes

import (
	"github.com/gin-gonic/gin"
	"server/internal/handler"
)

func Voices(router *gin.Engine){

	apiRoute := router.Group("/api")
	apiRoute.GET("/voices", handler.FetchVoices)
	apiRoute.POST("/stream/voices", handler.StreamAudio)

}