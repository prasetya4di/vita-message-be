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
	firebase := network.GetFirebase()
	localizer := translation.LoadTranslation()

	messageDao := impl.NewMessageDao(gormDb)
	imageDao := impl.NewImageDao(gormDb)
	userDao := impl.NewUserDao(gormDb)
	settingDao := impl.NewSettingDao(gormDb)
	messageService := impl2.NewMessageService(openAiClient, firebase)
	imageService := impl2.NewImageService(openAiClient)

	messageRepository := impl3.NewMessageRepository(messageDao, messageService)
	imageRepository := impl3.NewImageRepository(imageDao, imageService)
	userRepository := impl3.NewUserRepository(userDao)
	settingRepository := impl3.NewSettingRepository(settingDao)

	sendMessageUseCase := impl4.NewSendMessage(messageRepository, settingRepository)
	replyMessageUseCase := impl4.NewReplyMessage(messageRepository, settingRepository)
	saveMessageUseCase := impl4.NewSaveMessage(messageRepository)
	getMessageUseCase := impl4.NewGetMessage(messageRepository)
	getCurrentUserUseCase := impl4.NewGetCurrentUser(userRepository)
	uploadImageUseCase := impl4.NewUploadImage(imageRepository, settingRepository, messageRepository)
	loginUseCase := impl4.NewLoginUseCase(userRepository)
	registerUseCase := impl4.NewRegisterUseCase(userRepository)
	addInitialMessageUseCase := impl4.NewAddInitialMessage(messageRepository)
	saveMessagesUseCase := impl4.NewSaveMessages(messageRepository)

	messageHandler := impl5.NewMessageHandler(sendMessageUseCase, replyMessageUseCase, getMessageUseCase, getCurrentUserUseCase, saveMessagesUseCase)
	imageHandler := impl5.NewImageHandler(uploadImageUseCase, replyMessageUseCase, saveMessageUseCase, getCurrentUserUseCase, localizer)
	authHandler := impl5.NewAuthHandler(loginUseCase, registerUseCase, addInitialMessageUseCase)

	rest.LoadRoutes(messageHandler, imageHandler, authHandler)
}
