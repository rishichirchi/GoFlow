package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gotta-go/models"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"gofr.dev/pkg/gofr"
	"google.golang.org/api/option"
)

func GenerateColdEmailForGolang(inputPurpose string) (string, string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return "", "", err
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	geminiCtx := context.Background()
	client, err := genai.NewClient(geminiCtx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", "", err
	}
	defer client.Close()

	prompt := fmt.Sprintf(`
You are a skilled copywriter specializing in creating high-impact cold outreach emails for Golang developers. Your goal is to craft a concise, engaging, and professional email that resonates with the recipient and clearly conveys the value of the product or service.

Please generate the email with the following sections:

1. **Subject Line**: Create a compelling and direct subject line that would grab the attention of a Golang developer. Keep it brief but informative, such as "Boost Your Golang Development Efficiency Today."
2. **Greeting**: Start with a warm, professional greeting. Examples include "Hello [Name]," or "Hi [Developer]," tailored to the recipient's role or interests.
3. **Introduction**: Write a brief, 2-3 sentence introduction that clearly explains who you are, what your company does, and why your offering is relevant to the recipient. Make it clear how your product or service addresses the challenges Golang developers commonly face.
4. **Value Proposition (Main Body)**: 
    - Present a clear, concise explanation of how your product/service directly benefits Golang developers. Highlight key features and solutions, such as improving productivity, simplifying concurrency, enhancing code maintainability, or optimizing performance.
    - Use bullet points for easy readability and to emphasize key points. For example:
        * **Simplify Concurrency**: Eliminate the complexities of managing goroutines with our intuitive tools.
        * **Increase Productivity**: Streamline your development process and reduce debugging time.
        * **Improve Code Quality**: Write cleaner, more maintainable code with fewer errors.
5. **Call to Action**: End with a strong and clear call to action that encourages the recipient to take the next step. This could be an invitation for a demo, a link to learn more, or an offer to schedule a conversation. Examples:
    - "Request a free demo today to see how we can help."
    - "Click here to schedule a quick demo and experience the difference."
    - "Learn more about how we can enhance your Golang development."
6. **Closing**: Finish with a polite, professional sign-off. Example: "Best regards," followed by your name and company. Keep it friendly but formal.

**Purpose of the email**: %s
	`, inputPurpose)

	model := client.GenerativeModel("gemini-1.5-flash")
	response, err := model.GenerateContent(geminiCtx, genai.Text(prompt))
	if err != nil {
		return "", "", err
	}

	marshalResponse, _ := json.MarshalIndent(response, "", "  ")
	var geminiResp models.GeminiResponse
	err = json.Unmarshal(marshalResponse, &geminiResp)
	if err != nil {
		log.Println("Error unmarshalling response:", err)
		return "", "", err
	}

	if len(geminiResp.Candidates) > 0 && len(geminiResp.Candidates[0].Content.Parts) > 0 {
		emailBody := geminiResp.Candidates[0].Content.Parts[0]
		emailSubject := "Exciting Tool for Golang Developers"
		return emailBody, emailSubject, nil
	}

	return "No content found", "No subject", nil
}

func EmailOutreach(ctx *gofr.Context) (interface{}, error) {
	inputPurpose := "Introducing a tool to simplify concurrency management in Go applications."

	emailBody, emailSubject, err := GenerateColdEmailForGolang(inputPurpose)
	if err != nil {
		return nil, err
	}

	log.Printf("Generated Email Subject: %s\n", emailSubject)
	log.Printf("Generated Email Body: %s\n", emailBody)

	return models.ChatbotResponse{
		Response: fmt.Sprintf("Email subject: %s\n\nEmail body:\n%s", emailSubject, emailBody),
	}, nil
}
