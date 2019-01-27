package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/go-auth-manager/model"
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
		Success(ctx, users)
	}
}

// UserAdd godoc
// @Summary Add user
// @Description Add user
// @Tags user2
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param account body User true "user update info"
// @success 200 {array} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /user [post]
func UserAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// UserUpdate godoc
// @Summary Update user
// @Description Update user
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @Param account body User true "user update info"
// @success 200 {array} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /user/{id} [post]
func UserUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// UserShow godoc
// @Summary Show user
// @Description Show user
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @success 200 {array} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /user/{id} [get]
func UserShow(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// UserDelete godoc
// @Summary Delete user
// @Description Delete user
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @success 200 {array} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /user/{id} [delete]
func UserDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// UserPermissionList ...
func UserPermissionList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
