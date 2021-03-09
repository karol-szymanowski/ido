package events

import (
	"flag"
	"github.com/tetrash/event-recorder/internal"
	"strings"
	"time"
)

func Add(args []string) {
	file := flag.String("path", "events.log", "File path to log events. Default: events.log")

	flag.Parse()

	eventBody := strings.Join(args, " ")
	event := internal.CreateEvent(eventBody, time.Now())
	internal.AppendFile(*file, event)
}
