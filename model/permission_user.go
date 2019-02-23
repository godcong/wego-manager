package model

import (
	"golang.org/x/xerrors"
)

// PermissionUser ...
type PermissionUser struct {
	Model        `xorm:"-"`
	PermissionID string `json:"permission_id" xorm:"permission_id notnull unique(permission_user) uuid"`
	UserID       string `json:"user_id" xorm:"user_id notnull unique(permission_user) uuid"`
}

// Relate ...
func (obj *PermissionUser) Relate() (*Permission, *User, error) {
	var info struct {
		Permission Permission `xorm:"extends"`
		User       User       `xorm:"extends"`
	}
	session := DB().Table(&info.Permission).Select("permission.*, user.*").
		Join("left", obj, "permission_user.permission_id = permission.id").
		Join("left", &info.User, "permission_user.user_id = user.id")

	if obj.UserID != "" {
		session = session.Where("user.id = ? ", obj.UserID)
	}
	if obj.PermissionID != "" {
		session = session.Where("permission.id = ? ", obj.PermissionID)
	}
	b, err := session.Get(&info)
	if err != nil {
		return nil, nil, xerrors.Errorf("relate: %w", err)
	}
	if !b {
		return nil, nil, xerrors.Errorf("permission user not found")
	}
	return &info.Permission, &info.User, nil
}
