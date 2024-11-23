package handlers

import (
	"context"
	"encoding/json"
	"gotta-go/models"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"gofr.dev/pkg/gofr"
	"google.golang.org/api/option"
)

func GenerateCodeReview(ctx *gofr.Context)(interface{}, error){
	var filepath models.File

	err := ctx.Bind(&filepath)

	if err != nil{
		log.Println("Error binding file:", err)
		return nil, err
	}

	log.Println("File:", filepath)

	file, err := os.Open(filepath.Name)

	if err != nil {
		log.Println("Error opening file:", err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error reading file:", err)
	}

	log.Println("file content:", string(content))

	instructions := "instructions/code-review.md"

	instructionFile, err := os.Open(instructions)

	if err != nil{
		log.Println("Error opening file:", err)
	}

	defer instructionFile.Close()

	instructionContent, err := ioutil.ReadAll(instructionFile)

	if err != nil {
		log.Println("Error reading file:", err)
	}

	geminiPrompt := string(content) + string(instructionContent)

	api_key := os.Getenv("GEMINI_API_KEY")

	geminiCtx := context.Background()

	client, err := genai.NewClient(geminiCtx, option.WithAPIKey(api_key))

	if err != nil {
		return nil, err
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	// model.SetMaxOutputTokens(100)

	response, err := model.GenerateContent(geminiCtx, genai.Text(geminiPrompt))

	if err != nil {
		log.Println(err)
	}

	marshalResponse, _ := json.MarshalIndent(response, "", "  ")

	if err != nil {
		log.Println("Error marshalling response:", err)
		return nil, err
	}

	return string(marshalResponse), nil

}