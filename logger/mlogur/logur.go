package mlogur

import (
	"context"
	"github.com/soluty/x/logger"
	"logur.dev/logur"
)

type Logger struct {
	logger       logur.Logger
	ctxExtractor logger.ContextExtractor
}

func NewLogger(logger logur.Logger, ctxExtractor ...logger.ContextExtractor) *Logger {
	l := &Logger{
		logger: logger,
	}
	if len(ctxExtractor) > 0 {
		l.ctxExtractor = ctxExtractor[0]
	}
	return l
}

func (l *Logger) Trace(msg string, fields ...map[string]interface{}) {
	l.logger.Trace(msg, fields...)
}
func (l *Logger) Debug(msg string, fields ...map[string]interface{}) {
	l.logger.Debug(msg, fields...)
}
func (l *Logger) Info(msg string, fields ...map[string]interface{}) {
	l.logger.Info(msg, fields...)
}
func (l *Logger) Warn(msg string, fields ...map[string]interface{}) {
	l.logger.Warn(msg, fields...)
}
func (l *Logger) Error(msg string, fields ...map[string]interface{}) {
	l.logger.Error(msg, fields...)
}

func (l *Logger) WithFields(fields map[string]interface{}) logger.Logger {
	return &Logger{
		logger:       logur.WithFields(l.logger, fields),
		ctxExtractor: l.ctxExtractor,
	}
}

func (l *Logger) WithContext(ctx context.Context) logger.Logger {
	if l.ctxExtractor == nil {
		return l
	}
	return l.WithFields(l.ctxExtractor.Extract(ctx))
}

func (l *Logger) LevelEnabled(level logur.Level) bool {
	if e, ok := l.logger.(logur.LevelEnabler); ok {
		return e.LevelEnabled(level)
	} else {
		return false
	}
}
