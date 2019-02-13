package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

// RoleList godoc
// @Summary List roles
// @Description List roles
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @success 200 {array} model.Role
// @Failure 400 {object} controller.CodeMessage
// @Router /dashboard/role [get]
func RoleList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var role model.Role
		roles, err := role.Roles()
		if err != nil {
			Error(ctx, err)
			return
		}
		log.Info(roles)
		Success(ctx, roles)
	}
}

// RoleAdd godoc
// @Summary Add role
// @Description Add role
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param account body Role true "role update info"
// @success 200 {object} model.Role
// @Failure 400 {object} controller.CodeMessage
// @Router /dashboard/role [post]
func RoleAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var role model.Role
		err := ctx.BindJSON(&role)
		if err != nil {
			Error(ctx, err)
			return
		}
		_, err = model.Insert(nil, &role)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, role)
	}
}

// RoleUpdate godoc
// @Summary Update role
// @Description Update role
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Role ID"
// @Param account body Role true "role update info"
// @success 200 {object} model.Role
// @Failure 400 {object} controller.CodeMessage
// @Router /dashboard/role/{id} [post]
func RoleUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		role := model.NewRole(id)
		b, err := role.Get()
		if err != nil || !b {
			Error(ctx, xerrors.Errorf("no roles:%w", err))
			return
		}
		err = ctx.BindJSON(role)
		if err != nil {
			Error(ctx, err)
			return
		}

		_, err = model.Update(nil, id, role)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, role)
	}
}

// RoleShow godoc
// @Summary Show role
// @Description Show role
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Role ID"
// @success 200 {object} model.Role
// @Failure 400 {object} controller.CodeMessage
// @Router /dashboard/role/{id} [get]
func RoleShow(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		role := model.NewRole(id)
		_, err := model.Get(nil, role)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, role)
	}
}

// RoleDelete godoc
// @Summary Delete role
// @Description Delete role
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Role ID"
// @success 200 {object} model.Role
// @Failure 400 {object} controller.CodeMessage
// @Router /dashboard/role/{id} [delete]
func RoleDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		role := model.NewRole(id)
		_, err := model.Delete(nil, role)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, role)
	}
}

// RolePermissionList godoc
// @Summary List permission
// @Description List permission
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Role ID"
// @success 200 {array} model.Permission
// @Failure 400 {object} controller.CodeMessage
// @Router /dashboard/role/{id}/permission [get]
func RolePermissionList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		role := model.NewRole(id)
		permissions, err := role.Permissions()
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, permissions)
	}
}

// RoleUserList godoc
// @Summary List permission
// @Description List permission
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Role ID"
// @success 200 {array} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /dashboard/role/{id}/user [get]
func RoleUserList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		role := model.NewRole(id)
		users, err := role.Users()
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, users)
	}
}
