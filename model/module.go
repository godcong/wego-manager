package model

// Module ...
type Module struct {
	Model   `xorm:"extends" json:",inline"`
	Name    string   `xorm:"notnull default('') name" json:"name"`     //模块名称
	Alias   string   `xorm:"notnull default('') alias" json:"alias"`   //模块别名url+别名+process
	Site    string   `xorm:"notnull default('') site" json:"site"`     //模块地址
	OS      string   `xorm:"os" json:"os"`                             //运行中的系统类别
	Deploy  string   `xorm:"notnull default('') deploy" json:"deploy"` //部署方式
	Log     string   `xorm:"notnull default('') log" json:"log"`       //日志目录(类别)
	Depends []string `xorm:"depends" json:"depends"`                   //关联模块列表
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
