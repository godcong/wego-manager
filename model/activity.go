package model

import (
	"golang.org/x/xerrors"
)

// Activity ...
type Activity struct {
	Model      `xorm:"extends" json:",inline"`
	UserID     string `xorm:"notnull default('') comment(创建活动的用户ID) user_id" json:"user_id"`
	PropertyID string `xorm:"notnull default('') comment(配置ID) property_id" json:"property_id"`
	Name       string `xorm:"notnull default('') comment(名称) name" json:"name"`                        //活动名称
	Code       string `xorm:"notnull unique default('') comment(活动码) code" json:"code"`                //活动码
	IsPublic   bool   `xorm:"notnull default(false) comment(公开) is_public" json:"is_public"`           //是否公开
	NeedVerify bool   `xorm:"notnull default(false) comment(是否需要校验) need_verify" json:"need_verify"`   //是否需要校验
	Comment    string `xorm:"notnull default('')  comment(活动介绍) varchar(2048) comment" json:"comment"` //活动介绍
	Mode       string `xorm:"notnull default('') comment(活动模式) mode" json:"mode"`                      //活动模式
}

// NewActivity ...
func NewActivity(id string) *Activity {
	return &Activity{
		Model: Model{ID: id},
	}
}

// Get ...
func (obj *Activity) Get() (bool, error) {
	return Get(nil, obj)
}

// Update ...
func (obj *Activity) Update(cols ...string) (int64, error) {
	return Update(nil, obj.ID, obj)
}

// CodeProperty ...
func (obj *Activity) CodeProperty() (*Property, error) {
	var info struct {
		Activity Activity `xorm:"extends"`
		Property Property `xorm:"extends"`
	}
	b, e := Table(obj).Join("left", info.Property, "activity.property_id = property.id").
		Where("activity.code = ?", obj.Code).Get(&info)
	if e != nil {
		return nil, e
	}
	if !b {
		e = xerrors.New("property not found")
		return nil, e
	}
	*obj = info.Activity
	return &info.Property, nil
}

// Property ...
func (obj *Activity) Property() (*Property, error) {
	var info struct {
		Activity Activity `xorm:"extends"`
		Property Property `xorm:"extends"`
	}
	b, e := Table(obj).Join("left", info.Property, "activity.property_id = property.id").
		Where("activity.id = ?", obj.ID).Get(&info)
	if e != nil {
		return nil, e
	}
	if !b {
		e = xerrors.New("property not found")
		return nil, e
	}
	*obj = info.Activity
	return &info.Property, nil
}

// Activities ...
func (obj *Activity) Activities() ([]*Activity, error) {
	var activities []*Activity
	e := DB().Find(&activities)
	if e != nil {
		return nil, e
	}
	return activities, e
}

// User ...
func (obj *Activity) User() (*User, error) {
	return nil, nil
}
