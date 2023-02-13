package main

import (
	"github.com/joho/godotenv"
	"log"
	"vita-message-service/data/local"
	"vita-message-service/data/local/impl"
	"vita-message-service/data/network"
	impl2 "vita-message-service/data/network/impl"
	"vita-message-service/delivery/rest"
	impl5 "vita-message-service/delivery/rest/handler/impl"
	impl3 "vita-message-service/repository/impl"
	impl4 "vita-message-service/usecase/impl"
)

func init() {
	//Get environment data
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db := local.GetDB()
	openAiClient := network.GetOpenAi()

	messageDao := impl.NewMessageDao(db)
	messageService := impl2.NewMessageService(openAiClient)
	messageRepository := impl3.NewMessageRepository(messageDao, messageService)

	sendMessageUseCase := impl4.NewSendMessage(messageRepository)
	replyMessageUseCase := impl4.NewReplyMessage(messageRepository)
	getMessageUseCase := impl4.NewGetMessage(messageRepository)

	messageHandler := impl5.NewMessageHandler(sendMessageUseCase, replyMessageUseCase, getMessageUseCase)

	rest.LoadRoutes(messageHandler)
}
