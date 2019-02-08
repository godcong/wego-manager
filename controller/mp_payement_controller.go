package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
)

// MpPaymentBill godoc
// @Summary List permission
// @Description List permission
// @Tags permission
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Permission ID"
// @success 200 {array} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /mp/payment/bill [get]
func MpPaymentBill(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.MustUser(ctx.Get("user"))

		_, err := user.Property()
		if err != nil {
			Error(ctx, err)
			return
		}

	}
}
