package controller

import (
	"github.com/go-xorm/xorm"
	"github.com/godcong/wego-auth-manager/config"
	"github.com/godcong/wego-auth-manager/model"
	log "github.com/sirupsen/logrus"
	"math"
	"net/url"
	"strconv"
)

// Paginator ...
type Paginator interface {
	Find(*xorm.Session, interface{}) error
	FindWhere(*xorm.Session, interface{}) error
	Count(model.Modeler) error
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
func (obj *Paginate) FindWhere(m model.Modeler, detail interface{}) error {
	session := model.DB().Where(m).Limit(obj.Limit, obj.Current)
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
func (obj *Paginate) Count(m model.Modeler) (e error) {
	obj.Total, e = model.Count(nil, m)
	obj.TotalPage = PageNumber(float64(obj.Total), float64(obj.Limit))
	return
}

// Find ...
func (obj *Paginate) Find(detail interface{}) error {
	session := model.DB().Limit(obj.Limit, obj.Current)
	if obj.Desc {
		session = model.DB().Desc("created_at")
	}
	e := session.Find(detail)
	if e != nil {
		log.Error(e)
		return e
	}
	return nil
}

// PaginateUserActivity ...
type PaginateUserActivity struct {
	*Paginate `json:",inline"`
	Detail    []*model.UserActivity
}

// Page ...
func (obj *PaginateUserActivity) Page() error {
	e := obj.Find(&obj.Detail)
	if e != nil {
		log.Error(e)
		return e
	}
	return obj.Count(&model.UserActivity{})
}

// PageWhere ...
func (obj *PaginateUserActivity) PageWhere(m model.Modeler) error {
	e := obj.FindWhere(m, &obj.Detail)
	if e != nil {
		log.Error(e)
		return e
	}
	return obj.Count(m)
}

// PageNumber ...
func PageNumber(total, limit float64) int {
	if total > 0 && limit != 0 {
		return int(math.Ceil(total / limit))
	}
	return 0
}

// PageUserActivity ...
func PageUserActivity(v url.Values) *PaginateUserActivity {
	return &PaginateUserActivity{
		Paginate: parsePaginate(v),
		Detail:   nil,
	}
}

func parsePaginate(v url.Values) *Paginate {
	desc, _ := strconv.ParseBool(v.Get("desc"))
	return &Paginate{
		Current: config.MustInt(v.Get("current"), 0),
		Limit:   config.MustInt(v.Get("limit"), 50),
		Desc:    desc,
	}
}
