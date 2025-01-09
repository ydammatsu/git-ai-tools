package lib

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetDiff(opt string) (string, error) {
	cmdLastArg := ""
	if opt == "commit" {
		cmdLastArg = "HEAD"
	}
	if opt == "branch" {
		cmdLastArg = "origin/HEAD"
	}
	cmd := exec.Command("git", "diff", cmdLastArg)

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error running `git diff`: %v", err)
	}

	return strings.TrimSpace(string(output)), nil
}
