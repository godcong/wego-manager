package controller

import (
	"github.com/gin-gonic/gin"
)

// NewError ...
func NewError(ctx *gin.Context, msg string) {
	fail(ctx, msg)
}
