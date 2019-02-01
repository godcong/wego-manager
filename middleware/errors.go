package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CodeMessage ...
type CodeMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error ...
func Error(ctx *gin.Context, err error) {
	log.Println(err)
	ctx.JSON(http.StatusForbidden, CodeMessage{
		Code:    -1,
		Message: err.Error(),
	})
}
