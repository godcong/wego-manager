package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/config"
)

// Install ...
func Install(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !config.IsInitialized() {
			//config.InitConfig()
		}
	}
}
