package model

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
