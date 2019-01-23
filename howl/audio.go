package howl

import (
	"io"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
)

// InitAudio inits the audio speaker
func InitAudio() {
	sr := beep.SampleRate(22050)
	err := speaker.Init(sr, sr.N(time.Second/2))
	if err != nil {
		log.Fatal("audio: could not init speaker")
		return
	}
}

// PlayAudioStream will output a stream to device speakers
func PlayAudioStream(as io.ReadCloser, format string) {
	// we decode the ogg stream
	var s beep.StreamSeekCloser
	var f beep.Format

	if format == "vorbis" {
		s, f, _ = vorbis.Decode(as)
	} else if format == "mp3" {
		s, f, _ = mp3.Decode(as)
	}

	// channel, which will signal the end of the playback.
	playing := make(chan struct{})

	if format == "mp3" {
		// play the stream on the speaker
		speaker.Play(beep.Seq(beep.Resample(9, f.SampleRate, beep.SampleRate(22050), s), beep.Callback(func() {
			// call back when the stream ended
			close(playing)
		})))
	}

	if format == "vorbis" {
		// play the stream on the speaker
		speaker.Play(beep.Seq(beep.Resample(9, f.SampleRate, beep.SampleRate(44000), s), beep.Callback(func() {
			// call back when the stream ended
			close(playing)
		})))
	}

	// wait for the end of the stream
	<-playing
}
