package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"mime/multipart"
	"net/http"
	"vita-message-service/data/entity"
	"vita-message-service/delivery/rest/handler"
	"vita-message-service/usecase"
	constant "vita-message-service/util/const"
	time2 "vita-message-service/util/local_time"
	"vita-message-service/util/translation"
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
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	if !isValidImage(file) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid image format"})
		return
	}
	currentUser, err := ih.getCurrentUser.Invoke(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}

	scan, err := ih.uploadImage.Invoke(currentUser.Email, file, header)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	if len(scan.Possibilities) == 1 {
		replyMessage := entity.Message{
			Email:       currentUser.Email,
			Message:     scan.Possibilities[0].Description,
			CreatedDate: time2.CurrentTime(),
			MessageType: constant.Reply,
			FileType:    constant.Text,
		}
		messages, err := ih.replyMessage.Invoke(currentUser, replyMessage)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		scan.Messages = append(scan.Messages, messages...)
	} else if len(scan.Possibilities) == 0 {
		replyMessage := entity.Message{
			Email:       currentUser.Email,
			Message:     translation.UnknownImageMessage(ih.localizer),
			CreatedDate: time2.CurrentTime(),
			MessageType: constant.Reply,
			FileType:    constant.Text,
		}
		message, err := ih.saveMessage.Invoke(replyMessage)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		scan.Messages = append(scan.Messages, message)
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"data": scan})
}

func isValidImage(file multipart.File) bool {
	buff := make([]byte, 512)
	file.Read(buff)
	contentType := http.DetectContentType(buff)
	return contentType == "image/png" || contentType == "image/jpg" || contentType == "image/gif" || contentType == "image/webp" || contentType == "image/jpeg" || contentType == "image/bmp"
}
