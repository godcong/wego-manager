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
// @Failure 200 {object} controller.Ret
// @Router /user [get]
func UserList(ctx *gin.Context) {
	users, err := model.Users()
	if err != nil {
		NewError(ctx, err.Error())
		return
	}
	detail(ctx, users)
}
