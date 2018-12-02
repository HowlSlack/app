package howl

import (
	"fmt"
	"io"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
)

// PlayAudioStream will output a stream to device speakers
func PlayAudioStream(as io.ReadCloser) {
	// we decode the ogg stream
	s, format, _ := vorbis.Decode(as)

	// the sample rate is set to 12000 as s.SampleRate does
	// not return an accurate value
	err := speaker.Init(12000, format.SampleRate.N(time.Second))
	if err != nil {
		fmt.Println("audio: err: could not play audio stream")
		return
	}

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
