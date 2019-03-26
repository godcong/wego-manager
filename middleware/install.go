package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-manager/config"
)

// InstallCheck ...
func InstallCheck(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !config.IsInitialized() {
			//config.InitConfig()
		}
	}
}
