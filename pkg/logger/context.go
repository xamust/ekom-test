package logger

import "github.com/google/uuid"

type logCtx struct {
	UserUUID uuid.UUID
}
