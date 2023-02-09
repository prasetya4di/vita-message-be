package impl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vita-message-service/data/entity"
	"vita-message-service/delivery/rest/handler"
	"vita-message-service/usecase"
)

type messageHandler struct {
	sendMessage usecase.SendMessage
	getMessage  usecase.GetMessage
}

func NewMessageHandler(message usecase.SendMessage, getMessage usecase.GetMessage) handler.MessageHandler {
	return &messageHandler{
		sendMessage: message,
		getMessage:  getMessage,
	}
}

func (mh *messageHandler) SendMessage(c *gin.Context) {
	var newMessage entity.Message

	if err := c.BindJSON(&newMessage); err != nil {
		return
	}

	messages, err := mh.sendMessage.Invoke(newMessage)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		c.IndentedJSON(http.StatusCreated, messages)
	}
}

func (mh *messageHandler) GetMessage(c *gin.Context) {
	email := c.Param("email")
	messages, err := mh.getMessage.Invoke(email)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	} else {
		c.IndentedJSON(http.StatusOK, messages)
	}
}
