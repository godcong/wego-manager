package model

import (
	"github.com/godcong/go-auth-manager/config"
	"net/url"
)

// Paginate ...
type Paginate struct {
	Current string
	Total   string
	Limit   string
	Order   string
	Detail  interface{}
}

// ParsePaginate ...
func ParsePaginate(v url.Values) *Paginate {
	return &Paginate{
		Current: config.MustString(v.Get("current"), "0"),
		Limit:   config.MustString(v.Get("limit"), "50"),
		Order:   config.MustString(v.Get("order"), "desc"),
		Detail:  nil,
	}
}
