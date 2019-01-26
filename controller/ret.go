package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Ret struct {
	Code    int    `json:"code" example:"-1"`
	Message string `json:"message" example:"status bad request"`
}

type RetDetail struct {
	Ret    `json:",inline"`
	Detail interface{}
}

func result(ctx *gin.Context, code int, msg string, detail interface{}) {
	if detail != nil {
		ctx.JSON(http.StatusOK, &RetDetail{
			Ret: Ret{
				Code:    code,
				Message: msg,
			},
			Detail: detail,
		})
		return
	}
	ctx.JSON(http.StatusOK, &Ret{
		Code:    code,
		Message: msg,
	})
}

func fail(ctx *gin.Context, msg string) {
	result(ctx, -1, msg, nil)
}

func success(ctx *gin.Context, msg string) {
	result(ctx, 0, msg, nil)
}

func detail(ctx *gin.Context, d interface{}) {
	result(ctx, 0, "success", d)
}
