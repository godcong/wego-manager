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

func ModuleSeeders() {
	var modules []*model.Module
	modules = append(modules, &model.Module{
		Name:    "spread",
		Alias:   "spread",
		Site:    "https://github.com/godcong/wego-spread-service",
		OS:      "",
		Deploy:  "",
		Log:     "http://localhost:9200",
		Depends: []string{"manager"},
	})
}
