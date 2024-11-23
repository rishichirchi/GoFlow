package main

import (
	"gotta-go/handlers"
	"log"

	"github.com/joho/godotenv"
	"gofr.dev/pkg/gofr"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	app := gofr.New()

	app.GET("/chatbot", handlers.GeminiChatBot)
	app.GET("/cold-email", handlers.EmailOutreach)
	app.GET("/pull-requests", handlers.FetchPullRequests)
	app.GET("/make-post", handlers.MakeTweet)

	app.GET("/generate-posts", handlers.GenerateSocialMediaPost)

	app.Run()
}
