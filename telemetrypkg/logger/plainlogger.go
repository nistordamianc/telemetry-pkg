package logger

import (
	"fmt"
	"os"
)

type PlainLogger struct {
	baseLogger
	fileLocation string
	buffer       string
}

func (l *PlainLogger) log(log Log) {
	l.buffer += fmt.Sprintf(LogFormat, log.Timestamp, log.Level, log.Message, log.Metadata)
}

func (l *PlainLogger) Flush() {
	file, errOpen := os.OpenFile(l.fileLocation, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)

	if errOpen != nil {
		panic(errOpen)
	}

	defer file.Close()

	_, errWrite := file.WriteString(l.buffer)
	if errWrite != nil {
		panic(errWrite)
	}
}

func NewPlainLogger(fileLocation string) *PlainLogger {
	logger := &PlainLogger{fileLocation: fileLocation}
	logger.logFunc = logger.log

	return logger
}
