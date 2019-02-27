package model

// Module ...
type Module struct {
	Model   `xorm:"extends" json:",inline"`
	Name    string   `xorm:"notnull default('') name" json:"name"`   //模块名称
	Alias   string   `xorm:"notnull default('') alias" json:"alias"` //模块别名url+别名+process
	Site    string   `xorm:"notnull default('') site" json:"site"`
	Deploy  string   `xorm:"notnull default('') deploy" json:"deploy"`
	Log     string   `xorm:"notnull default('') log" json:"log"`
	Depends []string `xorm:"depends" json:"depends"`
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
