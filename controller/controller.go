package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/go-auth-manager/model"

	"net/http"
)

// ListAccounts godoc
// @Summary List users
// @Description get users
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {array} model.User
// @Failure 400 {object} controller.Ret
// @Router /accounts [get]
func ListAccounts(ctx *gin.Context) {
	users, err := model.Users()
	if err != nil {

		ctx.JSON(http.StatusOK, NewError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, users)
}
