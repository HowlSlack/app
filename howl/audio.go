package howl

import (
	"io"
	"log"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
)

// InitAudio inits the audio speaker
func InitAudio() {
	// the sample rate is set to 12000 as format.SampleRate does
	// not return an accurate value
	err := speaker.Init(12000, 12000)
	if err != nil {
		log.Fatal("audio: could not init speaker")
		return
	}
}

// PlayAudioStream will output a stream to device speakers
func PlayAudioStream(as io.ReadCloser) {
	// we decode the ogg stream
	s, _, _ := vorbis.Decode(as)

	// channel, which will signal the end of the playback.
	playing := make(chan struct{})

	// play the stream on the speaker
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		// call back when the stream ended
		close(playing)
	})))

	// wait for the end of the stream
	<-playing
}
