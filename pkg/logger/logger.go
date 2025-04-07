package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/xamust/ekom-test/internal/interfaces"
)

var _ interfaces.Logger = (*Logger)(nil)

type Logger struct {
	logger *slog.Logger
	opts   Options
}

func NewLogger(o ...Option) *Logger {
	l := &Logger{
		opts: NewOptions(o...),
	}
	handler := slog.Handler(slog.NewJSONHandler(
		l.opts.output,
		&slog.HandlerOptions{Level: l.opts.level}),
	)
	// set by Default logger
	slog.SetDefault(slog.New(NewHandlerMiddleware(handler)))
	return &Logger{
		logger: slog.Default(),
	}
}

func (l *Logger) Debug(message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case string:
		l.logger.Debug(msg, args...)
	case error:
		l.logger.Debug(msg.Error(), args...)
	default:
		l.logger.Debug(fmt.Sprintf("message %v has unknown type %v", message, msg), args...)
	}
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.logger.Info(message, args...)
}

func (l *Logger) Warn(message string, args ...interface{}) {
	l.logger.Warn(message, args...)
}

func (l *Logger) Error(message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case string:
		l.logger.Error(msg, args...)
	case error:
		l.logger.Error(msg.Error(), args...)
	default:
		l.logger.Error(fmt.Sprintf("message %v has unknown type %v", message, msg), args...)
	}
}

func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.Error(message, args...)
	os.Exit(1)
}
