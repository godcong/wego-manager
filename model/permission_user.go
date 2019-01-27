package model

import "github.com/google/uuid"

// PermissionUser ...
type PermissionUser struct {
	Model        `xorm:"-"`
	PermissionID uuid.UUID `json:"permission_id" xorm:"permission_id uuid"`
	UserID       uuid.UUID `json:"user_id" xorm:"user_id uuid"`
}
