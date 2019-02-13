package model

import "golang.org/x/xerrors"

// Permission ...
type Permission struct {
	Model           `xorm:"extends" json:",inline"`
	Name            string `xorm:"name"`
	Slug            string `xorm:"slug"`
	Description     string `xorm:"description"`
	PermissionModel string `xorm:"permission_model"`
}

// NewPermission ...
func NewPermission(id string) *Permission {
	return &Permission{
		Model: Model{
			ID: id,
		},
	}
}

// Get ...
func (obj *Permission) Get() (bool, error) {
	return Get(nil, obj)
}

// Update ...
func (obj *Permission) Update(cols ...string) (int64, error) {
	return Update(nil, obj.ID, obj)
}

// Permissions ...
func (obj *Permission) Permissions() ([]*Permission, error) {
	var permissions []*Permission
	err := DB().Table(obj).Find(&permissions)
	if err != nil {
		return nil, xerrors.Errorf("find: %w", err)
	}
	return permissions, nil
}

// Roles ...
func (obj *Permission) Roles() ([]*Role, error) {
	var roles []*Role
	session := DB().Table(&Role{}).Select("role.*").
		Join("left", &PermissionRole{}, "permission_role.role_id = role.id")

	if obj.ID != "" {
		session = session.Where("permission_role.permission_id = ? ", obj.ID)
	}

	err := session.Find(&roles)
	if err != nil {
		return nil, xerrors.Errorf("relate: %w", err)
	}

	return roles, nil
}

// Users ...
func (obj *Permission) Users() ([]*User, error) {
	var users []*User
	session := DB().Table(User{}).Select("user.*").
		Join("left", &RoleUser{}, "permission_user.user_id = user.id")

	if obj.ID != "" {
		session = session.Where("user.id = ? ", obj.ID)
	}

	err := session.Find(&users)
	if err != nil {
		return nil, xerrors.Errorf("relate: %w", err)
	}

	return users, nil
}
