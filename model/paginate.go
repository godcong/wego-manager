package model

import (
	"github.com/godcong/go-auth-manager/config"
	"net/url"
)

// Paginator ...
type Paginator interface {
	Paginate(values url.Values) (*Paginate, error)
	Object() interface{}
}

// Paginate ...
type Paginate struct {
	Current   int
	Total     int
	TotalPage int
	Limit     int
	Order     string
	Detail    interface{}
}

func parsePaginate(v url.Values) *Paginate {
	order := v.Get("order")
	if order != "asc" {
		order = "desc"
	}

	return &Paginate{
		Current: config.MustInt(v.Get("current"), 0),
		Limit:   config.MustInt(v.Get("limit"), 50),
		Order:   order,
		Detail:  nil,
	}
}
