package logger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlainLogger_Log(t *testing.T) {
	fileLocation := "test_plain.json"

	logger := NewPlainLogger(fileLocation)

	defer os.Remove(fileLocation)
	metadata := []Metadata{
		{Key: "user", Value: "test_user"},
	}

	logger.LogI("Test message", metadata)
	logger.Flush()

	fileContent, err := os.ReadFile(fileLocation)
	require.NoError(t, err)

	fileContentString := string(fileContent)

	assert.Contains(t, fileContentString, "INFO")
	assert.Contains(t, fileContentString, "Test message")
	assert.Contains(t, fileContentString, "test_user")
}
