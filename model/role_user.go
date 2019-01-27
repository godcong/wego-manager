package model

import "github.com/google/uuid"

// RoleUser ...
type RoleUser struct {
	Model  `xorm:"-"`
	RoleID uuid.UUID `json:"role_id" xorm:"role_id uuid"`
	UserID uuid.UUID `json:"user_id" xorm:"user_id uuid"`
}
