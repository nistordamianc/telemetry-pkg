package main

import (
	telemetrypkg "telemetry-pkg/logger"
)

func main() {
	logger := telemetrypkg.NewLogger()
	transactionalLogger := telemetrypkg.NewTransactionalLogger(logger)

	logger.LogW("Hello, vtm!", []telemetrypkg.Metadata{{Key: "testey1", Value: "testval1"}, {Key: "key1", Value: "value1"}})

	transactionalLogger.LogW("log1", []telemetrypkg.Metadata{{Key: "key", Value: "value"}, {Key: "key1", Value: "value1"}})
	transactionalLogger.LogW("log2", []telemetrypkg.Metadata{{Key: "key", Value: "value"}, {Key: "key2", Value: "value2"}})
	transactionalLogger.LogW("log3", []telemetrypkg.Metadata{{Key: "key", Value: "value"}, {Key: "key3", Value: "value3"}})

	defer logger.Flush()
}
