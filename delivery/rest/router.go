package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"vita-message-service/delivery/rest/handler"
)

func LoadRoutes(mh handler.MessageHandler) {
	router := gin.Default()
	router.POST("/message", mh.SendMessage)
	router.GET("/message/:email", mh.GetMessage)
	err := router.Run(os.Getenv("BASEURL") + ":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
