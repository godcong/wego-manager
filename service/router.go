package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-manager/controller"
	"github.com/godcong/wego-manager/middleware"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Handle ...
type Handle func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

// HandleFunc ...
type HandleFunc func(string) gin.HandlerFunc

// Router ...
type Router struct {
	Handle     Handle
	Name       string
	HandleFunc HandleFunc
}

// RouteLoader ...
type RouteLoader struct {
	Version string
	routers []*Router
}

// NewRouteLoader ...
func NewRouteLoader(version string) *RouteLoader {
	return &RouteLoader{Version: version}
}

func (l *RouteLoader) router(eng *gin.Engine) {
	eng.Use(middleware.InstallCheck(l.Version), middleware.VisitLog(l.Version))
	eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v0 := eng.Group(l.Version)

	notify := v0.Group("notify")
	notify.POST("/:sign/:backType/*uri", controller.NotifyServer(l.Version))

	payment := v0.Group("payment")
	payment.POST("/:sign/:payType", controller.NotifyPaymentUnify(l.Version))

	v0.POST("login", controller.UserLogin(l.Version))
	v0.POST("register", controller.UserRegister(l.Version))
	//超级管理员面板
	//账号、密码、所属组织、角色权限、邮箱、手机号码、授权证书和授权私钥
	admin := v0.Group("admin", middleware.AuthCheck(l.Version), middleware.PermissionCheck(l.Version))
	//admin.Use(middleware.AuthCheck(l.Version), middleware.PermissionCheck(l.Version))

	//r0.POST("user", controller.UserAdd(version))
	//r0.GET("user", controller.UserList(version))
	//r0.POST("user/:id", controller.UserUpdate(version))
	//r0.GET("user/:id", controller.UserShow(version))
	//r0.DELETE("user/:id", controller.UserDelete(version))
	//r0.GET("user/:id/role", controller.UserRoleList(version))
	//r0.GET("user/:id/permission", controller.UserPermissionList(version))
	//r0.POST("role", controller.RoleAdd(version))
	//r0.GET("role", controller.RoleList(version))
	//r0.POST("role/:id", controller.RoleUpdate(version))
	//r0.GET("role/:id", controller.RoleShow(version))
	//r0.DELETE("role/:id", controller.RoleDelete(version))
	//r0.GET("role/:id/permission", controller.RolePermissionList(version))
	//r0.GET("role/:id/user", controller.RoleUserList(version))
	//r0.POST("permission", controller.PermissionAdd(version))
	//r0.GET("permission", controller.PermissionList(version))
	//r0.POST("permission/:id", controller.PermissionUpdate(version))
	//r0.GET("permission/:id", controller.PermissionShow(version))
	//r0.DELETE("permission/:id", controller.PermissionDelete(version))
	//r0.GET("permission/:id/role", controller.PermissionRoleList(version))
	//r0.GET("permission/:id/user", controller.PermissionUserList(version))

	l.Register(admin.POST, "users", controller.UserAdd)
	l.Register(admin.GET, "users", controller.UserList)
	l.Register(admin.POST, "users/:id", controller.UserUpdate)
	l.Register(admin.GET, "users/:id", controller.UserShow)
	l.Register(admin.DELETE, "users/:id", controller.UserDelete)
	l.Register(admin.GET, "users/:id/roles", controller.UserRoleList)
	l.Register(admin.POST, "users/:id/roles/:rid", controller.UserRoleAdd)
	l.Register(admin.GET, "users/:id/permissions", controller.UserPermissionList)
	l.Register(admin.POST, "roles", controller.RoleAdd)
	l.Register(admin.GET, "roles", controller.RoleList)
	l.Register(admin.POST, "roles/:id", controller.RoleUpdate)
	l.Register(admin.GET, "roles/:id", controller.RoleShow)
	l.Register(admin.DELETE, "roles/:id", controller.RoleDelete)
	l.Register(admin.GET, "roles/:id/permissions", controller.RolePermissionList)
	l.Register(admin.POST, "roles/:id/permissions/:pid", controller.RolePermissionAdd)
	l.Register(admin.GET, "roles/:id/users", controller.RoleUserList)
	l.Register(admin.POST, "permissions", controller.PermissionAdd)
	l.Register(admin.GET, "permissions", controller.PermissionList)
	l.Register(admin.POST, "permissions/:id", controller.PermissionUpdate)
	l.Register(admin.GET, "permissions/:id", controller.PermissionShow)
	l.Register(admin.DELETE, "permissions/:id", controller.PermissionDelete)
	l.Register(admin.GET, "permissions/:id/roles", controller.PermissionRoleList)
	l.Register(admin.GET, "permissions/:id/users", controller.PermissionUserList)

	user := v0.Group("user", middleware.AuthCheck(l.Version), middleware.PermissionCheck(l.Version))
	l.Register(user.GET, "property", controller.UserPropertyList)
	l.Register(user.POST, "property", controller.UserPropertyAdd)
	l.Register(user.POST, "property/:id", controller.UserPropertyUpdate)
	l.Register(user.DELETE, "property/:id", controller.UserPropertyDelete)
	l.Register(user.GET, "menus", controller.UserMenuList)
	l.Register(user.POST, "menus", controller.UserMenuAdd)
	l.Register(user.POST, "menus/:id", controller.UserMenuUpdate)
	l.Register(user.DELETE, "menus/:id", controller.UserMenuDelete)

	for _, v := range l.routers {
		v.Handle(v.Name, v.HandleFunc(l.Version))
	}

}

// Register ...
func (l *RouteLoader) Register(handle Handle, name string, handleFunc HandleFunc) {
	l.routers = append(l.routers, &Router{
		Handle:     handle,
		Name:       name,
		HandleFunc: handleFunc,
	})
}

// Routers ...
func (l *RouteLoader) Routers() []*Router {
	return l.routers
}
