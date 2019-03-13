package model

import "github.com/go-xorm/xorm"

// Spread ...
type Spread struct {
	Model         `xorm:"extends"`
	UserID        string `xorm:"notnull default('') user_id" json:"user_id"`
	Code          string `xorm:"notnull default('') code" json:"code"`
	ActivityID    string `xorm:"notnull default('') activity_id" json:"activity_id"`
	ParentUserID1 string `xorm:"notnull default('') parent_user_id_1" json:"parent_user_id_1 comment(上1级)"`
	ParentUserID2 string `xorm:"notnull default('') parent_user_id_2" json:"parent_user_id_2 comment(上2级)"`
	ParentUserID3 string `xorm:"notnull default('') parent_user_id_3" json:"parent_user_id_3 comment(上3级)"`
	ParentUserID4 string `xorm:"notnull default('') parent_user_id_4" json:"parent_user_id_4 comment(上4级)"`
	ParentUserID5 string `xorm:"notnull default('') parent_user_id_5" json:"parent_user_id_5 comment(上5级)"`
	ParentUserID6 string `xorm:"notnull default('') parent_user_id_6" json:"parent_user_id_6 comment(上6级)"`
	ParentUserID7 string `xorm:"notnull default('') parent_user_id_7" json:"parent_user_id_7 comment(上7级)"`
	ParentUserID8 string `xorm:"notnull default('') parent_user_id_8" json:"parent_user_id_8 comment(上8级)"`
	ParentUserID9 string `xorm:"notnull default('') parent_user_id_9" json:"parent_user_id_9 comment(上9级)"`
}

// NewSpread ...
func NewSpread(id string) *Spread {
	return &Spread{
		Model: Model{
			ID: id,
		},
	}
}

// Get ...
func (obj *Spread) Get() (bool, error) {
	return Get(nil, obj)
}

// SpreadActivity ...
type SpreadActivity struct {
	Spread   Spread   `xorm:"extends"`
	Activity Activity `xorm:"extends"`
}

// SpreadActivity ...
func (obj *Spread) SpreadActivity(session *xorm.Session) ([]*SpreadActivity, error) {
	var info []*SpreadActivity
	e := MustSession(session).Table(obj).
		Join("left", &Activity{}, "activity.id = spread.activity_id").
		Find(&info)
	if e != nil {
		return nil, e
	}
	return info, e
}
