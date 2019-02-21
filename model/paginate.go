package model

import (
	"github.com/go-xorm/xorm"
	"github.com/godcong/wego-auth-manager/config"
	log "github.com/sirupsen/logrus"
	"math"
	"net/url"
	"strconv"
)

// Paginator ...
type Paginator interface {
	Find(*xorm.Session, interface{}) error
	FindWhere(*xorm.Session, interface{}) error
	Count(Modeler) error
	Page() error
	PageWhere(m Modeler) error
}

// Paginate ...
type Paginate struct {
	Current   int
	Total     int64
	TotalPage int
	Limit     int
	Desc      bool
}

// FindWhere ...
func (obj *Paginate) FindWhere(m Modeler, detail interface{}) error {
	session := DB().Where(m).Limit(obj.Limit, obj.Current)
	if obj.Desc {
		session = session.Desc("created_at")
	}
	e := session.Find(detail)
	if e != nil {
		log.Error(e)
		return e
	}
	return nil
}

// Count ...
func (obj *Paginate) Count(m Modeler) (e error) {
	obj.Total, e = Count(nil, m)
	obj.TotalPage = PageNumber(float64(obj.Total), float64(obj.Limit))
	return
}

// Find ...
func (obj *Paginate) Find(detail interface{}) error {
	session := DB().Limit(obj.Limit, obj.Current)
	if obj.Desc {
		session = DB().Desc("created_at")
	}
	e := session.Find(detail)
	if e != nil {
		return e
	}
	return nil
}

// PageNumber ...
func PageNumber(total, limit float64) int {
	if total > 0 && limit != 0 {
		return int(math.Ceil(total / limit))
	}
	return 0
}

// ParsePaginate ...
func ParsePaginate(v url.Values) *Paginate {
	desc, _ := strconv.ParseBool(v.Get("desc"))
	return &Paginate{
		Current: config.MustInt(v.Get("current"), 0),
		Limit:   config.MustInt(v.Get("limit"), 50),
		Desc:    desc,
	}
}
