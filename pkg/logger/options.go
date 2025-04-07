package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

type Option func(*Options)

type Options struct {
	output io.Writer
	level  slog.Level
}

func NewOptions(opts ...Option) Options {
	options := Options{
		level:  slog.LevelDebug,
		output: os.Stdout,
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

func WithOutput(output io.Writer) Option {
	return func(o *Options) {
		o.output = output
	}
}

func WithLevel(level string) Option {
	parseLevel, err := ParseLevel(level)
	if err != nil {
		panic(err)
	}
	return func(o *Options) {
		o.level = parseLevel
	}
}

func ParseLevel(level string) (slog.Level, error) {
	level = strings.TrimSpace(strings.ToLower(level))
	switch level {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return 0, fmt.Errorf("unknown level: %s", level)
	}
}
