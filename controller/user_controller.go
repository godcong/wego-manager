package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/go-auth-manager/model"
	"golang.org/x/exp/xerrors"
	"log"
)

// UserList godoc
// @Summary List users
// @Description List users
// @Tags user
// @Accept  json
// @Produce  json
// @Param current query string false "paginate:current"
// @Param limit query string false "paginate:limit"
// @Param order query string false "paginate:order"
// @success 200 {array} model.Paginate
// @Failure 400 {object} controller.CodeMessage
// @Router /user [get]
func UserList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		users, err := user.Users()
		if err != nil {
			Error(ctx, err)
			return
		}
		log.Println(users)
		Success(ctx, users)
	}
}

// UserAdd godoc
// @Summary Add user
// @Description Add user
// @Tags user
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param account body User true "user update info"
// @success 200 {object} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /user [post]
func UserAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		err := ctx.BindJSON(&user)
		if err != nil {
			Error(ctx, err)
			return
		}
		_, err = model.Insert(nil, &user)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, user)
	}
}

// UserUpdate godoc
// @Summary Update user
// @Description Update user
// @Tags user
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @Param account body User true "user update info"
// @success 200 {object} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /user/{id} [post]
func UserUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.NewUser(id)
		b, err := user.Get()
		if err != nil || !b {
			Error(ctx, xerrors.Errorf("no users:%w", err))
			return
		}
		err = ctx.BindJSON(user)
		if err != nil {
			Error(ctx, err)
			return
		}

		_, err = model.Update(nil, id, user)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, user)
	}
}

// UserShow godoc
// @Summary Show user
// @Description Show user
// @Tags user
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @success 200 {object} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /user/{id} [get]
func UserShow(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.NewUser(id)
		_, err := model.Get(nil, user)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, user)
	}
}

// UserDelete godoc
// @Summary Delete user
// @Description Delete user
// @Tags user
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @success 200 {object} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /user/{id} [delete]
func UserDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.NewUser(id)
		_, err := model.Delete(nil, user)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, user)
	}
}

// UserPermissionList godoc
// @Summary List permission
// @Description List permission
// @Tags user
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @success 200 {array} model.Permission
// @Failure 400 {object} controller.CodeMessage
// @Router /user/{id}/permission [get]
func UserPermissionList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.NewUser(id)
		permissions, err := user.Permissions()
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, permissions)
	}
}
