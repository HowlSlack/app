package main

import (
	"github.com/HowlSlack/app/howl"
)

func main() {
	// load config
	howl.LoadConfig()
	// init aws polly services
	howl.InitPolly()
	// connect to slack
	howl.InitSlack()

	// listen to slack events
	go howl.ListenForEvents()

	for {
		select {
		// we just received a message on our channel
		case message := <-howl.OnMessage:
			// synthesize it
			as := howl.StringToAudioStream(message)
			if as != nil {
				// play it on speakers
				howl.PlayAudioStream(as)
			}
		}
	}
}
