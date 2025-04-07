package logger

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type keyType int

const key = keyType(0)

type HandlerMiddlware struct {
	next slog.Handler
}

func (h HandlerMiddlware) Enabled(ctx context.Context, level slog.Level) bool {
	return h.next.Enabled(ctx, level)
}

func (h HandlerMiddlware) Handle(ctx context.Context, rec slog.Record) error {
	if c, ok := ctx.Value(key).(logCtx); ok {
		if c.UserUUID != uuid.Nil {
			rec.Add("userUUID", c.UserUUID)
		}
	}
	return h.next.Handle(ctx, rec)
}

func (h HandlerMiddlware) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &HandlerMiddlware{next: h.next.WithAttrs(attrs)}
}

func (h HandlerMiddlware) WithGroup(name string) slog.Handler {
	return &HandlerMiddlware{next: h.next.WithGroup(name)}
}

func NewHandlerMiddleware(next slog.Handler) *HandlerMiddlware {
	return &HandlerMiddlware{next: next}
}
