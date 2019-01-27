package model

// Role ...
type Role struct {
	Model       `xorm:"extends"`
	Name        string `xorm:"name"`
	Slug        string `xorm:"slug"`
	Description string `xorm:"description"`
	Level       int    `xorm:"level"`
}

// Count ...
func (Role) Count() (int64, error) {
	return Count(nil, m)
}
