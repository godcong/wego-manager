package model

import "github.com/google/uuid"

// PermissionRole ...
type PermissionRole struct {
	PermissionID uuid.UUID `json:"permission_id" xorm:"permission_id uuid"`
	RoleID       uuid.UUID `json:"role_id" xorm:"role_id uuid"`
}
