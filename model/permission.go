package model

// Permission ...
type Permission struct {
	Model           `xorm:"extends"`
	Name            string `xorm:"name"`
	Slug            string `xorm:"slug"`
	Description     string `xorm:"description"`
	PermissionModel string `xorm:"permission_model"`
}

// Count ...
func (p *Permission) Count() (int64, error) {
	return Count(nil, p)
}
