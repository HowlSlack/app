package howl

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
)

var (
	sess *session.Session
	svc  *polly.Polly

	format string
	voice  string
)

// InitPolly manually inits the AWS Polly service
func InitPolly() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc = polly.New(sess)

	format = "ogg_vorbis" // mp3, pcm
	voice = "Lea"         // fr
}

// StringToAudioStream transforms a text into
// a readable audio stream
func StringToAudioStream(text string) io.ReadCloser {
	input := &polly.SynthesizeSpeechInput{
		Text:         aws.String(text),
		OutputFormat: aws.String(format),
		VoiceId:      aws.String(voice),
		SampleRate:   aws.String("22050"),
	}

	output, err := svc.SynthesizeSpeech(input)
	if err != nil {
		fmt.Printf("s3: polly failed for %s\n", text)
		return nil
	}

	return output.AudioStream
}
