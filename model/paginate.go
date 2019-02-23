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
	Find(interface{}) error
	FindWhere(*xorm.Session, interface{}) error
	Count(*xorm.Session, Modeler) error
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

// Start ...
func (obj *Paginate) Start() int {
	if obj.Current == 0 || obj.Limit == 0 {
		return obj.Current
	}
	return obj.Current * obj.Limit
}

// FindWhere ...
func (obj *Paginate) FindWhere(session *xorm.Session, detail interface{}) error {
	session = MustSession(session).Limit(obj.Limit, obj.Start())
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

// FindWhere2 ...
func (obj *Paginate) FindWhere2(m Modeler, detail interface{}) error {
	session := DB().Table(m).Limit(obj.Limit, obj.Start())
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
func (obj *Paginate) Count(session *xorm.Session, m Modeler) (e error) {
	obj.Total, e = Count(session, m)
	obj.TotalPage = PageNumber(float64(obj.Total), float64(obj.Limit))
	return
}

// Find ...
func (obj *Paginate) Find(detail interface{}) error {
	session := DB().Limit(obj.Limit, obj.Start())
	if obj.Desc {
		session = session.Desc("created_at")
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
