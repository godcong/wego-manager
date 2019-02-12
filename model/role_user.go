package model

import "golang.org/x/xerrors"

// RoleUser ...
type RoleUser struct {
	Model  `xorm:"-"`
	RoleID string `json:"role_id" xorm:"role_id notnull uuid"`
	UserID string `json:"user_id" xorm:"user_id notnull uuid"`
}

// Relate ...
func (obj *RoleUser) Relate() (*Role, *User, error) {
	var info []struct {
		Role *Role `xorm:"extends"`
		User *User `xorm:"extends"`
	}
	session := DB().Table(&Role{}).Select("role.*, user.*").
		Join("left", obj, "role_user.role_id = role.id").
		Join("left", &User{}, "role_user.user_id = user.id")

	if obj.UserID != "" {
		session = session.Where("user.id = ? ", obj.UserID)
	}
	if obj.RoleID != "" {
		session = session.Where("role.id = ? ", obj.RoleID)
	}
	i, err := session.FindAndCount(&info)
	if err != nil {
		return nil, nil, xerrors.Errorf("relate: %w", err)
	}
	if i > 1 {
		return nil, nil, xerrors.Errorf("count %d > 1 ", i)
	}
	return (info)[0].Role, (info)[0].User, nil
}
