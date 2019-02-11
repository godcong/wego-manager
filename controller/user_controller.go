package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/config"
	"github.com/godcong/wego-auth-manager/model"
	"github.com/godcong/wego-auth-manager/util"
	"golang.org/x/exp/xerrors"
	"log"
)

// UserLogin godoc
// @Summary Login user
// @Description Login user
// @Tags default
// @Accept  json
// @Produce  json
// @Param account body UserLogin true "user update info"
// @success 200 {object} util.WebToken
// @Failure 400 {object} controller.CodeMessage
// @Router /login [post]
func UserLogin(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var u model.UserLogin
		err := ctx.BindJSON(&u)

		if err != nil {
			Error(ctx, err)
			return
		}
		user := model.User{
			Username: u.Username,
		}

		b, err := user.Get()
		if err != nil {
			log.Println("get error1")
			Error(ctx, err)
			return
		}

		if !b {
			log.Println("get error2")
			Error(ctx, xerrors.New("username password is not correct"))
			return
		}

		b = user.Validate(&u, config.Config().General.TokenKey)
		if !b {
			log.Println("validate error")
			Error(ctx, xerrors.New("username password is not correct"))
			return
		}
		token := util.NewWebToken(user.ID)
		token.Username = user.Username
		token.Nickname = user.Nickname
		t, err := util.ToToken(config.Config().General.TokenKey, token)
		if err != nil {
			log.Println(err)
			Error(ctx, xerrors.New("username password is not correct"))
			return
		}
		Success(ctx, gin.H{
			"token": t,
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
		user.Password = util.SHA256(user.Password, config.Config().General.TokenKey, user.Salt)
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
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Param current query string false "paginate:current"
// @Param limit query string false "paginate:limit"
// @Param order query string false "paginate:order"
// @success 200 {array} model.Paginate
// @Failure 400 {object} controller.CodeMessage
// @Router /dashboard/user [get]
func UserList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		users, err := user.Users()
		if err != nil {
			Error(ctx, err)
			return
		}
		log.Println(users)
		Success(ctx, users)
	}
}

// UserAdd godoc
// @Summary Add user
// @Description Add user
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param account body User true "user update info"
// @success 200 {object} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /dashboard/user [post]
func UserAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		err := ctx.BindJSON(&user)
		if err != nil {
			Error(ctx, err)
			return
		}
		user.Password = util.SHA256(user.Password, config.Config().General.TokenKey, util.GenerateRandomString(16))
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
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @Param account body User true "user update info"
// @success 200 {object} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /dashboard/user/{id} [post]
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

// UserShow godoc
// @Summary Show user
// @Description Show user
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @success 200 {object} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /dashboard/user/{id} [get]
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
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @success 200 {object} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /dashboard/user/{id} [delete]
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
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @success 200 {array} model.Permission
// @Failure 400 {object} controller.CodeMessage
// @Router /dashboard/user/{id}/permission [get]
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
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "User ID"
// @success 200 {array} model.Role
// @Failure 400 {object} controller.CodeMessage
// @Router /dashboard/user/{id}/role [get]
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
