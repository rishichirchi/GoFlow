package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"net/smtp"

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
    You are an expert copywriter crafting cold outreach emails for Golang developers.
    The goal is to create a concise yet highly engaging and professional email that conveys value to the recipient.

    Please generate the email with the following sections:
    
    1. **Subject Line**: A catchy and clear subject line for a Golang developer (e.g., "Enhance Your Golang Development Today").
    2. **Greeting**: A personalized greeting, e.g., "Hi there," or "Hello Golang Enthusiast,".
    3. **Introduction**: A brief 2-3 sentence introduction to your product/service, explaining what it does and why it’s relevant to the recipient.
    4. **Value Proposition (Main Body)**: 
        - A few bullet points that clearly explain how your product/service solves problems Golang developers face (e.g., increasing productivity, simplifying code, handling concurrency).
    5. **Call to Action**: A strong and clear CTA like "Get a demo today," or "Learn more about our solution."
    6. **Closing**: A polite sign-off, like "Best regards," followed by your name and company.

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


func SendEmail(to, subject, body string) error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return err
	}

	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	senderEmail := os.Getenv("SENDER_EMAIL")
	senderPassword := os.Getenv("SENDER_PASSWORD")

	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)

	message := []byte(fmt.Sprintf("Subject: %s\n\n%s", subject, body))
	err = smtp.SendMail(smtpServer+":"+smtpPort, auth, senderEmail, []string{to}, message)
	if err != nil {
		return err
	}
	return nil
}


func EmailOutreach(ctx *gofr.Context) (interface{}, error) {
	
	inputPurpose := "Introducing a tool to simplify concurrency management in Go applications."

	
	emailBody, emailSubject, err := GenerateColdEmailForGolang(inputPurpose)
	if err != nil {
		return nil, err
	}

	
	log.Printf("Generated Email Subject: %s\n", emailSubject)
	log.Printf("Generated Email Body: %s\n", emailBody)

	
	recipientEmail := "eduksh01@gmail.com"
	err = SendEmail(recipientEmail, emailSubject, emailBody)
	if err != nil {
		return nil, fmt.Errorf("failed to send email: %v", err)
	}

	
	return models.ChatbotResponse{
		Response: fmt.Sprintf("Email subject: %s\n\nEmail body:\n%s", emailSubject, emailBody),
	}, nil
}
