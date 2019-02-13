package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/controller"
	"github.com/godcong/wego-auth-manager/middleware"
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
	eng.Use(middleware.VisitLog(l.Version))
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
	dashboard := v0.Group("dashboard", middleware.AuthCheck(l.Version), middleware.PermissionCheck(l.Version))
	//dashboard.Use(middleware.AuthCheck(l.Version), middleware.PermissionCheck(l.Version))

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

	l.Register(dashboard.POST, "user", controller.UserAdd)
	l.Register(dashboard.GET, "user", controller.UserList)
	l.Register(dashboard.POST, "user/:id", controller.UserUpdate)
	l.Register(dashboard.GET, "user/:id", controller.UserShow)
	l.Register(dashboard.DELETE, "user/:id", controller.UserDelete)
	l.Register(dashboard.GET, "user/:id/role", controller.UserRoleList)
	l.Register(dashboard.POST, "user/:id/role/:rid", controller.UserRoleAdd)
	l.Register(dashboard.GET, "user/:id/permission", controller.UserPermissionList)
	l.Register(dashboard.POST, "role", controller.RoleAdd)
	l.Register(dashboard.GET, "role", controller.RoleList)
	l.Register(dashboard.POST, "role/:id", controller.RoleUpdate)
	l.Register(dashboard.GET, "role/:id", controller.RoleShow)
	l.Register(dashboard.DELETE, "role/:id", controller.RoleDelete)
	l.Register(dashboard.GET, "role/:id/permission", controller.RolePermissionList)
	l.Register(dashboard.POST, "role/:id/permission/{pid}", controller.RolePermissionAdd)
	l.Register(dashboard.GET, "role/:id/user", controller.RoleUserList)
	l.Register(dashboard.POST, "permission", controller.PermissionAdd)
	l.Register(dashboard.GET, "permission", controller.PermissionList)
	l.Register(dashboard.POST, "permission/:id", controller.PermissionUpdate)
	l.Register(dashboard.GET, "permission/:id", controller.PermissionShow)
	l.Register(dashboard.DELETE, "permission/:id", controller.PermissionDelete)
	l.Register(dashboard.GET, "permission/:id/role", controller.PermissionRoleList)
	l.Register(dashboard.GET, "permission/:id/user", controller.PermissionUserList)

	user := v0.Group("user", middleware.AuthCheck(l.Version), middleware.PermissionCheck(l.Version))
	l.Register(user.GET, "property", controller.UserPropertyList)
	l.Register(user.POST, "property", controller.UserPropertyAdd)
	l.Register(user.POST, "property/:id", controller.UserPropertyUpdate)
	l.Register(user.DELETE, "property/:id", controller.UserPropertyDelete)
	l.Register(user.GET, "menu", controller.UserMenuList)
	l.Register(user.POST, "menu", controller.UserMenuAdd)
	l.Register(user.POST, "menu/:id", controller.UserMenuUpdate)
	l.Register(user.DELETE, "menu/:id", controller.UserMenuDelete)

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
