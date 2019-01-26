package model

import (
	"github.com/godcong/go-auth-manager/config"
	"github.com/google/uuid"
	"github.com/xormplus/xorm"
)

// DB ...
type DataBase struct {
	config *config.Configure
	eng    *xorm.Engine
}

var globalDB *DataBase

// DB ...
func DB() *DataBase {
	if globalDB == nil {
		InitDB(config.Config())
	}
	return globalDB
}

// Model ...
type Model struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt int64     `json:"created_at" xorm:"created"`
	UpdatedAt int64     `json:"deleted_at" xorm:"updated"`
	DeletedAt *int64    `json:"deleted_at" xorm:"deleted"`
	Version   int       `json:"version" xorm:"version"`
}

// InitDB ...
func InitDB(cfg *config.Configure) {
	engine, err := xorm.NewEngine(cfg.Database.Type, cfg.Database.Source())
	if err != nil {
		panic(err)
	}
	globalDB = &DataBase{
		config: cfg,
		eng:    engine,
	}
}

// FindByID ...
func FindByID(id string, model interface{}) error {

}
