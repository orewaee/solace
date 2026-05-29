package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"orewaee.dev/solace/internal/ollama"
)

func main() {
	system := fmt.Sprintf(`
You are the user reminder planner agent. You have two tools to choose from: create_event, ask_clarification.
You should definitely choose one of them. If the data provided by the user in the request is enough for you,
then use create_event, otherwise specify the unknown using ask_clarification. Communication with the user and
the text of the event must be in the user's language. The user currently has %s in RFC 3339.

For create_event: the 'text' field must be a normalized version of the user's request - not just keywords,
not a verbose replay - but a clean, natural sentence suitable for a reminder.

CRITICAL: Preserve the exact lexical form of all nouns, proper names, and specialized terminology.
Do NOT apply declensions, case changes, or any morphological modifications. Output each term in the
same grammatical case and number as the user provided.
		`, time.Now().Format(time.RFC3339))

	chat := &ollama.Chat{
		Model: "gpt-oss:120b-cloud",
		Messages: []ollama.Message{
			{Role: "user", Content: "Install Arch tomorrow"},
			{Role: "system", Content: system},
		},
		Stream: false,
		Tools: []ollama.Tool{
			{
				Type: "object",
				Function: ollama.Function{
					Name:        "create_event",
					Description: "Creates a reminder for the user",
					Parameters: ollama.Parameters{
						Type:     "object",
						Required: []string{"timestamp", "text"},
						Properties: map[string]ollama.Property{
							"timestamp": {
								Type:        "string",
								Description: "RFC 3339 timestamp when the reminder will be activated",
							},
							"text": ollama.Property{
								Type:        "string",
								Description: "The message that will be sent to the user",
							},
						},
					},
				},
			},
			{
				Type: "object",
				Function: ollama.Function{
					Name:        "ask_clarification",
					Description: "Ask the user for information",
					Parameters: ollama.Parameters{
						Type:     "object",
						Required: []string{"text"},
						Properties: map[string]ollama.Property{
							"text": {
								Type:        "string",
								Description: "Clarifying question to the user",
							},
						},
					},
				},
			},
		},
	}

	data, err := json.Marshal(chat)
	if err != nil {
		panic(err)
	}

	response, err := http.Post("http://localhost:11434/api/chat", "application/json", bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var result map[string]any
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		panic(err)
	}

	pretty, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(pretty))
}
