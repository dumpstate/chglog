package main

import (
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
)

func main() {
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
			{
				Name:    "commit",
				Aliases: []string{"c"},
				Usage:   "generate a commit message",
				Action: func(ctx *cli.Context) error {
					cfg := readConfig()
					oaiClient := openai.NewClient(cfg.OpenAI.ApiKey)
					return commit(oaiClient, cfg)
				},
			},
			{
				Name: "changelog",
				Aliases: []string{"cl"},
				Usage: "generate a changelog entry",
				Action: func(ctx *cli.Context) error {
					cfg := readConfig()
					oaiClient := openai.NewClient(cfg.OpenAI.ApiKey)
					return changelogEntry(oaiClient, cfg)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
