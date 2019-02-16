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
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @success 200 {array} model.Role
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/role [get]
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
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param account body Role true "role update info"
// @success 200 {object} model.Role
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/role [post]
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
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Role ID"
// @Param account body Role true "role update info"
// @success 200 {object} model.Role
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/role/{id} [post]
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
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Role ID"
// @success 200 {object} model.Role
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/role/{id} [get]
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
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Role ID"
// @success 200 {object} model.Role
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/role/{id} [delete]
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
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Role ID"
// @success 200 {array} model.Permission
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/role/{id}/permission [get]
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

// RolePermissionAdd godoc
// @Summary add permission
// @Description add permission
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Role ID"
// @Param pid path string true "Permission ID"
// @success 200 {array} model.PermissionRole
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/role/{id}/permission/{pid} [get]
func RolePermissionAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		pid := ctx.Param("pid")
		role := model.NewRole(id)
		permission := model.NewPermission(pid)
		b, e := role.Get()
		if e != nil || !b {
			log.Error(e, b)
			Error(ctx, xerrors.New("no roleser"))
			return
		}
		b, e = permission.Get()
		if e != nil || !b {
			log.Error(e, b)
			Error(ctx, xerrors.New("no permission"))
			return
		}

		pr := model.PermissionRole{
			PermissionID: permission.ID,
			RoleID:       role.ID,
		}
		i, e := model.Insert(nil, &pr)
		if e != nil || i == 0 {
			log.Error(e, i)
			Error(ctx, xerrors.New("insert permission error"))
			return
		}
		pr.Permission = permission
		pr.Role = role
		Success(ctx, &pr)
	}
}

// RoleUserList godoc
// @Summary List permission
// @Description List permission
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Role ID"
// @success 200 {array} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/role/{id}/user [get]
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
