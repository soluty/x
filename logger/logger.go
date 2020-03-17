package logger

import (
	"context"
	"logur.dev/logur"
)

type ContextExtractor interface {
	Extract(ctx context.Context) map[string]interface{}
}

type Logger interface {
	Trace(msg string, fields ...map[string]interface{})
	Debug(msg string, fields ...map[string]interface{})
	Info(msg string, fields ...map[string]interface{})
	Warn(msg string, fields ...map[string]interface{})
	Error(msg string, fields ...map[string]interface{})
	WithFields(fields map[string]interface{}) Logger
	WithContext(ctx context.Context) Logger
}

type Level uint8

const (
	DisableLevel Level = iota
	TraceLevel
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
)

var defaultLogger Logger

func init() {
	defaultLogger = NewNoopLogger()
}

func SetLogger(l Logger) {
	defaultLogger = l
}

func levelTrans(l Level) logur.Level {
	switch l {
	case TraceLevel:
		return logur.Trace
	case DebugLevel:
		return logur.Debug
	case InfoLevel:
		return logur.Info
	case WarnLevel:
		return logur.Warn
	case ErrorLevel:
		return logur.Error
	}
	return 999
}

func LevelEnable(l Level) bool {
	if e, ok := defaultLogger.(logur.LevelEnabler); ok {
		return e.LevelEnabled(levelTrans(l))
	} else {
		return false
	}
}

func Debug(msg string, fields ...map[string]interface{}) {
	defaultLogger.Debug(msg, fields...)
}
func Info(msg string, fields ...map[string]interface{}) {
	defaultLogger.Info(msg, fields...)
}
func Warn(msg string, fields ...map[string]interface{}) {
	defaultLogger.Warn(msg, fields...)
}
func Error(msg string, fields ...map[string]interface{}) {
	defaultLogger.Error(msg, fields...)
}
func Trace(msg string, fields ...map[string]interface{}) {
	defaultLogger.Trace(msg, fields...)
}
