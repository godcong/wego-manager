package model

import "github.com/go-xorm/xorm"

// UserActivityPaginate ...
type UserActivityPaginate struct {
	*Paginate `json:",inline"`
	Detail    []*UserActivity
}

// Page ...
func (obj *UserActivityPaginate) Page() error {
	e := obj.Find(&obj.Detail)
	if e != nil {
		return e
	}
	return obj.Count(nil, &UserActivity{})
}

// PageWhere ...
func (obj *UserActivityPaginate) PageWhere(session *xorm.Session) error {
	e := obj.FindWhere(session.Clone(), &obj.Detail)
	if e != nil {
		return e
	}
	return obj.Count(session, &UserActivity{})
}

// PageUserActivity ...
func PageUserActivity(paginate *Paginate) *UserActivityPaginate {
	return &UserActivityPaginate{
		Paginate: paginate,
		Detail:   nil,
	}
}
