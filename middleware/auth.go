package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

// AuthCheck ...
func AuthCheck(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := User(ctx)
		role, err := user.Role()
		if err == nil {
			//超级管理员拥有所有权限
			if role.Slug == model.SlugGenesis {
				ctx.Next()
				return
			}
		}

		if user.Block {
			nop(ctx, "this account has been blocked")
			return
		}

		p := model.NewPermission()
		//logger := Logger(ctx)
		p.Slug = role.Slug
		err = p.Find()
		if err != nil {
			log.Println(err.Error())
			nop(ctx, err.Error())
			ctx.Abort()
			return
		}

		err = user.CheckPermission(p)
		if err != nil {
			log.Println(err.Error())
			nop(ctx, "this account has no permissions to visit this url")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
