package main

import (
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	//Get environment data
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	openAiKey := os.Getenv("OPENAIKEY")
	if openAiKey == "" {
		log.Fatalln("Missing OPEN AI API KEY")
	}

	ctx := context.Background()
	openAiClient := gpt3.NewClient(openAiKey, gpt3.WithDefaultEngine(gpt3.TextDavinci003Engine))

	resp, err := openAiClient.Completion(ctx, gpt3.CompletionRequest{
		Prompt: []string{"How are you today ?"},
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Choices[0].Text)
}
