package model

import "golang.org/x/xerrors"

// UserActivity ...
type UserActivity struct {
	Model        `xorm:"extends" json:",inline"`
	ActivityID   string `xorm:"notnull unique default('') comment(活动ID) activity_id" json:"activity_id"`
	UserID       string `xorm:"notnull unique default('') comment(参加活动的用户ID) user_id" json:"user_id"`
	SpreadCode   string `xorm:"notnull unique default('') comment(参加活动的用户推广码) spread_code"  json:"spread_code"`
	SpreadNumber int64  `xorm:"notnull default('') comment(推广数) spread_code" json:"spread_number"`
}

// CodeSpread ...
func (obj *UserActivity) CodeSpread() (*Spread, error) {
	var info struct {
		UserActivity *UserActivity `xorm:"extends"`
		Spread       *Spread       `xorm:"extends"`
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
	return info.Spread, nil
}
