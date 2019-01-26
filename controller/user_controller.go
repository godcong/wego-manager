package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/go-auth-manager/model"
)

// UserList godoc
// @Summary List users
// @Description get users
// @Accept  json
// @Produce  json
// @Param current query string false "paginate:current"
// @Param limit query string false "paginate:limit"
// @Param order query string false "paginate:order"
// @success 200 {array} model.User
// @Failure 200 {object} controller.RetDetail
// @Router /user [get]
func UserList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		val := ctx.Request.URL.Query()
		model.ParsePaginate(val)
		users, err := model.Users()
		if err != nil {
			NewError(ctx, err.Error())
			return
		}
		detail(ctx, users)
	}

}

// UserUpdate godoc
// @Summary Update user
// @Description Update user
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @Param account body model.User true "user update info"
// @success 200 {array} model.User
// @Failure 200 {object} controller.RetDetail
// @Router /user/{id} [post]
func UserUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var user model.User
		var err error
		err = model.FindByID(id, &user)

		users, err := model.Users()
		if err != nil {
			NewError(ctx, err.Error())
			return
		}
		detail(ctx, users)
	}
}

// UserShow godoc
// @Summary Show user
// @Description Show user
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @Param account body model.User true "user update info"
// @success 200 {array} model.User
// @Failure 200 {object} controller.RetDetail
// @Router /user/{id} [post]
func UserShow(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var user model.User
		var err error
		err = model.FindByID(id, &user)

		users, err := model.Users()
		if err != nil {
			NewError(ctx, err.Error())
			return
		}
		detail(ctx, users)
	}
}

// UserDelete godoc
// @Summary Delete user
// @Description Delete user
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @Param account body model.User true "user update info"
// @success 200 {array} model.User
// @Failure 200 {object} controller.RetDetail
// @Router /user/{id} [post]
func UserDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var user model.User
		var err error
		err = model.FindByID(id, &user)

		users, err := model.Users()
		if err != nil {
			NewError(ctx, err.Error())
			return
		}
		detail(ctx, users)
	}
}
