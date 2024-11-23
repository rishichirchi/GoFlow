package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
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

type ChatbotResponse struct{
	Response string `json:"response"`
}


func GeminiChatBot(ctx *gofr.Context) (interface{}, error) {
	prompt := ctx.Request.Param("prompt")

	filepath := "gofr-introduction.md"
	file, err := os.Open(filepath)

	if err != nil{
		log.Println("Error opening file:", err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error reading file:", err)
	}

	log.Println("file content:", string(content))

	geminiPrompt := string(content) + prompt

	api_key := os.Getenv("GEMINI_API_KEY")

	geminiCtx := context.Background()

	client, err := genai.NewClient(geminiCtx, option.WithAPIKey(api_key))

	if err != nil {
		return ChatbotResponse{Response: err.Error()}, err
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	model.SetMaxOutputTokens(100)

	response, err := model.GenerateContent(geminiCtx, genai.Text(geminiPrompt))

	if err != nil {
		log.Fatal(err)
	}

	marshalResponse, _ := json.MarshalIndent(response, "", "  ")

	var geminiResp GeminiResponse
	err = json.Unmarshal(marshalResponse, &geminiResp)
	if err != nil {
		log.Println("Error unmarshalling response:", err)
		return ChatbotResponse{Response: err.Error()}, err
	}

	if len(geminiResp.Candidates) > 0 && len(geminiResp.Candidates[0].Content.Parts) > 0 {
		parts := geminiResp.Candidates[0].Content.Parts[0]
		log.Println("response:", parts)
		return ChatbotResponse{Response: parts}, nil
	}

	return ChatbotResponse{Response: "error"}, nil
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