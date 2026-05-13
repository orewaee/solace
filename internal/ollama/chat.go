package ollama

import "time"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Chat struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
	Tools    []Tool    `json:"tools"`
}

type Response struct {
	TotalDuration time.Duration `json:"total_duration"`
	Message       `json:"message"`
}
