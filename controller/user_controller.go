package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/config"
	"github.com/godcong/wego-auth-manager/model"
	"github.com/godcong/wego-auth-manager/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

// UserLogin godoc
// @Summary Login user
// @Description Login user
// @Tags default
// @Accept  json
// @Produce  json
// @Param account body Login true "user update info"
// @success 200 {object} util.WebToken
// @Failure 400 {object} controller.CodeMessage
// @Router /login [post]
func UserLogin(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var u model.Login
		e := ctx.BindJSON(&u)

		if e != nil {
			Error(ctx, e)
			return
		}
		user := model.User{
			Username: u.Username,
		}

		b, e := user.Get()
		if e != nil || !b {
			log.Error(e, b)
			Error(ctx, xerrors.New("user not found"))
			return
		}

		b = user.Validate(&u, config.Config().WebToken.Key)
		if !b {
			Error(ctx, xerrors.New("username password is not correct"))
			return
		}
		token, e := user.Login()
		if e != nil {
			Error(ctx, e)
			return
		}
		Success(ctx, gin.H{
			"token": token,
		})
		return
	}
}

// UserRegister godoc
// @Summary register user
// @Description register user
// @Tags default
// @Accept  json
// @Produce  json
// @Param account body User true "user update info"
// @success 200 {object} util.WebToken
// @Failure 400 {object} controller.CodeMessage
// @Router /register [post]
func UserRegister(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		err := ctx.BindJSON(&user)
		if err != nil {
			Error(ctx, err)
			return
		}
		user.Salt = util.GenerateRandomString(16)
		user.Password = util.SHA256(user.Password, config.Config().WebToken.Key, user.Salt)
		_, err = model.Insert(nil, &user)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, user)
	}
}

// UserList godoc
// @Summary List users
// @Description List users
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @success 200 {array} model.Paginate
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/user [get]
func UserList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		users, err := user.Users()
		if err != nil {
			Error(ctx, err)
			return
		}
		log.Info(users)
		Success(ctx, users)
	}
}

// UserAdd godoc
// @Summary Add user
// @Description Add user
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param account body User true "user update info"
// @success 200 {object} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/user [post]
func UserAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		err := ctx.BindJSON(&user)
		if err != nil {
			Error(ctx, err)
			return
		}
		user.Password = util.SHA256(user.Password, config.Config().WebToken.Key, util.GenerateRandomString(16))
		_, err = model.Insert(nil, &user)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, user)
	}
}

// UserUpdate godoc
// @Summary Update user
// @Description Update user
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @Param account body User true "user update info"
// @success 200 {object} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/user/{id} [post]
func UserUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.NewUser(id)
		b, err := user.Get()
		if err != nil || !b {
			Error(ctx, xerrors.Errorf("no users:%w", err))
			return
		}
		err = ctx.BindJSON(user)
		if err != nil {
			Error(ctx, err)
			return
		}

		_, err = model.Update(nil, id, user)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, user)
	}
}

// UserReset godoc
// @Summary reset user password
// @Description reset user password
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @Param account body User true "user update info"
// @success 200 {object} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/user/{id}/reset [post]
func UserReset(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.NewUser(id)
		b, err := user.Get()
		if err != nil || !b {
			Error(ctx, xerrors.Errorf("no users:%w", err))
			return
		}
		login := model.Login{}
		err = ctx.BindJSON(&login)
		if err != nil {
			Error(ctx, err)
			return
		}
		if login.Password == "" {
			login.Password = "123456"
		}

		user.Password = util.SHA256(login.Password, config.Config().WebToken.Key, user.Salt)
		_, err = model.UpdateWithColumn(nil, id, user, "password")
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, user)
	}
}

// UserShow godoc
// @Summary Show user
// @Description Show user
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @success 200 {object} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/user/{id} [get]
func UserShow(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.NewUser(id)
		_, err := model.Get(nil, user)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, user)
	}
}

// UserDelete godoc
// @Summary Delete user
// @Description Delete user
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @success 200 {object} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/user/{id} [delete]
func UserDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.NewUser(id)
		_, err := model.Delete(nil, user)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, user)
	}
}

// UserPermissionList godoc
// @Summary List permission
// @Description List permission
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @success 200 {array} model.Permission
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/user/{id}/permission [get]
func UserPermissionList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.NewUser(id)
		permissions, err := user.Permissions()
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, permissions)
	}
}

// UserRoleList godoc
// @Summary List role
// @Description List role
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @success 200 {array} model.Role
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/user/{id}/role [get]
func UserRoleList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.NewUser(id)
		roles, err := user.Roles()
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, roles)
	}
}

// UserRoleAdd godoc
// @Summary List role
// @Description List role
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @success 200 {array} model.RoleUser
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/user/{id}/role/{rid} [post]
func UserRoleAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		rid := ctx.Param("rid")
		user := model.NewUser(id)
		role := model.NewRole(rid)
		b, e := user.Get()
		if e != nil || !b {
			log.Error(e, b)
			Error(ctx, xerrors.New("no user"))
			return
		}
		b, e = role.Get()
		if e != nil || !b {
			log.Error(e, b)
			Error(ctx, xerrors.New("no role"))
			return
		}

		ru := model.RoleUser{
			RoleID: role.ID,
			UserID: user.ID,
		}
		i, e := model.Insert(nil, &ru)
		if e != nil || i == 0 {
			log.Error(e, i)
			Error(ctx, xerrors.New("insert role error"))
			return
		}
		ru.User = user
		ru.Role = role
		Success(ctx, &ru)
	}
}
