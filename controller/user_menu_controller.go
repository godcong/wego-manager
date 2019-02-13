package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

// UserMenuList godoc
// @Summary List user properties
// @Description List user properties
// @Tags menu
// @Accept  json
// @Produce  json
// @Param current query string false "paginate:current"
// @Param limit query string false "paginate:limit"
// @Param order query string false "paginate:order"
// @success 200 {array} model.Menu
// @Failure 400 {object} controller.CodeMessage
// @Router /user/menu [get]
func UserMenuList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var menu model.Menu
		menus, err := menu.Menus()
		if err != nil {
			Error(ctx, err)
			return
		}
		log.Info(menus)
		Success(ctx, menus)
	}
}

// UserMenuAdd godoc
// @Summary Add user menu
// @Description Add user menu
// @Tags menu
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param account body Property true "menu update info"
// @success 200 {object} model.Menu
// @Failure 400 {object} controller.CodeMessage
// @Router /user/menu [post]
func UserMenuAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var menu model.Menu
		err := ctx.BindJSON(&menu)
		if err != nil {
			Error(ctx, err)
			return
		}
		_, err = model.Insert(nil, &menu)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, menu)
	}
}

// UserMenuUpdate godoc
// @Summary Update user menu
// @Description Update user menu
// @Tags menu
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Property ID"
// @Param account body Property true "menu update info"
// @success 200 {object} model.Menu
// @Failure 400 {object} controller.CodeMessage
// @Router /user/menu/{id} [post]
func UserMenuUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		menu := model.NewMenu(id)
		b, err := menu.Get()
		if err != nil || !b {
			Error(ctx, xerrors.Errorf("no menu:%w", err))
			return
		}
		err = ctx.BindJSON(menu)
		if err != nil {
			Error(ctx, err)
			return
		}

		_, err = model.Update(nil, id, menu)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, menu)
	}
}

// UserMenuShow godoc
// @Summary Show user menu
// @Description Show user menu
// @Tags menu
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Property ID"
// @success 200 {object} model.Menu
// @Failure 400 {object} controller.CodeMessage
// @Router /user/menu/{id} [get]
func UserMenuShow(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		menu := model.NewMenu(id)
		_, err := model.Get(nil, menu)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, menu)
	}
}

// UserMenuDelete godoc
// @Summary Delete user menu
// @Description Delete user menu
// @Tags menu
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Property ID"
// @success 200 {object} model.Menu
// @Failure 400 {object} controller.CodeMessage
// @Router /user/menu/{id} [delete]
func UserMenuDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		menu := model.NewMenu(id)
		_, err := model.Delete(nil, menu)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, menu)
	}
}
