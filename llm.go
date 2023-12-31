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

func changelogEntryTemplate(branch string, diff string) string {
	return fmt.Sprintf(`Given the provided context of a git change,
write a changelog entry. Changelog should follow GNU guidelines.

The result should be a list of changes, split into following buckets:
- 'Added' - for new features,
- 'Changed' - for changes in existing functionality,
- 'Deprecated' - for soon-to-be removed features,
- 'Removed' - for now removed features,
- 'Fixed' - for any bug fixes,
- 'Security' - in case of vulnerabilities.

Each bucket should contain a list of changes and each entry should
be a short and concise sentend.

The result should be formatted as a markdown list, with each bucket
being a markdown header (with '###'), e.g.:

START OF AN EXAMPLE
### Added

- Arabic translation
- Russian translation

### Changed

- Upgrade dependencies

### Removed

- Unused normalize.css file
END OF AN EXAMPLE

Do not prefix or suffix your answer, provide it right away.

BRANCH NAME: %s

START OF DIFF
%s
END OF DIFF`, branch, diff)
}

func openaiModel(cfg *Config) string {
	switch cfg.OpenAI.Model {
	case Gpt35Turbo:
		return openai.GPT3Dot5Turbo
	case Gpt4:
		return openai.GPT4
	default:
		return openai.GPT4
	}
}

func genCommitMessage(
	oaiClient *openai.Client,
	cfg *Config,
	branch string,
	diff string,
) string {
	res, err := oaiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openaiModel(cfg),
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

func genChangelogEntry(
	oaiClient *openai.Client,
	cfg *Config,
	branch string,
	diff string,
) string {
	res, err := oaiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openaiModel(cfg),
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: SystemMessage,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: changelogEntryTemplate(branch, diff),
				},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(res.Choices[0].Message.Content)
}
