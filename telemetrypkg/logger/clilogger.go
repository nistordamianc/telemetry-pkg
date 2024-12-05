package logger

import (
	"fmt"
)

type CLILogger struct {
	baseLogger
}

func (l *CLILogger) log(log Log) {
	fmt.Printf(LogFormat, log.Timestamp, log.Level, log.Message, log.Metadata)
}

func (l *CLILogger) Flush() {
}

func NewCLILogger() *CLILogger {
	logger := &CLILogger{}
	logger.logFunc = logger.log

	return logger
}
