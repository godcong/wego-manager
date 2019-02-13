package model

import (
	"golang.org/x/xerrors"
)

// Role ...
type Role struct {
	Model       `xorm:"extends" json:",inline"`
	Enable      bool   `xorm:"enable" json:"enable"`
	Name        string `xorm:"name" json:"name"`
	Slug        string `xorm:"slug" json:"slug"`
	Description string `xorm:"description" json:"description"`
	Level       int    `xorm:"level" json:"level"`
}

// RoleSlugAdmin ...
const RoleSlugAdmin = "admin"

// RoleSlugUser ...
const RoleSlugUser = "user"

// NewRole ...
func NewRole(id string) *Role {
	return &Role{
		Model: Model{
			ID: id,
		},
	}
}

// Get ...
func (obj *Role) Get() (bool, error) {
	return Get(nil, obj)
}

// Update ...
func (obj *Role) Update(cols ...string) (int64, error) {
	return Update(nil, obj.ID, obj)
}

// Roles ...
func (obj *Role) Roles() ([]*Role, error) {
	var roles []*Role
	err := DB().Table(obj).Find(&roles)
	if err != nil {
		return nil, xerrors.Errorf("find: %w", err)
	}
	return roles, nil
}

// Permissions ...
func (obj *Role) Permissions() ([]*Permission, error) {
	var permissions []*Permission
	session := DB().Table(&Permission{}).Select("permission.*").
		Join("left", &PermissionRole{}, "permission_role.role_id = role.id")

	if obj.ID != "" {
		session = session.Where("role.id = ? ", obj.ID)
	}

	err := session.Find(&permissions)
	if err != nil {
		return nil, xerrors.Errorf("relate: %w", err)
	}

	return permissions, nil
}

// Users ...
func (obj *Role) Users() ([]*User, error) {
	var users []*User
	session := DB().Table(&User{}).Select("user.*").
		Join("left", &RoleUser{}, "role_user.role_id = role.id")

	if obj.ID != "" {
		session = session.Where("role.id = ? ", obj.ID)
	}

	err := session.Find(&users)
	if err != nil {
		return nil, xerrors.Errorf("relate: %w", err)
	}

	return users, nil
}
