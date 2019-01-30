package model

import (
	"golang.org/x/exp/xerrors"
)

// Role ...
type Role struct {
	Model       `xorm:"extends"`
	Name        string `xorm:"name"`
	Slug        string `xorm:"slug"`
	Description string `xorm:"description"`
	Level       int    `xorm:"level"`
}

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
