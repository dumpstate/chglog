package main

import (
	"log"
	"os/exec"
	"strings"
)

func execGitCmd(args ...string) string {
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(out))
}

func isGitRepo() bool {
	out := execGitCmd("remote")
	return !strings.HasPrefix(out, "fatal: not a git repository")
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
