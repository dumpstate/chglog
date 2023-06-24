package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

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

func main() {
	app := &cli.App{
		Name:  "chglog",
		Usage: "chglog",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			fmt.Println("Genering changelog...")
			cfg := loadConfig()
			fmt.Println(cfg)
			// TODO
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
