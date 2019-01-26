package model

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	ID        uuid.UUID   `json:"id"`
	CreatedAt time.Timer  `json:"created_at" xorm:"created"`
	DeletedAt *time.Timer `json:"deleted_at" xorm:"deleted"`
	UpdatedAt *time.Timer `json:"deleted_at" xorm:"updated"`
	Version   int         `json:"version" xorm:"version"`
}
