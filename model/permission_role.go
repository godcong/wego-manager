package model

import "golang.org/x/xerrors"

// PermissionRole ...
type PermissionRole struct {
	Model        `xorm:"-"`
	PermissionID string      `json:"-" xorm:"permission_id notnull unique(permission_role) uuid"`
	Permission   *Permission `xorm:"-"`
	RoleID       string      `json:"-" xorm:"role_id notnull unique(permission_role) uuid"`
	Role         *Role       `xorm:"-"`
}

// Relate ...
func (obj *PermissionRole) Relate() (*Permission, *Role, error) {
	var info struct {
		Permission Permission `xorm:"extends"`
		Role       Role       `xorm:"extends"`
	}
	session := Table(&info.Permission).Select("permission.*, role.*").
		Join("left", obj, "permission_role.permission_id = permission.id").
		Join("left", &info.Role, "permission_role.role_id = role.id")

	if obj.RoleID != "" {
		session = session.Where("role.id = ? ", obj.RoleID)
	}
	if obj.PermissionID != "" {
		session = session.Where("permission.id = ? ", obj.PermissionID)
	}
	b, err := session.Get(&info)
	if err != nil {
		return nil, nil, xerrors.Errorf("relate: %w", err)
	}
	if !b {
		return nil, nil, xerrors.Errorf("permission role not found")
	}
	return &info.Permission, &info.Role, nil
}
