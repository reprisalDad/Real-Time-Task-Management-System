package services

import (
    "context"
    "time"

    openai "github.com/sashabaranov/go-openai"
)

var openAIClient = openai.NewClient("YOUR_OPENAI_API_KEY") // Replace with env variable

// GetTaskSuggestions calls the AI API to generate task breakdowns.
func GetTaskSuggestions(description string) (string, error) {
    req := openai.ChatCompletionRequest{
        Model: openai.GPT3Dot5Turbo, // or use Gemini if available
        Messages: []openai.ChatCompletionMessage{
            {Role: "system", Content: "You are an expert project manager. Break down the following task into actionable subtasks."},
            {Role: "user", Content: description},
        },
        MaxTokens: 150,
        Timeout:   10 * time.Second,
    }
    resp, err := openAIClient.CreateChatCompletion(context.Background(), req)
    if err != nil {
        return "", err
    }
    if len(resp.Choices) > 0 {
        return resp.Choices[0].Message.Content, nil
    }
    return "", nil
}
