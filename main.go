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
	"vita-message-service/util/translation"
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
	gormDb := local.GetGormDb()
	openAiClient := network.GetOpenAi()
	localizer := translation.LoadTranslation()

	messageDao := impl.NewMessageDao(db)
	imageDao := impl.NewImageDao(db)
	userDao := impl.NewUserDao(gormDb)
	messageService := impl2.NewMessageService(openAiClient)
	imageService := impl2.NewImageService()

	messageRepository := impl3.NewMessageRepository(messageDao, messageService)
	imageRepository := impl3.NewImageRepository(imageDao, imageService)
	userRepository := impl3.NewUserRepository(userDao)

	sendMessageUseCase := impl4.NewSendMessage(messageRepository)
	replyMessageUseCase := impl4.NewReplyMessage(messageRepository)
	getMessageUseCase := impl4.NewGetMessage(messageRepository)
	uploadImageUseCase := impl4.NewUploadImage(imageRepository)
	loginUseCase := impl4.NewLoginUseCase(userRepository)
	registerUseCase := impl4.NewRegisterUseCase(userRepository)

	messageHandler := impl5.NewMessageHandler(sendMessageUseCase, replyMessageUseCase, getMessageUseCase)
	imageHandler := impl5.NewImageHandler(uploadImageUseCase, replyMessageUseCase, localizer)
	authHandler := impl5.NewAuthHandler(loginUseCase, registerUseCase)

	rest.LoadRoutes(messageHandler, imageHandler, authHandler)
}
