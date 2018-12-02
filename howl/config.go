package howl

import (
	"fmt"
	"os"

	"github.com/tkanos/gonfig"
)

// Configuration handles the configuration for howl
type Configuration struct {
	// SlackToken is the bot token you should generate
	// in https://YOUR_SUBDOMAIN.slack.com/apps/manage/custom-integrations
	SlackToken string
}

var (
	config Configuration
)

// LoadConfig loads config.json which
// has all the auth tokens
func LoadConfig() {
	config = Configuration{}
	err := gonfig.GetConf("config.json", &config)
	if err != nil {
		fmt.Printf("err: could not load config.json (invalid syntax?)\n")
		os.Exit(0)
	}

	// check for errors in json
	if config.SlackToken == "" {
		fmt.Printf("err: put your token in config.json\n")
		os.Exit(0)
	}
}

// GetSlackToken return the slack token
func GetSlackToken() string {
	return config.SlackToken
}
