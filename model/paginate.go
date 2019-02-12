package model

import (
	"github.com/godcong/wego-auth-manager/config"
	"net/url"
)

// Paginator ...
type Paginator interface {
	Paginate(values url.Values) (*Paginate, error)
}

// Paginate ...
type Paginate struct {
	Current   int
	Total     int
	TotalPage int
	Limit     int
	Order     string
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
	}
}
