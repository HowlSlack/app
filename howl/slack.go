package howl

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

var (
	api *slack.Client
	rtm *slack.RTM

	// OnMessage will receive slack messages
	OnMessage chan string
)

// InitSlack will init the connection to slack
func InitSlack() {
	api = slack.New(config.SlackToken)

	rtm = api.NewRTM()
	go rtm.ManageConnection()

	OnMessage = make(chan string)
}

// ListenForEvents will wait for events
// on the RTM api
func ListenForEvents() {
	// wait for events
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {

		// just connected
		case *slack.ConnectedEvent:
			fmt.Println("slack: connected")

		// received a message
		case *slack.MessageEvent:
			if ev.Type == "message" {
				fmt.Printf("slack: message: %v\n", ev.Text)
				OnMessage <- ev.Text
			}

		// errors
		case *slack.RTMError:
			fmt.Printf("slackk: RTM err: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Println("slack: invalid credentials")
			os.Exit(1)
			return

		default:
			// Ignore other events..
		}
	}
}
