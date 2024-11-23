package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"gotta-go/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/generative-ai-go/genai"
	"gofr.dev/pkg/gofr"
	"google.golang.org/api/option"
)

func FetchPullRequests()(interface{}, error) {
	const (
		owner = "gofr-dev"
		repo = "gofr"
		apiVersion = "2022-11-28"
	)

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls?state=all", owner, repo)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("Error creating request: " + err.Error())
	}

	token := os.Getenv("GITHUB_TOKEN")

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", apiVersion)

	client := &http.Client{}

	response, err :=client.Do(req)

	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", response.StatusCode)
	}

	// Read and display the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	// log.Println(string(body))

	var results []models.GithubPR

	err = json.Unmarshal(body, &results)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return err.Error(), nil
	}

	return results[0], nil
}


func GenerateSocialMediaPost(ctx *gofr.Context) (interface{}, error) {
	filepath := "instructions/sm-post.md"
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

	prompt, _ := FetchPullRequests()

	geminiPrompt := string(content) + fmt.Sprintf("%v", prompt)

	api_key := os.Getenv("GEMINI_API_KEY")

	geminiCtx := context.Background()

	client, err := genai.NewClient(geminiCtx, option.WithAPIKey(api_key))

	if err != nil {
		return models.ChatbotResponse{Response: err.Error()}, err
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	model.SetMaxOutputTokens(300)

	response, err := model.GenerateContent(geminiCtx, genai.Text(geminiPrompt))

	if err != nil {
		log.Fatal(err)
	}

	marshalResponse, _ := json.MarshalIndent(response, "", "  ")

	var geminiResp models.GeminiResponse
	err = json.Unmarshal(marshalResponse, &geminiResp)
	if err != nil {
		log.Println("Error unmarshalling response:", err)
		return err.Error(), err
	}

	if len(geminiResp.Candidates) > 0 && len(geminiResp.Candidates[0].Content.Parts) > 0 {
		parts := geminiResp.Candidates[0].Content.Parts[0]
		log.Println("response:", parts)
		return parts, nil
	}

	return err.Error(), nil
}

