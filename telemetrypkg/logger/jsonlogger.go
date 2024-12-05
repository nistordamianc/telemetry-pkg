package logger

import (
	"encoding/json"
	"os"
)

type JSONLogger struct {
	baseLogger
	fileLocation string
	buffer       []JSONLog
}

type JSONLog struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	Metadata  string `json:"metadata"`
}

func (l *JSONLogger) log(log Log) {
	jsonMetadata, _ := json.Marshal(log.Metadata)

	l.buffer = append(l.buffer, JSONLog{
		Timestamp: log.Timestamp,
		Level:     log.Level,
		Message:   log.Message,
		Metadata:  string(jsonMetadata),
	})
}

func NewJSONLogger(fileLocation string) *JSONLogger {
	logger := &JSONLogger{fileLocation: fileLocation}
	logger.logFunc = logger.log

	return logger
}

func (l *JSONLogger) Flush() {
	if len(l.buffer) > 0 {
		existingLogs := []JSONLog{}

		if fileData, err := os.ReadFile(l.fileLocation); err == nil {
			json.Unmarshal(fileData, &existingLogs)
		}

		updatedLogs := append(existingLogs, l.buffer...)

		logData, _ := json.MarshalIndent(updatedLogs, "", "  ")

		os.WriteFile(l.fileLocation, logData, 0o644)

		l.buffer = []JSONLog{}
	}
}
