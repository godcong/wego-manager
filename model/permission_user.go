package model

import "github.com/google/uuid"

// PermissionUser ...
type PermissionUser struct {
	Model        `xorm:"-"`
	PermissionID uuid.UUID `json:"permission_id" xorm:"permission_id uuid"`
	UserID       uuid.UUID `json:"user_id" xorm:"user_id uuid"`
}

// Relate ...
func (obj *PermissionUser) Relate(permission *Permission, user *User) error {
	return DB().Table(permission).
		Join("left", obj, "permission_user.permission_id = permission.id").
		Join("left", user, "permission_user.user_id = user.id").
		Find(permission, user)
}
