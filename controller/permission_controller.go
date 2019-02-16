package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

// PermissionList godoc
// @Summary List permissions
// @Description List permissions
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @success 200 {array} model.Permission
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/permission [get]
func PermissionList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var permission model.Permission
		permissions, err := permission.Permissions()
		if err != nil {
			Error(ctx, err)
			return
		}
		log.Info(permissions)
		Success(ctx, permissions)
	}
}

// PermissionAdd godoc
// @Summary Add permission
// @Description Add permission
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param account body Permission true "permission update info"
// @success 200 {object} model.Permission
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/permission [post]
func PermissionAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var permission model.Permission
		err := ctx.BindJSON(&permission)
		if err != nil {
			Error(ctx, err)
			return
		}
		_, err = model.Insert(nil, &permission)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, permission)
	}
}

// PermissionUpdate godoc
// @Summary Update permission
// @Description Update permission
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Permission ID"
// @Param account body Permission true "permission update info"
// @success 200 {object} model.Permission
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/permission/{id} [post]
func PermissionUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		permission := model.NewPermission(id)
		b, err := permission.Get()
		if err != nil || !b {
			Error(ctx, xerrors.Errorf("no permissions:%w", err))
			return
		}
		err = ctx.BindJSON(permission)
		if err != nil {
			Error(ctx, err)
			return
		}

		_, err = model.Update(nil, id, permission)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, permission)
	}
}

// PermissionShow godoc
// @Summary Show permission
// @Description Show permission
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Permission ID"
// @success 200 {object} model.Permission
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/permission/{id} [get]
func PermissionShow(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		permission := model.NewPermission(id)
		_, err := model.Get(nil, permission)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, permission)
	}
}

// PermissionDelete godoc
// @Summary Delete permission
// @Description Delete permission
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Permission ID"
// @success 200 {object} model.Permission
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/permission/{id} [delete]
func PermissionDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		permission := model.NewPermission(id)
		_, err := model.Delete(nil, permission)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, permission)
	}
}

// PermissionRoleList godoc
// @Summary List role
// @Description List role
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Permission ID"
// @success 200 {array} model.Role
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/permission/{id}/role [get]
func PermissionRoleList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		permission := model.NewPermission(id)
		roles, err := permission.Roles()
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, roles)
	}
}

// PermissionUserList godoc
// @Summary List permission
// @Description List permission
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Permission ID"
// @success 200 {array} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/permission/{id}/user [get]
func PermissionUserList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		permission := model.NewPermission(id)
		users, err := permission.Users()
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, users)
	}
}
