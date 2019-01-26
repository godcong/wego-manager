package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/godcong/go-auth-manager/config"
	"github.com/google/uuid"
	"github.com/xormplus/xorm"
	"log"
)

// DataBase ...
type DataBase struct {
	config   *config.Configure
	modelers []Modeler
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

// Sync ...
func Sync() error {
	var err error
	for _, m := range DB().modelers {
		log.Println("syncing")
		err = DB().Sync2(m)
		if err != nil {
			return err
		}
	}
	return nil
}

// Modeler ...
type Modeler interface {
	GetID() string
}

// Model ...
type Model struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt int64     `json:"created_at" xorm:"created"`
	UpdatedAt int64     `json:"deleted_at" xorm:"updated"`
	DeletedAt *int64    `json:"deleted_at" xorm:"deleted"`
	Version   int       `json:"version" xorm:"version"`
}

// ID ...
func (m *Model) GetID() string {
	return m.ID.String()
}

// InitDB ...
func InitDB(cfg *config.Configure) *DataBase {
	engine, err := xorm.NewEngine(cfg.Database.Type, cfg.Database.Source())
	if err != nil {
		panic(err)
	}
	globalDB = &DataBase{
		config:   cfg,
		modelers: modelerTables(),
		Engine:   engine,
	}
	return globalDB
}

func modelerTables() []Modeler {
	return []Modeler{
		&User{},
		&Permission{},
		&Role{},
	}
}

// FindByID ...
func FindByID(id string, model interface{}) error {
	return DB().Where("id = ?", id).Find(model)
}
