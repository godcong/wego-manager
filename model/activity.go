package model

import (
	"golang.org/x/xerrors"
)

// Activity ...
type Activity struct {
	Model      `xorm:"extends" json:",inline"`
	UserID     string `xorm:"notnull unique default('') comment(创建活动的用户ID) user_id" json:"user_id"`
	PropertyID string `xorm:"notnull unique default('') comment(配置ID) property_id" json:"property_id"`
	Code       string `xorm:"notnull unique default('') comment(活动码) code" json:"code"`  //活动码
	Verify     bool   `xorm:"notnull default(false) comment(是否校验) verify" json:"verify"` //是否校验
	Mode       string `xorm:"notnull default('') comment(活动模式) mode" json:"mode"`        //活动模式
}

// NewActivity ...
func NewActivity(id string) *Activity {
	return &Activity{
		Model: Model{ID: id},
	}
}

// CodeProperty ...
func (obj *Activity) CodeProperty() (*Property, error) {
	var info struct {
		Activity *Activity `xorm:"extends"`
		Property *Property `xorm:"extends"`
	}
	b, e := DB().Table(obj).Join("left", info.Property, "activity.property_id = property.id").
		Where("activity.code = ?", obj.Code).Get(&info)
	if e != nil {
		return nil, e
	}
	if !b {
		e = xerrors.New("property not found")
		return nil, e
	}
	return info.Property, nil
}
