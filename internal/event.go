package internal

import (
	"regexp"
	"strings"
	"time"
)

type Event struct {
	timestamp time.Time
	body string
	separator string
	dateFormat string
	tag string
}

type EventRequest struct {
	Body string `json:"body"`
	Tag string `json:"tag"`
}

type event interface {
	String() string
	ParseEvent(event string) Event
}

func CreateEvent(body string, timeStamp time.Time, eventTag *string) Event {
	tag := "none"
	if eventTag != nil {
		tag = *eventTag
	}

	return Event{
		timestamp: timeStamp,
		body: body,
		separator: " = ",
		dateFormat: time.RFC3339,
		tag: tag,
	}
}

func (e Event) String() string {
	return "[" + e.tag + "] " + e.timestamp.Format(e.dateFormat) + e.separator + e.body
}

func (e Event) ParseEvent(event string) (Event, error) {
	res := strings.SplitN(event, e.separator, 1)

	tagExpression := regexp.MustCompile("\\[(.*?)]")
	tag := tagExpression.FindString(res[0])

	timestampExpression := regexp.MustCompile("([0-9]+)-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])[Tt]([01][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9]|60)(\\.[0-9]+)?(([Zz])|([+|\\-]([01][0-9]|2[0-3]):[0-5][0-9]))")
	timestamp, err := time.Parse(e.dateFormat, timestampExpression.FindString(res[0]))

	if err != nil {
		return Event{}, err
	}

	return CreateEvent(res[1], timestamp, &tag), nil
}

