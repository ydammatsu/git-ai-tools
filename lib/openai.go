package lib

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func CallOpenAI(prompt string) (string, error) {
	token, err := getOpenAIToken()
	if err != nil {
		return "", err
	}

	resp, err := openai.NewClient(token).CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func getOpenAIToken() (string, error) {
	token := os.Getenv("OPENAI_API_KEY")

	if token == "" {
		return "", fmt.Errorf("error: OPENAI_API_KEY environment variable is not set")
	}

	return token, nil
}
