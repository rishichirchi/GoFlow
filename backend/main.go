package main

import (
	"gotta-go/handlers"
	"gotta-go/models"
	"log"

	"github.com/joho/godotenv"
	"gofr.dev/pkg/gofr"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	app := gofr.New()

	app.GET("/chatbot", handlers.GeminiChatBot)
	app.GET("/coldemail", func(ctx *gofr.Context) (interface{}, error) {
		inputPurpose := "Introducing a tool to simplify concurrency management in Go applications."
		email, err := handlers.GenerateColdEmailForGolang(inputPurpose)
		if err != nil {
			return nil, err
		}
		return models.ChatbotResponse{Response: email}, nil
	})
	app.GET("/pull-requests", handlers.FetchPullRequests)
	
	app.Run()
}