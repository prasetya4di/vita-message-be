package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"vita-message-service/delivery/rest/handler"
)

func LoadRoutes(mh handler.MessageHandler, ih handler.ImageHandler, ah handler.AuthHandler) {
	router := gin.Default()
	router.POST("/login", ah.Login)
	router.POST("/register", ah.Register)

	router.POST("/message", mh.SendMessage)
	router.POST("/message/reply", mh.ReplyMessage)
	router.GET("/message/:email", mh.GetMessage)

	router.POST("/image/:email", ih.UploadImage)
	router.Static("/image", "upload/images")
	err := router.Run(os.Getenv("BASEURL") + ":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
