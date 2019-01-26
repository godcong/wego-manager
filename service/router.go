package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/go-auth-manager/controller"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func router(eng *gin.Engine) {
	version := "v0"
	eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r0 := eng.Group(version)
	r0.GET("user", controller.UserList(version))
}
