package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/sashabaranov/go-openai"
)

const SystemMessage = `You're a programmer's assistant.
Your task is to help the programmer write nice and concise git
messages and changelog entries.`

func commitMessageTemplate(branch string, diff string) string {
	return fmt.Sprintf(`Given the provided context of a git change,
write a single line git commit message, max 120 characters.

Provide your answer right away, without any prefixes or suffixes.

BRANCH NAME: %s

START OF DIFF
%s
END OF DIFF`, branch, diff)
}

func genCommitMessage(
	oaiClient *openai.Client,
	branch string,
	diff string,
) string {
	res, err := oaiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: SystemMessage,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: commitMessageTemplate(branch, diff),
				},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(res.Choices[0].Message.Content)
}