package logger

import (
	"testing"
)

type MockLogger struct {
	Logs []Log
}

func (m *MockLogger) log(log Log) {
	m.Logs = append(m.Logs, log)
}

func (m *MockLogger) LogE(message string, metadata []Metadata) {
	m.log(Log{
		Timestamp: "mock-timestamp",
		Level:     levelError,
		Message:   message,
		Metadata:  metadata,
	})
}

func (m *MockLogger) LogW(message string, metadata []Metadata) {
	m.log(Log{
		Timestamp: "mock-timestamp",
		Level:     levelWarning,
		Message:   message,
		Metadata:  metadata,
	})
}

func (m *MockLogger) LogI(message string, metadata []Metadata) {
	m.log(Log{
		Timestamp: "mock-timestamp",
		Level:     levelInfo,
		Message:   message,
		Metadata:  metadata,
	})
}

func (m *MockLogger) Flush() {
	// No-op for mock
}

func TestTransactionalLogger_LogW(t *testing.T) {
	mockLogger := &MockLogger{}
	transactionalLogger := NewTransactionalLogger(mockLogger)

	message := "Test warning log"
	metadata := []Metadata{
		{Key: "key1", Value: "value1"},
	}

	transactionalLogger.LogW(message, metadata)

	if len(mockLogger.Logs) != 1 {
		t.Fatalf("Expected 1 log entry, got %d", len(mockLogger.Logs))
	}

	log := mockLogger.Logs[0]
	if log.Message != message {
		t.Errorf("Expected message '%s', got '%s'", message, log.Message)
	}

	foundTransactionID := false
	for _, meta := range log.Metadata {
		if meta.Key == "transaction_id" && meta.Value == transactionalLogger.TransactionID {
			foundTransactionID = true
			break
		}
	}

	if !foundTransactionID {
		t.Error("Transaction ID not found in log metadata")
	}
}

func TestTransactionalLogger_LogE(t *testing.T) {
	mockLogger := &MockLogger{}
	transactionalLogger := NewTransactionalLogger(mockLogger)

	message := "Test error log"
	metadata := []Metadata{
		{Key: "key2", Value: "value2"},
	}

	transactionalLogger.LogE(message, metadata)

	if len(mockLogger.Logs) != 1 {
		t.Fatalf("Expected 1 log entry, got %d", len(mockLogger.Logs))
	}

	log := mockLogger.Logs[0]
	if log.Message != message {
		t.Errorf("Expected message '%s', got '%s'", message, log.Message)
	}

	if log.Level != levelError {
		t.Errorf("Expected level '%s', got '%s'", levelError, log.Level)
	}

	foundTransactionID := false
	for _, meta := range log.Metadata {
		if meta.Key == "transaction_id" && meta.Value == transactionalLogger.TransactionID {
			foundTransactionID = true
			break
		}
	}

	if !foundTransactionID {
		t.Error("Transaction ID not found in log metadata")
	}
}

func TestTransactionalLogger_Flush(t *testing.T) {
	mockLogger := &MockLogger{}
	transactionalLogger := NewTransactionalLogger(mockLogger)

	transactionalLogger.Flush()

	t.Log("Flush executed without error")
}
