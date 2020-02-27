package logger

import (
	"github.com/rogpeppe/go-internal/lockedfile"
	"log"
	"os"
	"time"
)

const (
	logfile = "../../tools/responsetime.log"
)

var t time.Time

// Start keeps track of the start time of a process request
func Start() {
	t = time.Now()
}

// End keeps track of the end time of a process request
func End(task string) {
	f, _ := lockedfile.OpenFile(logfile, os.O_APPEND|os.O_WRONLY, 0666)
	defer f.Close()

	log.SetOutput(f)

	t2 := time.Since(t).Nanoseconds()
	log.Println(task, t2)
}
