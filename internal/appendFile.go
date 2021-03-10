package internal

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

func DefaultFileName() string {
	return time.Now().Format("2006-01-02") + "-events.log"
}

func parsePath(fileName string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir

	if fileName == "~" {
		fileName = dir
	} else if strings.HasPrefix(fileName, "~/") {
		fileName = filepath.Join(dir, fileName[2:])
	}

	return fileName
}

func AppendFile(fileName string, event Event) {
	pfn := parsePath(fileName)
	path := parsePath(filepath.Dir(fileName))

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm|0777)
		if err != nil {
			log.Println(err)
		}
	}

	f, err := os.OpenFile(pfn,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm|0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(event.String() + "\n"); err != nil {
		log.Println(err)
	}
}
