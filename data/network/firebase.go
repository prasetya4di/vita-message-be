package network

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
	"log"
	"os"
)

func GetFirebase() *messaging.Client {
	// Replace with your project's credentials file path
	opt := option.WithCredentialsFile(os.Getenv("FIREBASECREDENTIAL"))

	// Initialize the app with credentials
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v\n", err)
	}

	// Get a messaging client
	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("Failed to create messaging client: %v\n", err)
	}

	return client
}
