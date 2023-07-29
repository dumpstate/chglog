package main

import (
	"fmt"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
)

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
			{
				Name:    "commit",
				Aliases: []string{"c"},
				Usage:   "generate a commit message",
				Action: func(ctx *cli.Context) error {
					return commit(oaiClient)
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			diff := getCurrentDiff()
			fmt.Printf("Changelog entry:\n%s\n", genSummary(oaiClient, diff))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
