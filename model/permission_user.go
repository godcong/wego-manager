package model

import (
	"golang.org/x/xerrors"
)

// PermissionUser ...
type PermissionUser struct {
	Model        `xorm:"-"`
	PermissionID string `json:"permission_id" xorm:"permission_id notnull uuid"`
	UserID       string `json:"user_id" xorm:"user_id notnull uuid"`
}

// Relate ...
func (obj *PermissionUser) Relate() (*Permission, *User, error) {
	var info []struct {
		Permission *Permission `xorm:"extends"`
		User       *User       `xorm:"extends"`
	}
	session := DB().Table(&Permission{}).Select("permission.*, user.*").
		Join("left", obj, "permission_user.permission_id = permission.id").
		Join("left", &User{}, "permission_user.user_id = user.id")

	if obj.UserID != "" {
		session = session.Where("user.id = ? ", obj.UserID)
	}
	if obj.PermissionID != "" {
		session = session.Where("permission.id = ? ", obj.PermissionID)
	}
	i, err := session.FindAndCount(&info)
	if err != nil {
		return nil, nil, xerrors.Errorf("relate: %w", err)
	}
	if i > 1 {
		return nil, nil, xerrors.Errorf("count %d > 1 ", i)
	}
	return (info)[0].Permission, (info)[0].User, nil
}
