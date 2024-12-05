package logger

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCLILogger_Log(t *testing.T) {
	logger := NewCLILogger()
	metadata := []Metadata{
		{Key: "user", Value: "test_user"},
	}

	originalStdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	logger.LogI("Test message", metadata)

	w.Close()
	os.Stdout = originalStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	capturedOutput := buf.String()

	assert.Contains(t, capturedOutput, "INFO")
	assert.Contains(t, capturedOutput, "Test message")
	assert.Contains(t, capturedOutput, "test_user")
}
