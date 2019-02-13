package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"strings"
)

// NotifyServer ...
func NotifyServer(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Info(ctx.Params)
		paths := strings.Split(ctx.Request.URL.Path, "/")
		if len(paths) < 6 {
			log.Info("path error", paths)
			Error(ctx, xerrors.New("path error"))
			return
		}
		var back model.Notify
		back.Ver = paths[1]
		back.Sign = paths[3]
		back.URI = paths[4]
		back.BackType = paths[5]
		b, e := model.Get(nil, &back)
		if e != nil || !b {
			log.Info("wrong back address", paths)
			Error(ctx, xerrors.New("wrong back address"))
			return
		}
		log.Info(ctx.HandlerName())
		log.Info(ctx.Request.URL.RawPath)
		log.Info(ctx.Request.URL.Path)
		log.Info(ctx.Request.URL.EscapedPath())

		Success(ctx, back)
		log.Info(back)
		return
	}
}
