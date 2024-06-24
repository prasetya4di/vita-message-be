package rest

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
	"vita-message-service/delivery/rest/handler"
	"vita-message-service/delivery/rest/middlewares"
)

func LoadRoutes(mh handler.MessageHandler, ih handler.ImageHandler, ah handler.AuthHandler) {
	gin.SetMode(os.Getenv(gin.EnvGinMode))
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://vita-377401.web.app", "http://localhost:5555", "https://*.ngrok.io"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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

	err := router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatal(err)
	}
}
