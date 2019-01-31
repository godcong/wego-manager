package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/controller"
	"github.com/godcong/wego-auth-manager/middleware"
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
	r0.GET("user/:id/role", controller.UserRoleList(version))
	r0.GET("user/:id/permission", controller.UserPermissionList(version))

	r0.POST("role", controller.RoleAdd(version))
	r0.GET("role", controller.RoleList(version))
	r0.POST("role/:id", controller.RoleUpdate(version))
	r0.GET("role/:id", controller.RoleShow(version))
	r0.DELETE("role/:id", controller.RoleDelete(version))
	r0.GET("role/:id/permission", controller.RolePermissionList(version))
	r0.GET("role/:id/user", controller.RoleUserList(version))

	r0.POST("permission", controller.PermissionAdd(version))
	r0.GET("permission", controller.PermissionList(version))
	r0.POST("permission/:id", controller.PermissionUpdate(version))
	r0.GET("permission/:id", controller.PermissionShow(version))
	r0.DELETE("permission/:id", controller.PermissionDelete(version))
	r0.GET("permission/:id/role", controller.PermissionRoleList(version))
	r0.GET("permission/:id/user", controller.PermissionUserList(version))

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
