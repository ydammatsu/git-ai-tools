package main

import (
	"fmt"
	"os"

	"github.com/ydammatsu/git-ai-tools/lib"
)

func main() {
	arg := os.Args[0]

	if arg == "title" {
		title, err := genGitHubTitle()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(title)
		return
	}

	if arg == "body" {
		body, err := genGitHubBody()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(body)
		return
	}

	if arg == "commit" {
		message, err := genCommitMessage()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(message)
		return
	}

	fmt.Println("Invalid argument")
	os.Exit(1)
}

func genCommitMessage() (string, error) {
	gitDiff, err := lib.GetDiff("commit")
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

func genGitHubTitle() (string, error) {
	gitDiff, err := lib.GetDiff("branch")
	if err != nil {
		return "", err
	}

	title, err := lib.CallOpenAI(
		lib.GetGitHubTitlePrompt(gitDiff),
	)
	if err != nil {
		return "", err
	}

	return title, nil
}

func genGitHubBody() (string, error) {
	gitDiff, err := lib.GetEntireDiff()
	if err != nil {
		return "", err
	}

	jiraLink, err := lib.GetJiraLink()
	if err != nil {
		jiraLink = "none"
	}

	body, err := lib.CallOpenAI(
		lib.GetGitHubBodyPrompt(gitDiff, jiraLink),
	)
	if err != nil {
		return "", err
	}

	return body, nil
}
