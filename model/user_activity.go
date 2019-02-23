package model

import (
	"github.com/go-xorm/xorm"
	"golang.org/x/xerrors"
)

// UserActivity ...
type UserActivity struct {
	Model        `xorm:"extends" json:",inline"`
	PropertyID   string `xorm:"notnull default('') comment(配置ID) property_id" json:"property_id"`
	ActivityID   string `xorm:"notnull unique(user_activity) default('') comment(活动ID) activity_id" json:"activity_id"`
	UserID       string `xorm:"notnull unique(user_activity) default('') comment(参加活动的用户ID) user_id" json:"user_id"`
	SpreadCode   string `xorm:"notnull unique default('') comment(参加活动的用户推广码) spread_code"  json:"spread_code"`
	Verified     bool   `xorm:"notnull default('')  comment(校验通过) verified" json:"verified"`
	SpreadNumber int64  `xorm:"notnull default(0) comment(推广数) spread_number" json:"spread_number"`
}

// NewUserActivity ...
func NewUserActivity(id string) *UserActivity {
	return &UserActivity{
		Model: Model{
			ID: id,
		},
	}
}

// Get ...
func (obj *UserActivity) Get() (bool, error) {
	return Get(nil, obj)
}

// CodeSpread ...
func (obj *UserActivity) CodeSpread() (*Spread, error) {
	var info struct {
		UserActivity UserActivity `xorm:"extends"`
		Spread       Spread       `xorm:"extends"`
	}
	b, e := DB().Table(obj).Join("left", info.Spread, "user_activity.user_id = spread_code.user_id").
		Where("user_activity.spread_code = ?", obj.SpreadCode).Get(&info)
	if e != nil {
		return nil, e
	}
	if !b {
		e = xerrors.New("property not found")
		return nil, e
	}
	*obj = info.UserActivity

	return &info.Spread, nil
}

// Property ...
func (obj *UserActivity) Property(session *xorm.Session) (*Property, error) {
	var info struct {
		UserActivity UserActivity `xorm:"extends"`
		Property     Property     `xorm:"extends"`
	}
	if session == nil {
		session = DB().NewSession()
	}
	b, e := session.Table(obj).Join("left", info.Property, "user_activity.property_id = property.id").
		Where("user_activity.id = ?", obj.ID).
		Get(&info)
	if e != nil {
		return nil, e
	}
	if !b {
		e = xerrors.New("property not found")
		return nil, e
	}
	*obj = info.UserActivity
	return &info.Property, nil
}
