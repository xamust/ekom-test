package entity

import (
	"time"

	"github.com/google/uuid"
)

type Common struct {
	UUID      uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
