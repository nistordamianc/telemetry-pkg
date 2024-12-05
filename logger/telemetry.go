package logger

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

const (
	driverTypeCLI           = "CLI"
	driverTypeJSONFile      = "JSON"
	driverTypePlainTextFile = "PLAIN"
)

const LogFormat = "[%s] - [%s] %s %s\n"

const (
	levelError   = "ERROR"
	levelWarning = "WARNING"
	levelInfo    = "INFO"
)

type Metadata struct {
	Key   string
	Value string
}

type Logger interface {
	log(log Log)
	LogE(message string, metadata []Metadata)
	LogW(message string, metadata []Metadata)
	LogI(message string, metadata []Metadata)

	Flush()
}

type baseLogger struct {
	logFunc func(log Log)
}

func (l *baseLogger) LogE(message string, metadata []Metadata) {
	l.logFunc(Log{
		Timestamp: time.Now().Format(time.RFC3339Nano),
		Level:     levelError,
		Message:   message,
		Metadata:  metadata,
	})
}

func (l *baseLogger) LogW(message string, metadata []Metadata) {
	l.logFunc(Log{
		Timestamp: time.Now().Format(time.RFC3339Nano),
		Level:     levelError,
		Message:   message,
		Metadata:  metadata,
	})
}

func (l *baseLogger) LogI(message string, metadata []Metadata) {
	l.logFunc(Log{
		Timestamp: time.Now().Format(time.RFC3339Nano),
		Level:     levelInfo,
		Message:   message,
		Metadata:  metadata,
	})
}

type Log struct {
	Timestamp string
	Level     string
	Message   string
	Metadata  []Metadata
}

func NewLogger() Logger {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("ini")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	driverType := viper.GetString("default.driver_type")
	logsStorage := viper.GetString("default.logs_storage_location")

	switch driverType {
	case driverTypeCLI:
		return NewCLILogger()
	case driverTypeJSONFile:
		if logsStorage == "" {
			logsStorage = "logs.json"
		}
		return NewJSONLogger(logsStorage)
	case driverTypePlainTextFile:
		if logsStorage == "" {
			logsStorage = "logs.txt"
		}
		return NewPlainLogger(logsStorage)
	default:
		logger := NewCLILogger()

		logger.LogE(fmt.Sprintf("Driver type %s is not supported, defaulting to CLI logger.", driverType), nil)

		return logger
	}
}
