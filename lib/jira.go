package lib

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// GetJiraLink generates a Jira ticket link based on the current Git branch name.
// Returns the Jira link and an error (if any).
func GetJiraLink() (string, error) {
	// Get the current branch name
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get the current branch name: %w", err)
	}

	branchName := strings.TrimSpace(string(output))

	// Define a regex pattern to match the Jira ticket format
	pattern := `^(CAMID|CAMIDSRE)-\d+`
	re := regexp.MustCompile(pattern)

	// Extract the Jira ticket ID
	matches := re.FindString(branchName)
	if matches == "" {
		return "", errors.New("no Jira ticket ID found in the branch name")
	}

	// Construct the Jira link
	jiraBaseURL := "https://moneyforward.atlassian.net/browse/"
	jiraLink := jiraBaseURL + matches

	return jiraLink, nil
}
