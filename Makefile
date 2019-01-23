# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=howl-slack
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_ARM=$(BINARY_NAME)_arm

all: build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/tkanos/gonfig
	$(GOGET) github.com/faiface/beep
	$(GOGET) github.com/aws/aws-sdk-go/aws
	$(GOGET) github.com/nlopes/slack
	$(GOGET) github.com/hajimehoshi/oto
	$(GOGET) github.com/jfreymuth/oggvorbis
	$(GOGET) github.com/pkg/errors
	$(GOGET) github.com/hajimehoshi/oto

build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
build-raspberry:
	CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 $(GOBUILD) -o $(BINARY_ARM) -v
