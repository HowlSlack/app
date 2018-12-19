package howl

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/nlopes/slack"
)

var (
	api *slack.Client
	rtm *slack.RTM

	// OnMessage will receive slack messages
	OnMessage chan string

	// UserIDToDisplayName we will store user display name
	// to avoid requesting the API each time there is a user mention
	UserIDToDisplayName map[string]string
	reUserIDs           *regexp.Regexp
)

// InitSlack will init the connection to slack
func InitSlack() {
	api = slack.New(config.SlackToken)

	rtm = api.NewRTM()
	go rtm.ManageConnection()

	OnMessage = make(chan string)
	UserIDToDisplayName = make(map[string]string)
	reUserIDs = regexp.MustCompile(`<@(\S+)>`)
}

// ListenForEvents will wait for events
// on the RTM api
func ListenForEvents() {
	// wait for events
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {

		// just connected
		case *slack.ConnectedEvent:
			log.Print("slack: connected")

		// received a message
		case *slack.MessageEvent:
			if ev.Type == "message" {

				message := ev.Text

				// replace all userIDs with their display names
				matches := reUserIDs.FindAllStringSubmatch(message, -1)
				for _, match := range matches {
					// string_to_replace = match[0]
					userID := match[1]

					// retrieve its display name
					username, present := UserIDToDisplayName[userID]

					if !present {
						user, err := api.GetUserInfo(userID)
						if err != nil {
							username = ""
						} else {
							username = user.Profile.DisplayName
						}
					}

					reUsername := regexp.MustCompile("<@" + userID + ">")
					message = reUsername.ReplaceAllString(message, username)
				}

				log.Printf("slack: message: %v\n", message)
				OnMessage <- message
			}

			// errors
		case *slack.RTMError:
			fmt.Fprintf(os.Stderr, "slack: RTM err: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			log.Fatal("slack: invalid credentials")
			return

		default:
			// Ignore other events..
		}
	}
}
