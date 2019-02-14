package model

import "golang.org/x/xerrors"

// PermissionRole ...
type PermissionRole struct {
	Model        `xorm:"-"`
	PermissionID string      `json:"-" xorm:"permission_id notnull uuid"`
	Permission   *Permission `xorm:"-"`
	RoleID       string      `json:"-" xorm:"role_id notnull uuid"`
	Role         *Role       `xorm:"-"`
}

// Relate ...
func (obj *PermissionRole) Relate() (*Permission, *Role, error) {
	var info []struct {
		Permission *Permission `xorm:"extends"`
		Role       *Role       `xorm:"extends"`
	}
	session := DB().Table(&Permission{}).Select("permission.*, role.*").
		Join("left", obj, "permission_role.permission_id = permission.id").
		Join("left", &Role{}, "permission_role.role_id = role.id")

	if obj.RoleID != "" {
		session = session.Where("role.id = ? ", obj.RoleID)
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
	return (info)[0].Permission, (info)[0].Role, nil
}
