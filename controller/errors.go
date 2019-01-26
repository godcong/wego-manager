package controller

import (
	"github.com/gin-gonic/gin"
)

func NewError(ctx *gin.Context, msg string) {
	fail(ctx, msg)
}
