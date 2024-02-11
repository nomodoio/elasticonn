// Package main ...
package main

import (
	"flag"

	"github.com/davecgh/go-spew/spew"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/pterm/pterm"
	"github.com/rs/zerolog/log"
)

func main() {
	url := ""
	userName := ""
	password := ""

	flag.StringVar(&url, "url", "", "url as a connection string")
	flag.StringVar(&userName, "user", "", "elastic user name")
	flag.StringVar(&password, "password", "", "password for elastic user")
	flag.Parse()

	if url == "" {
		url, _ = pterm.DefaultInteractiveTextInput.
			WithDefaultText("elastic url").
			Show()
	}

	if userName == "" {
		userName, _ = pterm.DefaultInteractiveTextInput.
			WithDefaultText("elastic user name").
			Show()
	}

	if password == "" {
		password, _ = pterm.DefaultInteractiveTextInput.
			WithDefaultText("password for user " + userName).
			Show()
	}

	cfg := elasticsearch.Config{
		Addresses: []string{url},
		Username:  userName,
		Password:  password,
	}
	spew.Dump(cfg)

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal().Err(err).Caller().Send()
	}

	client.Ping()

	infores, err := client.Info()
	if err != nil {
		log.Fatal().Err(err).Caller().Send()
	}

	log.Info().Interface("info", infores).Send()
}
