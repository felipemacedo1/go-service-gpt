package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type OpenAIService struct {
	apiKey string
}

func NewOpenAIService(apiKey string) *OpenAIService {
	return &OpenAIService{
		apiKey: apiKey,
	}
}

type ChatCompletionRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func (s *OpenAIService) SendMessage(message string) (string, error) {
	requestBody := ChatCompletionRequest{
		Model: "gpt-4o",
		Messages: []Message{
			{Role: "user", Content: message},
		},
	}

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	var response ChatCompletionResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}

	return response.Choices[0].Message.Content, nil
}
