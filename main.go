package main

import (
	"github.com/tetrash/event-recorder/cmd/events"
	"os"
)

func main() {
	switch os.Args[1] {
		case "add":
			args := os.Args[2:]
			events.Add(args)
		default:
			args := os.Args[1:]
			events.Add(args)
	}
}
