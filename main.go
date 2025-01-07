package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	// Get the OpenAI API key from environment variables
	token := os.Getenv("OPENAI_API_KEY")
	if token == "" {
		fmt.Println("Error: OPENAI_API_KEY environment variable is not set.")
		return
	}

	// Run `git diff` command
	cmd := exec.Command("git", "diff")
	gitDiffOutput, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error running `git diff`: %v\n", err)
		return
	}

	// Trim and format the git diff output
	gitDiff := strings.TrimSpace(string(gitDiffOutput))
	if gitDiff == "" {
		fmt.Println("No changes detected in `git diff`. Please ensure you have staged or modified files.")
		return
	}

	// Prepare the prompt
	prompt := fmt.Sprintf(
		"Based on the following git diff, generate a one-liner commit message with an emoji at the beginning (in English, within 100 characters):\n%s",
		gitDiff,
	)

	// Create OpenAI client
	client := openai.NewClient(token)

	// Make a request to OpenAI GPT-4
	resp, err := client.CreateChatCompletion(
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
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	// Print the generated commit message
	fmt.Println("Generated Commit Message:")
	fmt.Println(resp.Choices[0].Message.Content)
}
