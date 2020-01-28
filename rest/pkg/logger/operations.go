package logger

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	operationsFile = "../../tools/operations.log"
)

var opLog []string
var f *os.File
var w *bufio.Writer

// Log maintains the mainpulation of the dbs in memory and flushed to file at specific time interval
func Log(task, operation string) {
	l := fmt.Sprintf("%s%s\n", task, operation)
	opLog = append(opLog, l)
}

// LogToFile flushes in memory data to file
func LogToFile() {
	nextTime := time.Now().Truncate(time.Minute)
	nextTime = nextTime.Add(time.Minute)
	time.Sleep(time.Until(nextTime))

	f, _ = os.OpenFile(operationsFile, os.O_APPEND|os.O_WRONLY, 0666)
	w = bufio.NewWriter(f)

	for _, record := range opLog {
		_, err := w.WriteString(record)
		if err != nil {
		}
	}
	w.Flush()
	opLog = opLog[:0]

	go LogToFile()
}
