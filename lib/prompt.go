package lib

import "fmt"

var maxCommitMessageLength = 60

func GenCommitMessagePrompt(gitDiff string) string {
	return fmt.Sprintf(
		"You are a super engineer. Based on the following git diff, generate a one-liner commit message with an emoji at the beginning that matches the content (in English, within %d characters):\n%s",
		maxCommitMessageLength,
		gitDiff,
	)
}

var maxGitHubTitleLength = 40

func GetGitHubTitlePrompt(gitDiff string) string {
	return fmt.Sprintf(
		"Generate a GitHub issue title based on the following git diff (in English, within %d characters):\n%s",
		maxGitHubTitleLength,
		gitDiff,
	)
}

func GetGitHubBodyPrompt(gitDiff string, jiraLink string) string {
	return fmt.Sprintf(
		"Generate a GitHub issue body based on the following git diff (in English):\n%s",
		gitDiff,
	)
}
