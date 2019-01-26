package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/go-auth-manager/model"
	"log"
)

// VisitLog ...
func VisitLog(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println(ctx.Request.Header)
		log.Println(ctx.Accepted)
		l.Permission = handleFuncName(ctx)
		l.Method = ctx.Request.Method
		l.URL = ctx.Request.URL.String()
		token := ctx.GetHeader("token")
		log.Println(token)

		user, err := decodeUser(ctx)
		if err != nil {
			l.Err = err.Error()
		}
		l.UserID = user.ID

		RemoteIP(ctx)
	}
}

// RemoteIP ...
func RemoteIP(ctx *gin.Context) {
	host := ctx.GetHeader("REMOTE-HOST")
	if host == "" {
		host = ctx.Request.RemoteAddr
	}
	log.Println("host:", host)
}
