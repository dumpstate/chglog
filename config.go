package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/urfave/cli/v2"
)

const (
	Gpt35Turbo string = "gpt-3.5-turbo"
	Gpt4       string = "gpt-4"
)

type OpenAIConfig struct {
	ApiKey string `json:"apiKey"`
	Model  string `json:"model"`
}

type Config struct {
	OpenAI OpenAIConfig `json:"openai"`
	Ignore []string     `json:"ignore"`
}

func validateModel(model string) {
	if (model != Gpt35Turbo) && (model != Gpt4) {
		log.Fatalf("Invalid OpenAI model: %s", model)
	}
}

func cfgPath() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	dir = path.Join(dir, "chglog")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	return path.Join(dir, "config.json")
}

func writeConfig(path string, cfg *Config) {
	content, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(path, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func readConfig() *Config {
	content, err := ioutil.ReadFile(cfgPath())
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	err = json.Unmarshal(content, &cfg)
	validateModel(cfg.OpenAI.Model)
	return &cfg
}

func initConfig(ctx *cli.Context) error {
	path := cfgPath()
	_, err := os.Stat(path)
	if err == nil {
		fmt.Printf("Config already exists at: %s\nWould you like to overwrite? [Y/n] ", path)
		if readStdIn() != "Y" {
			fmt.Println("Aborting...")
			os.Exit(0)
		}
	}

	fmt.Printf("OpenAI API Key: ")
	openaiApiKey := readStdIn()
	fmt.Printf("OpenAI Model (gpt-3.5-turbo/gpt-4): ")
	openaiModel := readStdIn()
	validateModel(openaiModel)

	config := Config{
		OpenAI: OpenAIConfig{
			ApiKey: openaiApiKey,
			Model:  openaiModel,
		},
		Ignore: []string{},
	}
	writeConfig(path, &config)
	fmt.Printf("Config initialised: %s\n", path)
	return nil
}
