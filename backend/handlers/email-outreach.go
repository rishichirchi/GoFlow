package handlers

import (
	"context"
	"encoding/json"
	"gotta-go/models"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"gofr.dev/pkg/gofr"
	"google.golang.org/api/option"
)

func GenerateColdEmailForGolang(inputPurpose string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return "", err
	}

	api_key := os.Getenv("GEMINI_API_KEY")
	geminiCtx := context.Background()
	client, err := genai.NewClient(geminiCtx, option.WithAPIKey(api_key))
	if err != nil {
		return "", err
	}
	defer client.Close()

	prompt := `You are an email copywriting assistant. Generate a professional and engaging cold email tailored for reaching out to Golang developers. 
		The email should follow this structure:

		1. **Subject Line** - Catchy and relevant to Golang developers.
		2. **Greeting** - Personalized greeting using the recipient’s name.
		3. **Introduction** - Briefly introduce the product or service, its key benefits, and how it addresses challenges Golang developers face.
		4. **Main Body** - Detailed explanation of how your product/service can help Golang developers, such as improving workflow, scaling, or simplifying tasks. Use bullet points if needed.
		5. **Call to Action** - Encourage the recipient to take action, such as scheduling a demo or checking out resources.
		6. **Closing** - End with a polite sign-off, including your name and company.

		The purpose of the email is as follows: ` + inputPurpose + `

		Make sure the email is concise, personalized, and highlights how our product or service addresses the challenges faced by Golang developers. Include a strong call to action at the end.`

	model := client.GenerativeModel("gemini-1.5-flash")
	response, err := model.GenerateContent(geminiCtx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	marshalResponse, _ := json.MarshalIndent(response, "", "  ")
	var geminiResp models.GeminiResponse
	err = json.Unmarshal(marshalResponse, &geminiResp)
	if err != nil {
		log.Println("Error unmarshalling response:", err)
		return "", err
	}

	if len(geminiResp.Candidates) > 0 && len(geminiResp.Candidates[0].Content.Parts) > 0 {
		parts := geminiResp.Candidates[0].Content.Parts[0]
		return parts, nil
	}

	return "No content found", nil
}

func EmailOutreach(ctx *gofr.Context) (interface{}, error) {
	inputPurpose := "Introducing a tool to simplify concurrency management in Go applications."
	email, err := GenerateColdEmailForGolang(inputPurpose)
	if err != nil {
		return nil, err
	}
	return models.ChatbotResponse{Response: email}, nil
}

