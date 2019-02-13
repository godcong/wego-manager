package model

import "golang.org/x/xerrors"

// Menu ...
type Menu struct {
	Model       `xorm:"extends" json:",inline"`
	PID         string `xorm:"pid"` //parent id
	Name        string `xorm:"name"`
	Icon        string `xorm:"icon"`
	Slug        string `xorm:"slug"`
	URL         string `xorm:"url"`
	Active      string `xorm:"active"`
	Description string `xorm:"description"`
	Sort        int    `xorm:"sort"`
}

// NewMenu ...
func NewMenu(id string) *Menu {
	return &Menu{
		Model: Model{
			ID: id,
		},
	}
}

// Menus ...
func (obj *Menu) Menus() ([]*Menu, error) {
	var menus []*Menu
	err := DB().Table(obj).Find(&menus)
	if err != nil {
		return nil, xerrors.Errorf("find: %w", err)
	}
	return menus, nil
}
