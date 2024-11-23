package main

import (
	"log"
	"context"
	"encoding/json"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

	"github.com/joho/godotenv"
	"gofr.dev/pkg/gofr"
)

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []string `json:"Parts"`
		} `json:"Content"`
	} `json:"Candidates"`
}


func GeminiChatBot(ctx *gofr.Context) (interface{}, error) {
	api_key := os.Getenv("GEMINI_API_KEY")

	geminiCtx := context.Background()

	client, err := genai.NewClient(geminiCtx, option.WithAPIKey(api_key))

	if err != nil {
		return nil, err
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	response, err := model.GenerateContent(geminiCtx, genai.Text("Hello, how are you today?"))

	if err != nil {
		log.Fatal(err)
	}

	marshalResponse, _ := json.MarshalIndent(response, "", "  ")

	var geminiResp GeminiResponse
	err = json.Unmarshal(marshalResponse, &geminiResp)
	if err != nil {
		log.Println("Error unmarshalling response:", err)
		return nil, err
	}

	if len(geminiResp.Candidates) > 0 && len(geminiResp.Candidates[0].Content.Parts) > 0 {
		parts := geminiResp.Candidates[0].Content.Parts[0]
		log.Println("response:", parts)
		return parts, nil
	}

	return "No content found", nil
}
func main(){
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	app := gofr.New()

	app.GET("/chatbot", GeminiChatBot)
	
	app.Run()
}