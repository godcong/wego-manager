package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"strings"
)

func handleFuncName(ctx *gin.Context) string {
	hn := strings.Split(ctx.HandlerName(), ".")
	size := len(hn)
	if size < 2 {
		return ""
	}
	return hn[size-2]
}

// PermissionCheck ...
func PermissionCheck(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Debug(strings.Split(ctx.Request.URL.Path, "/"))
		var err error
		user := User(ctx)
		roles, err := user.Roles()
		defer func() {
			if err != nil {
				Error(ctx, err)
				ctx.Abort()
				return
			}
		}()
		log.Printf("%+v", roles)
		if err == nil {
			//超级管理员拥有所有权限
			for _, role := range roles {
				if role.Slug == model.RoleSlugAdmin {
					ctx.Next()
					return
				}
			}

		}

		if user.Enable {
			err = xerrors.New("this account is not enable")
			return
		}

		b := user.CheckPermission(handleFuncName(ctx))
		if !b {
			err = xerrors.New("no permission")
			return
		}
		ctx.Next()
	}
}
