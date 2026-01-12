package infra

import (
	"time"

	"github.com/google/uuid"
)

type TimestampMixin struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type UUIDMixin struct {
	ID uuid.UUID `json:"id"`
}

type VarCharIDMixin struct {
	ID string `json:"id"`
}
