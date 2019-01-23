package main

import (
	"os"

	"github.com/HowlSlack/app/howl"
)

func main() {
	// load config
	howl.LoadConfig()
	// init audio speakers
	howl.InitAudio()
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
				f, _ := os.Open("audio_effects/cha_ching.mp3")
				howl.PlayAudioStream(f, "mp3")
				f.Close()
				howl.PlayAudioStream(as, "vorbis")
			}
		}
	}
}
