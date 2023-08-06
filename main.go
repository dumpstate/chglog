package main

import (
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
)

func main() {
	var appendChangelog bool
	var changelogFile string

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
				Name:    "changelog",
				Aliases: []string{"cl"},
				Usage:   "generate a changelog entry",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "append",
						Usage:       "append to existing changelog file",
						Destination: &appendChangelog,
					},
					&cli.StringFlag{
						Name:        "target",
						Usage:       "target changelog file",
						Value:       "CHANGELOG.md",
						Destination: &changelogFile,
					},
				},
				Action: func(ctx *cli.Context) error {
					cfg := readConfig()
					oaiClient := openai.NewClient(cfg.OpenAI.ApiKey)
					return changelogEntry(oaiClient, cfg, appendChangelog, changelogFile)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
