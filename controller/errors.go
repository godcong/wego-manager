package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CodeMessage ...
type CodeMessage struct {
	Code    int    `json:"code" example:"-1"`
	Message string `json:"message" example:"status bad request"`
}

// Error ...
func Error(ctx *gin.Context, err error) {
	log.Println(err)
	ctx.JSON(http.StatusBadRequest, CodeMessage{
		Code:    -1,
		Message: err.Error(),
	})
}
