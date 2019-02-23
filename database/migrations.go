package database

import (
	"github.com/godcong/wego-auth-manager/model"
	"github.com/godcong/wego-auth-manager/util"
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
	&model.Spread{},
	&model.Activity{},
	&model.UserActivity{},
	&model.Menu{},
	&model.Notify{},
	&model.Turnover{},
	&model.UserInformation{},
	&model.WechatUser{},
}

// Migrate ...
func Migrate() error {
	var err error
	for _, m := range migrateionTable {
		log.WithField("name", util.StructureName(m)).Info("syncing")
		err = model.DB().Sync2(m)
		if err != nil {
			return err
		}
	}
	return nil
}
