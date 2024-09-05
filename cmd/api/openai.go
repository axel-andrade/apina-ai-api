package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const apiKey = "YOUR_API_KEY"

type ChatGPTRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

func formatText(input string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	requestBody, err := json.Marshal(ChatGPTRequest{
		Model: "gpt-4",
		Messages: []Message{
			{
				Role:    "system",
				Content: "You are a helpful assistant that corrects and formats poorly formatted English text.",
			},
			{
				Role:    "user",
				Content: input,
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("error creating request body: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making API request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	var chatGPTResponse ChatGPTResponse
	if err := json.Unmarshal(body, &chatGPTResponse); err != nil {
		return "", fmt.Errorf("error unmarshalling response: %v", err)
	}

	if len(chatGPTResponse.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}

	return chatGPTResponse.Choices[0].Message.Content, nil
}

func init() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide text to format")
		return
	}

	input := os.Args[1]
	formattedText, err := formatText(input)
	if err != nil {
		fmt.Printf("Error formatting text: %v\n", err)
		return
	}

	fmt.Println("Formatted text:")
	fmt.Println(formattedText)
}
