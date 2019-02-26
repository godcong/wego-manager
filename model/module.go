package model

// Module ...
type Module struct {
	Model `xorm:"extends" json:",inline"`
	Name  string `xorm:"name" json:"name"`
	URL   string `xorm:"url" json:"url"`
}
