package main

import (
	"fmt"
	"log"

	"github.com/sashabaranov/go-openai"
)

func commit(oaiClient *openai.Client, cfg *Config) error {
	if !isGitRepo() {
		log.Fatal("ERR: not a git repository")
	}

	branch := getBranchName()
	diff := getCurrentDiff()
	msg := genCommitMessage(oaiClient, cfg, branch, diff)
	fmt.Printf("%s\n", msg)

	return nil
}

func changelogEntry(oaiClient *openai.Client, cfg *Config, append bool, changelogFile string) error {
	if !isGitRepo() {
		log.Fatal("ERR: not a git repository")
	}

	branch := getBranchName()
	diff := getCurrentDiff()
	msg := genChangelogEntry(oaiClient, cfg, branch, diff)

	if append {
		appendChangelog(msg, changelogFile)
	} else {
		fmt.Printf("%s\n", msg)
	}

	return nil
}
