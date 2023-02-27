package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"vita-message-service/delivery/rest/handler"
	"vita-message-service/delivery/rest/middlewares"
)

func LoadRoutes(mh handler.MessageHandler, ih handler.ImageHandler, ah handler.AuthHandler) {
	router := gin.Default()
	router.POST("/login", ah.Login)
	router.POST("/register", ah.Register)
	router.Static("/image", "upload/images")

	routeMessage := router.Group("/message")
	routeMessage.Use(middlewares.JwtAuthMiddleware())
	routeMessage.POST("", mh.SendMessage)
	routeMessage.POST("/reply", mh.ReplyMessage)
	routeMessage.GET("", mh.GetMessage)

	routeImage := router.Group("/image")
	routeImage.Use(middlewares.JwtAuthMiddleware())
	routeImage.POST("", ih.UploadImage)

	err := router.Run(os.Getenv("BASEURL") + ":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
