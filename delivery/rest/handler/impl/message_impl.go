package impl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vita-message-service/data/entity"
	"vita-message-service/delivery/rest/handler"
	"vita-message-service/usecase"
)

type messageHandler struct {
	sendMessage    usecase.SendMessage
	replyMessage   usecase.ReplyMessage
	getMessage     usecase.GetMessage
	getCurrentUser usecase.GetCurrentUser
}

func NewMessageHandler(message usecase.SendMessage, replyMessage usecase.ReplyMessage, getMessage usecase.GetMessage, getCurrentUser usecase.GetCurrentUser) handler.MessageHandler {
	return &messageHandler{
		sendMessage:    message,
		replyMessage:   replyMessage,
		getMessage:     getMessage,
		getCurrentUser: getCurrentUser,
	}
}

func (mh *messageHandler) SendMessage(c *gin.Context) {
	var newMessage entity.Message

	if err := c.BindJSON(&newMessage); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}

	currentUser, err := mh.getCurrentUser.Invoke(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}

	newMessage.Email = currentUser.Email
	messages, err := mh.sendMessage.Invoke(newMessage)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		c.IndentedJSON(http.StatusCreated, gin.H{"data": messages})
	}
}

func (mh *messageHandler) ReplyMessage(c *gin.Context) {
	var newMessage entity.Message

	if err := c.BindJSON(&newMessage); err != nil {
		return
	}

	currentUser, err := mh.getCurrentUser.Invoke(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}

	newMessage.Email = currentUser.Email
	messages, err := mh.replyMessage.Invoke(newMessage)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		c.IndentedJSON(http.StatusCreated, gin.H{"data": messages})
	}
}

func (mh *messageHandler) GetMessage(c *gin.Context) {
	currentUser, err := mh.getCurrentUser.Invoke(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}

	messages, err := mh.getMessage.Invoke(currentUser.Email)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"data": messages})
	}
}
