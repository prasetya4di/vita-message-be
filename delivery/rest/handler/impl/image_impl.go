package impl

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"time"
	"vita-message-service/data/entity"
	"vita-message-service/delivery/rest/handler"
	"vita-message-service/usecase"
	constant "vita-message-service/util/const"
)

type imageHandler struct {
	uploadImage  usecase.UploadImage
	replyMessage usecase.ReplyMessage
}

func NewImageHandler(uploadImage usecase.UploadImage, message usecase.ReplyMessage) handler.ImageHandler {
	return &imageHandler{
		uploadImage,
		message,
	}
}

func (ih *imageHandler) UploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	if !isValidImage(file) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid image format"})
		return
	}
	email := c.Param("email")
	scan, err := ih.uploadImage.Invoke(email, file, header)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	if len(scan.Possibilities) == 1 {
		replyMessage := entity.Message{
			Email:       email,
			Message:     scan.Possibilities[0].Description,
			CreatedDate: time.Now(),
			MessageType: constant.Reply,
			FileType:    constant.Text,
		}
		messages, err := ih.replyMessage.Invoke(replyMessage)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		scan.Messages = append(scan.Messages, messages...)
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"data": scan})
}

func isValidImage(file multipart.File) bool {
	buff := make([]byte, 512)
	file.Read(buff)
	contentType := http.DetectContentType(buff)
	return contentType == "image/png" || contentType == "image/jpg" || contentType == "image/gif" || contentType == "image/webp" || contentType == "image/jpeg" || contentType == "image/bmp"
}
