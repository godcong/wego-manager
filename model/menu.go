package model

// Menu ...
type Menu struct {
	Model       `xorm:"extends"`
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
