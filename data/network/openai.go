package network

import (
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

func GetOpenAi() *openai.Client {
	openAiKey := os.Getenv("OPENAIKEY")
	if openAiKey == "" {
		log.Fatalln("Missing OPEN AI API KEY")
	}

	return openai.NewClient(openAiKey)
}
