package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"log"
	"mime/multipart"
	"net/http"
	"vita-message-service/delivery/rest/handler"
	"vita-message-service/usecase"
)

type imageHandler struct {
	uploadImage    usecase.UploadImage
	replyMessage   usecase.ReplyMessage
	saveMessage    usecase.SaveMessage
	getCurrentUser usecase.GetCurrentUser
	localizer      *i18n.Localizer
}

func NewImageHandler(uploadImage usecase.UploadImage, message usecase.ReplyMessage, saveMessage usecase.SaveMessage, getCurrentUser usecase.GetCurrentUser, localizer *i18n.Localizer) handler.ImageHandler {
	return &imageHandler{
		uploadImage,
		message,
		saveMessage,
		getCurrentUser,
		localizer,
	}
}

func (ih *imageHandler) UploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	message := c.Request.FormValue("message")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatalf("error insert image 3: %v", err)
		return
	}
	if !isValidImage(file) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid image format"})
		log.Fatalf("error insert image 3: %v", err)
		return
	}
	currentUser, err := ih.getCurrentUser.Invoke(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}
	messages, err := ih.uploadImage.Invoke(currentUser, file, header, message)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		log.Fatalf("error insert image 3: %v", err)
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": messages})
}

func isValidImage(file multipart.File) bool {
	buff := make([]byte, 512)
	file.Read(buff)
	contentType := http.DetectContentType(buff)
	return contentType == "image/png" || contentType == "image/gif" || contentType == "image/webp" || contentType == "image/jpeg"
}
