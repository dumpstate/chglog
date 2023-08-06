package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const CHANGELOG_TEMPLATE = `# Changelog

## %s

%s
`

func getToday() string {
	ts := time.Now()
	return fmt.Sprintf("%d-%02d-%02d", ts.Year(), ts.Month(), ts.Day())
}

func parseAndAppend(message string, changelogFile string) {
	res := []string{}
	f, err := os.Open(changelogFile)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	appended := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "## ") {
			res = append(
				res,
				fmt.Sprintf("## %s", getToday()),
				"",
				message,
				"",
			)
			appended = true
		}

		res = append(res, line)
	}

	if !appended {
		res = append(
			res,
			fmt.Sprintf("## %s", getToday()),
			"",
			message,
			"",
		)
	}

	f.Close()

	f, err = os.Create(changelogFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for _, line := range res {
		f.WriteString(line + "\n")
	}
}

func appendChangelog(message string, changelogFile string) {
	if _, err := os.Stat(changelogFile); err == nil {
		parseAndAppend(message, changelogFile)
	} else {
		f, err := os.Create(changelogFile)
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatal(err)
			}
		}()

		if err != nil {
			log.Fatal(err)
		}

		f.WriteString(fmt.Sprintf(CHANGELOG_TEMPLATE, getToday(), message))
	}
}
