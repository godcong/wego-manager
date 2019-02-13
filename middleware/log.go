package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// VisitLog ...
func VisitLog(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//log.Info(ctx.Request.Header)
		//log.Info(ctx.Accepted)
		//l.Permission = handleFuncName(ctx)
		//l.Method = ctx.Request.Method
		//l.URL = ctx.Request.URL.String()
		//token := ctx.GetHeader("token")
		//log.Info(token)
		//
		//user, err := decodeUser(ctx)
		//if err != nil {
		//	l.Err = err.Error()
		//}
		//l.UserID = user.ID

		RemoteIP(ctx)
	}
}

// RemoteIP ...
func RemoteIP(ctx *gin.Context) {
	host := ctx.GetHeader("REMOTE-HOST")
	if host == "" {
		host = ctx.Request.RemoteAddr
	}
	log.Info("host:", host)
}
