package utils

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	Info  *log.Logger
	Error *log.Logger
}

// console -> logs in console
// default: create log file from path
func NewLogger(path string) *Logger {
	var output io.Writer
	if path == "console" {
		output = os.Stdout
	} else {
		file, err := os.Create(path)
		if err != nil {
			log.Fatal("Error creating log file:", err)
		}
		output = file
	}

	return &Logger{
		Info:  log.New(output, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		Error: log.New(output, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
