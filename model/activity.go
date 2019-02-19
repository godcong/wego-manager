package model

// Activity ...
type Activity struct {
	Model      `xorm:"extends" json:",inline"`
	SpreadMark string `xorm:"notnull unique default('') spread_mark" json:"spread_mark"`
}
