package events

import (
	"flag"
	"github.com/tetrash/event-recorder/internal"
	"strings"
	"time"
)

func Add(subcommand string, args []string) {
	var file string
	var tag string

	defaultFileName := internal.DefaultFilePath() + internal.DefaultFileName()
	defaultTag := "cli"

	addCommand := flag.NewFlagSet(subcommand, flag.ExitOnError)
	addCommand.StringVar(&file, "file", defaultFileName, "File path to log events, eg.: ~/events.log")
	addCommand.StringVar(&file, "f", defaultFileName, "File path to log events, eg.: ~/events.log")
	addCommand.StringVar(&tag, "tag", defaultTag, "Event tag, eg.: cli")
	addCommand.StringVar(&tag, "t", defaultTag, "Event tag, eg.: cli")

	addCommand.Parse(args)

	eventBody := strings.Join(args, " ")
	event := internal.CreateEvent(eventBody, time.Now(), &tag)
	internal.AppendFile(file, event)
}
