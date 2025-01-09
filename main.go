package main

import (
	"fmt"
	"os"

	"github.com/ydammatsu/git-ai-tools/lib"
)

func main() {
	message, err := genCommitMessage()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(message)
}

func genCommitMessage() (string, error) {
	gitDiff, err := lib.GetDiff()
	if err != nil {
		return "", err
	}

	message, err := lib.CallOpenAI(
		lib.GenCommitMessagePrompt(gitDiff),
	)
	if err != nil {
		return "", err
	}

	return message, nil
}
