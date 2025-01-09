package lib

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetDiff() (string, error) {
	cmd := exec.Command("git", "diff", "HEAD")

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error running `git diff`: %v", err)
	}

	return strings.TrimSpace(string(output)), nil
}
