package model

import (
	"github.com/google/uuid"
)

// Model ...
type Model struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt int64     `json:"created_at" xorm:"created"`
	UpdatedAt int64     `json:"deleted_at" xorm:"updated"`
	DeletedAt *int64    `json:"deleted_at" xorm:"deleted"`
	Version   int       `json:"version" xorm:"version"`
}
