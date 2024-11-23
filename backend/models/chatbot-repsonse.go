package models

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