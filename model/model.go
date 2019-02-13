package model

import (
	"github.com/go-xorm/xorm"
	"github.com/godcong/wego-auth-manager/config"
	"github.com/godcong/wego-auth-manager/util"
	"github.com/google/uuid"
	"github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

// DataBase ...
type DataBase struct {
	config *config.Configure
	//modelers []Modeler
	*xorm.Engine
}

var globalDB *DataBase

// DB ...
func DB() *DataBase {
	if globalDB == nil {
		InitDB(config.Config())
	}
	return globalDB
}

// Modeler ...
type Modeler interface {
	BeforeInsert()
	GetID() string
	Get() (bool, error)
	Update(cols ...string) (int64, error)
	Count() (int64, error)
}

// Count ...
func Count(session *xorm.Session, obj Modeler) (int64, error) {
	return MustSession(session).Count(obj)
}

// FindByID ...
func FindByID(id string, obj Modeler) error {
	return DB().Where("id = ?", id).Find(obj)
}

// Find ...
func Find(session *xorm.Session, obj Modeler) error {
	return MustSession(session).Find(obj)
}

// Get ...
func Get(session *xorm.Session, obj Modeler) (bool, error) {
	return MustSession(session).Get(obj)
}

// Update ...
func Update(session *xorm.Session, id string, obj Modeler) (int64, error) {
	return MustSession(session).ID(id).Update(obj)
}

// UpdateWithColumn ...
func UpdateWithColumn(session *xorm.Session, id string, obj Modeler, cols ...string) (int64, error) {
	return MustSession(session).ID(id).Cols(cols...).Update(obj)
}

// Insert ...
func Insert(session *xorm.Session, obj Modeler) (int64, error) {
	return MustSession(session).InsertOne(obj)
}

// Delete ...
func Delete(session *xorm.Session, obj Modeler) (int64, error) {
	return MustSession(session).Delete(obj)
}

// MustSession ...
func MustSession(session *xorm.Session) *xorm.Session {
	if session == nil {
		session = DB().NewSession()
	}
	return session
}

// CSRF ...
type CSRF struct {
	CSRFToken string `json:"csrf_token"`
}

// Model ...
type Model struct {
	ID        string `json:"id" xorm:"id uuid pk comment(默认主键)"`
	CreatedAt int64  `json:"-" xorm:"created comment(创建时间)"`
	UpdatedAt int64  `json:"-" xorm:"updated comment(更新时间)"`
	DeletedAt *int64 `json:"-" xorm:"deleted comment(删除时间)"`
	Version   int    `json:"-" xorm:"version comment(版本)"`
}

// Get ...
func (m *Model) Get() (bool, error) {
	return Get(nil, m)
}

// Update ...
func (m *Model) Update(cols ...string) (int64, error) {
	if cols == nil {
		return Update(nil, m.ID, m)
	}
	return UpdateWithColumn(nil, m.ID, m, cols...)
}

// BeforeInsert ...
func (m *Model) BeforeInsert() {
	m.ID = uuid.Must(uuid.NewUUID()).String()
}

// Count ...
func (m *Model) Count() (int64, error) {
	return Count(nil, m)
}

// GetID ...
func (m *Model) GetID() string {
	return m.ID
}

// InitDB ...
func InitDB(cfg *config.Configure) *DataBase {
	engine, err := xorm.NewEngine(cfg.Database.Type, cfg.Database.Source())
	if err != nil {
		panic(err)
	}
	if cfg.Database.ShowSQL {
		engine.ShowSQL(true)
	}
	globalDB = &DataBase{
		config: cfg,
		//modelers: modelerTables(),
		Engine: engine,
	}
	return globalDB
}

// TokenSub ...
type TokenSub struct {
	ID string
}

// DecodeUser ...
func DecodeUser(token string) (*User, error) {
	t := TokenSub{}
	sub, err := util.DecryptJWT([]byte(globalDB.config.General.TokenKey), token)
	log.Info("sub", sub)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal([]byte(sub), &t)
	if err != nil {
		if err != nil {
			return nil, err
		}
	}
	return &User{}, nil
}
