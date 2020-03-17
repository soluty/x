package logger

import (
	"context"
	"logur.dev/logur"
)

type unionLogger []Logger

func (u unionLogger) Trace(msg string, fields ...map[string]interface{}) {
	for _, value := range u {
		value.Trace(msg, fields...)
	}
}

func (u unionLogger) Debug(msg string, fields ...map[string]interface{}) {
	for _, value := range u {
		value.Debug(msg, fields...)
	}
}

func (u unionLogger) Info(msg string, fields ...map[string]interface{}) {
	for _, value := range u {
		value.Info(msg, fields...)
	}
}

func (u unionLogger) Warn(msg string, fields ...map[string]interface{}) {
	for _, value := range u {
		value.Warn(msg, fields...)
	}
}

func (u unionLogger) Error(msg string, fields ...map[string]interface{}) {
	for _, value := range u {
		value.Error(msg, fields...)
	}
}

func (u unionLogger) WithFields(fields map[string]interface{}) Logger {
	for idx, value := range u {
		u[idx] = value.WithFields(fields)
	}
	return u
}

func (u unionLogger) WithContext(ctx context.Context) Logger {
	for idx, value := range u {
		u[idx] = value.WithContext(ctx)
	}
	return u
}

func (u unionLogger) LevelEnabled(l logur.Level) bool {
	for _, value := range u {
		if e, ok := value.(logur.LevelEnabler); ok {
			if e.LevelEnabled(l) {
				return true
			}
		}
	}
	return false
}

func NewUnionLogger(loggers ...Logger) Logger {
	return unionLogger(loggers)
}
