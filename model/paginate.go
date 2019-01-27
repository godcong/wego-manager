package model

import (
	"github.com/godcong/go-auth-manager/config"
	"github.com/xormplus/xorm"
	"net/url"
)

// Paginate ...
type Paginate struct {
	Current   int
	Total     int
	TotalPage int
	Limit     int
	Order     string
	Detail    interface{}
}

// ParsePaginate ...
func ParsePaginate(v url.Values) *Paginate {

	return &Paginate{
		Current: config.MustInt(v.Get("current"), 0),
		Limit:   config.MustInt(v.Get("limit"), 50),
		Order:   config.MustString(v.Get("order"), "desc"),
		Detail:  nil,
	}
}

// Engine ...
func (p *Paginate) Engine() *xorm.Session {
	return DB().Limit(p.Limit, p.Current*p.Limit).OrderBy(p.Order)
}
