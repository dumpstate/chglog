package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	openai "github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
)

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
	cfg := readConfig()
	oaiClient := openai.NewClient(cfg.OpenAI.ApiKey)
	app := &cli.App{
		Name:  "chglog",
		Usage: "chglog",
		Flags: []cli.Flag{},
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "initialise chglog config file",
				Action:  initConfig,
			},
		},
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
