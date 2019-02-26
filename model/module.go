package model

// Module ...
type Module struct {
	Model `xorm:"extends" json:",inline"`
	Name  string `xorm:"name" json:"name"`
	URL   string `xorm:"url" json:"url"`
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
