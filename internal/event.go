package internal

import (
	"log"
	"strings"
	"time"
)

type Event struct {
	timestamp time.Time
	body string
	separator string
	dateFormat string
}

type event interface {
	String() string
	ParseEvent(event string) Event
}

func CreateEvent(body string, timeStamp time.Time) Event {
	return Event{
		timestamp: timeStamp,
		body: body,
		separator: " = ",
		dateFormat: time.RFC3339,
	}
}

func (e Event) String() string {
	return e.timestamp.Format(e.dateFormat) + e.separator + e.body
}

func (e Event) ParseEvent(event string) Event {
	res := strings.SplitN(event, e.separator, 1)
	timestamp, err := time.Parse(e.dateFormat, res[0])
	if err != nil {
		log.Println(err)
	}
	return CreateEvent(res[1], timestamp)
}

