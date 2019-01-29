package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/go-auth-manager/controller"
	"github.com/godcong/go-auth-manager/middleware"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func router(eng *gin.Engine) {
	version := "v0"
	eng.Use(middleware.VisitLog(version))
	eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r0 := eng.Group(version)
	r0.Use(middleware.AuthCheck(version))
	r0.POST("user", controller.UserAdd(version))
	r0.GET("user", controller.UserList(version))
	r0.POST("user/:id", controller.UserUpdate(version))
	r0.GET("user/:id", controller.UserShow(version))
	r0.DELETE("user/:id", controller.UserDelete(version))
	r0.GET("user/:id/permission", controller.UserPermissionList(version))
}

// Router ...
type Router struct {
	Name   string
	Handle func(string) gin.HandlerFunc
}

// Routers ...
type Routers struct {
	Group   string
	routers []*Router
}

// Register ...
func (r *Router) Register(eng *gin.Engine) {

}
