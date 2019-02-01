package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/config"
	"github.com/godcong/wego-auth-manager/model"
	"github.com/godcong/wego-auth-manager/util"
	"golang.org/x/exp/xerrors"
	"log"
	"strings"
)

// AuthCheck ...
func AuthCheck(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		var err error
		defer func() {
			if err != nil {
				Error(ctx, err)
				ctx.Abort()
				return
			}
		}()
		if token == "" {
			err = xerrors.New("token is null")
			return
		}
		t, err := util.FromToken(config.Config().General.TokenKey, token)
		if err != nil {
			return
		}
		log.Printf("%+v", t)

		user := model.User{}
		user.ID = t.UID
		b, err := user.Get()
		if err != nil {
			return
		}
		if !b {
			err = xerrors.New("no users")
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}

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
		user := User(ctx)
		roles, err := user.Roles()
		defer func() {
			if err != nil {
				Error(ctx, err)
				ctx.Abort()
				return
			}
		}()
		if err == nil {
			//超级管理员拥有所有权限
			for _, role := range roles {
				if role.Slug == model.RoleSlugAdmin {
					ctx.Next()
					return
				}
			}

		}

		if user.Block {
			err = xerrors.New("this account has been blocked")
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

// User ...
func User(ctx *gin.Context) *model.User {
	if v, b := ctx.Get("user"); b {
		if v0, b := v.(*model.User); b {
			log.Printf("%+v\n", v0)
			return v0
		}
	}
	return nil
}
