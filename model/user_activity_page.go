package model

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
	return obj.Count(&UserActivity{})
}

// PageWhere ...
func (obj *UserActivityPaginate) PageWhere(m Modeler) error {
	e := obj.FindWhere(m, &obj.Detail)
	if e != nil {
		return e
	}
	return obj.Count(m)
}

// PageUserActivity ...
func PageUserActivity(paginate *Paginate) *UserActivityPaginate {
	return &UserActivityPaginate{
		Paginate: paginate,
		Detail:   nil,
	}
}
