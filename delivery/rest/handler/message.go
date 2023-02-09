package handler

import "github.com/gin-gonic/gin"

type MessageHandler interface {
	SendMessage(c *gin.Context)
	GetMessage(c *gin.Context)
}
