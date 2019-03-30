package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-manager/model"
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

// Permission ...
type Permission struct {
	Version string
	Menu    string
}

// PermissionCheck ...
func PermissionCheck(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Debug(strings.Split(ctx.Request.URL.Path, "/"))
		var err error
		user := User(ctx)
		defer func() {
			if err != nil {
				Error(ctx, err)
				ctx.Abort()
				return
			}
		}()
		roles, err := user.Roles()
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

		if user.Block {
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

// ParseCRUD ...
func ParseCRUD(funcName string) string {
	n := len(funcName)
	switch {
	case strings.LastIndex(funcName, "List") == (n - 4):
		return "list"
	case strings.LastIndex(funcName, "Get") == (n - 3):
		return "list"
	case strings.LastIndex(funcName, "Add") == (n - 3):
		return "add"
	case strings.LastIndex(funcName, "Delete") == (n - 6):
		return "delete"
	case strings.LastIndex(funcName, "Update") == (n - 6):
		return "update"
	case strings.LastIndex(funcName, "Show") == (n - 6):
		return "list"
	}
	return ""
}
