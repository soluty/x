package logger

import (
	"context"
	"os"
)

func NewNoopLogger() Logger {
	return &emptyLogger{}
}

type emptyLogger struct {
}

func (e *emptyLogger) Trace(msg string, fields ...map[string]interface{}) {
}

func (e *emptyLogger) Debug(msg string, fields ...map[string]interface{}) {
}

func (e *emptyLogger) Info(msg string, fields ...map[string]interface{}) {
}

func (e *emptyLogger) Warn(msg string, fields ...map[string]interface{}) {
}

func (e *emptyLogger) Error(msg string, fields ...map[string]interface{}) {
}

func (e *emptyLogger) Fatal(msg string, fields ...map[string]interface{}) {
	os.Exit(1)
}

func (e *emptyLogger) WithFields(fields map[string]interface{}) Logger {
	return e
}

func (e *emptyLogger) WithContext(ctx context.Context) Logger {
	return e
}

var _ Logger = &emptyLogger{}
