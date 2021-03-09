package internal

import (
	"log"
	"os"
)

func AppendFile(fileName string, event Event) {
	f, err := os.OpenFile(fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(event.String() + "\n"); err != nil {
		log.Println(err)
	}
}
