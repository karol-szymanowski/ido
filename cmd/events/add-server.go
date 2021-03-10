package events

import (
	"encoding/json"
	"fmt"
	"github.com/tetrash/ido/internal"
	"net/http"
	"time"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	body := &internal.EventRequest{}
	err := d.Decode(body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	event := internal.CreateEvent(body.Body, time.Now(), &body.Tag)
	file := internal.DefaultFilePath() + internal.DefaultFileName()

	internal.AppendFile(file, event)
	_, err = fmt.Fprintf(w, "Ok")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
