package main

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {
	//Get environment data
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

}
