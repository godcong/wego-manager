package database

import (
	"github.com/godcong/wego-auth-manager/model"
	log "github.com/sirupsen/logrus"
)

var migrateionTable = []model.Modeler{
	&model.User{},
	&model.Permission{},
	&model.Role{},
	&model.PermissionRole{},
	&model.PermissionUser{},
	&model.RoleUser{},
	&model.Property{},
	&model.Notify{},
}

// Migrate ...
func Migrate() error {
	var err error
	for _, m := range migrateionTable {
		log.Info("syncing")
		err = model.DB().Sync2(m)
		if err != nil {
			return err
		}
	}
	return nil
}
