package model

// Module ...
type Module struct {
	Model  `xorm:"extends" json:",inline"`
	Name   string `xorm:"name" json:"name"`
	Alias  string `xorm:"alias" json:"alias"`
	Site   string `xorm:"site" json:"site"`
	Deploy string `xorm:"deploy" json:"deploy"`
	Log    string `xorm:"log" json:"log"`
}

// NewModule ...
func NewModule(id string) *Module {
	return &Module{
		Model: Model{
			ID: id,
		},
	}
}

// Get ...
func (obj *Module) Get() (bool, error) {
	return Get(nil, obj)
}
