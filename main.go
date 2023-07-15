package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	openai "github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
)

type Config struct {
	OpenAI struct {
		Token string `json:"token"`
	} `json:"openai"`
}

func loadConfig() *Config {
	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	cfgPath := filepath.Join(userHome, ".chglog.json")
	_, err = os.Stat(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	err = json.Unmarshal(content, &cfg)
	return &cfg
}

func getDiff() string {
	out, err := exec.Command("git", "diff", "--no-color", "HEAD~1", "HEAD").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func summarise(client *openai.Client, diff string) string {
	res, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleAssistant,
					Content: fmt.Sprintf(`
						You're a programmer's assistant. Given the git diff below, summarise and write a changelog
						entry for the change. Provide your output right away, no prefixes, such that it could be
						used as a changelog entry.

						START OF DIFF
						%s
						END OF DIFF
					`, diff),
				},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	return res.Choices[0].Message.Content
}

func commitMessage(client *openai.Client, diff string) string {
	res, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleAssistant,
					Content: fmt.Sprintf(`
						You're a programmer's assistant. Given the git diff below, prepare a short
						, single line, git commit title, 120 characters max. Prepare your output
						right away, no prefixes.

						START OF DIFF
						%s
						END OF DIFF
					`, diff),
				},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	return res.Choices[0].Message.Content
}

func main() {
	cfg := loadConfig()
	oaiClient := openai.NewClient(cfg.OpenAI.Token)
	app := &cli.App{
		Name:  "chglog",
		Usage: "chglog",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			diff := getDiff()
			fmt.Printf("Commit message:\n%s\n\n", commitMessage(oaiClient, diff))
			fmt.Printf("Changelog entry:\n%s\n", summarise(oaiClient, diff))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
