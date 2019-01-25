package model

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	ID        uuid.UUID   `json:"id"`
	CreatedAt time.Timer  `json:"created_at"`
	DeletedAt *time.Timer `json:"deleted_at"`
	Version   int         `json:"version"`
}
