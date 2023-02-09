package rest

import (
	"github.com/gin-gonic/gin"
	"vita-message-service/delivery/rest/handler"
)

func LoadRoutes(mh handler.MessageHandler) {
	router := gin.Default()
	router.POST("/message", mh.SendMessage)
	router.GET("/message/:email", mh.GetMessage)
}
