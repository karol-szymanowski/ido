package main

import (
	"github.com/tetrash/event-recorder/cmd/events"
	"os"
)

func main() {
	args := os.Args[2:]
	switch os.Args[1] {
		case "add":
			events.Add("add", args)
		case "server":
			events.StartHttpServer("server", args)
		default:
			args := os.Args[1:]
			events.Add("", args)
	}
}
