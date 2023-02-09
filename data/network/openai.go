package network

import (
	"github.com/PullRequestInc/go-gpt3"
	"log"
	"os"
)

func GetOpenAi() gpt3.Client {
	openAiKey := os.Getenv("OPENAIKEY")
	if openAiKey == "" {
		log.Fatalln("Missing OPEN AI API KEY")
	}

	openAiClient := gpt3.NewClient(openAiKey, gpt3.WithDefaultEngine(gpt3.TextDavinci003Engine))
	return openAiClient
}
