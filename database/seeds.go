package database

import (
	"github.com/godcong/wego-auth-manager/model"
)

// MenuSeeders ...
func MenuSeeders() {
	var menus []*model.Menu
	menus = append(menus, &model.Menu{
		Name:        "控制台",
		PID:         "",
		Icon:        "",
		Slug:        "system.index",
		URL:         "dashboard/index",
		Description: "后台首页",
	})

}
