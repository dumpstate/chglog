package main

import (
	"log"
	"os/exec"
	"strings"
)

func isGitRepo() bool {
	// TODO
	return true
}

func execGitCmd(args ...string) string {
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(out))
}

func getBranchName() string {
	return execGitCmd("rev-parse", "--abbrev-ref", "HEAD")
}

func getCurrentDiff() string {
	diff := execGitCmd("diff", "--no-color", "--staged")
	if diff == "" {
		diff = execGitCmd("diff", "--no-color")
	}
	return diff
}
