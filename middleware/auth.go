package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/config"
	"github.com/godcong/wego-auth-manager/model"
	"github.com/godcong/wego-auth-manager/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
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

		log.Info(strings.Split(token, "."))
		if user.Token != token {
			err = xerrors.New("login expired")
			return
		}

		ctx.Set("user", &user)
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
