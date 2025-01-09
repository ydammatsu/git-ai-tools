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
		"Generate a GitHub issue title string one liner based on the following git diff (in English, within %d characters):\n%s",
		maxGitHubTitleLength,
		gitDiff,
	)
}

func GetGitHubBodyPrompt(gitDiff string, jiraLink string) string {
	// Generate the instruction prompt
	format := `Based on the following git diff, generate a Markdown document with the following structure:

## Jira
%s

## What
From the provided git diff, extract and summarize 1 to 3 key points describing the changes made. Ensure these are concise, relevant, and easy to understand.

## Why
Leave this section empty.

Here is the git diff: %s`
	return fmt.Sprintf(
		format,
		jiraLink,
		gitDiff,
	)
}
