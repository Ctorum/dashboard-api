package infra

import (
	"time"

	"github.com/google/uuid"
	"go.jetify.com/typeid/v2"
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

type TypeIDMixin struct {
	ID string `json:"id"`
}

func (id TypeIDMixin) ToString() (string, error) {
	tid, err := typeid.Parse(id.ID)
	if err != nil {
		return "", err
	}
	return tid.String(), nil
}
