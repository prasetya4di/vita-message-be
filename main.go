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
	gormDb := local.GetGormDb()
	openAiClient := network.GetOpenAi()
	localizer := translation.LoadTranslation()

	messageDao := impl.NewMessageDao(gormDb)
	imageDao := impl.NewImageDao(gormDb)
	userDao := impl.NewUserDao(gormDb)
	energyDao := impl.NewEnergyDao(gormDb)
	messageService := impl2.NewMessageService(openAiClient)
	imageService := impl2.NewImageService()

	messageRepository := impl3.NewMessageRepository(messageDao, messageService)
	imageRepository := impl3.NewImageRepository(imageDao, imageService)
	userRepository := impl3.NewUserRepository(userDao)
	energyRepository := impl3.NewEnergyRepository(energyDao)

	sendMessageUseCase := impl4.NewSendMessage(messageRepository)
	replyMessageUseCase := impl4.NewReplyMessage(messageRepository)
	saveMessageUseCase := impl4.NewSaveMessage(messageRepository)
	getMessageUseCase := impl4.NewGetMessage(messageRepository)
	getCurrentUserUseCase := impl4.NewGetCurrentUser(userRepository)
	uploadImageUseCase := impl4.NewUploadImage(imageRepository)
	loginUseCase := impl4.NewLoginUseCase(userRepository)
	registerUseCase := impl4.NewRegisterUseCase(userRepository)
	addInitialMessageUseCase := impl4.NewAddInitialMessage(messageRepository)
	addEnergy := impl4.NewAddEnergy(energyRepository)

	messageHandler := impl5.NewMessageHandler(sendMessageUseCase, replyMessageUseCase, getMessageUseCase, getCurrentUserUseCase)
	imageHandler := impl5.NewImageHandler(uploadImageUseCase, replyMessageUseCase, saveMessageUseCase, getCurrentUserUseCase, localizer)
	authHandler := impl5.NewAuthHandler(loginUseCase, registerUseCase, addInitialMessageUseCase, addEnergy)

	rest.LoadRoutes(messageHandler, imageHandler, authHandler)
}
