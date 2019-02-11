package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
	"golang.org/x/exp/xerrors"
	"log"
)

// UserPropertyList godoc
// @Summary List user properties
// @Description List user properties
// @Tags property
// @Accept  json
// @Produce  json
// @Param current query string false "paginate:current"
// @Param limit query string false "paginate:limit"
// @Param order query string false "paginate:order"
// @success 200 {array} model.UserProperty
// @Failure 400 {object} controller.CodeMessage
// @Router /user/property [get]
func UserPropertyList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var property model.UserProperty
		properties, err := property.Properties()
		if err != nil {
			Error(ctx, err)
			return
		}
		log.Println(properties)
		Success(ctx, properties)
	}
}

// UserPropertyAdd godoc
// @Summary Add user property
// @Description Add user property
// @Tags property
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param account body Property true "property update info"
// @success 200 {object} model.UserProperty
// @Failure 400 {object} controller.CodeMessage
// @Router /user/property [post]
func UserPropertyAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var property model.UserProperty
		err := ctx.BindJSON(&property)
		if err != nil {
			Error(ctx, err)
			return
		}
		_, err = model.Insert(nil, &property)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, property)
	}
}

// UserPropertyUpdate godoc
// @Summary Update user property
// @Description Update user property
// @Tags property
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Property ID"
// @Param account body Property true "property update info"
// @success 200 {object} model.UserProperty
// @Failure 400 {object} controller.CodeMessage
// @Router /user/property/{pid} [post]
func UserPropertyUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("pid")
		property := model.NewUserProperty(id)
		property.UserID = User(ctx).ID
		b, err := property.Get()
		if err != nil || !b {
			Error(ctx, xerrors.Errorf("no property:%w", err))
			return
		}
		err = ctx.BindJSON(property)
		if err != nil {
			Error(ctx, err)
			return
		}

		_, err = model.Update(nil, id, property)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, property)
	}
}

// UserPropertyShow godoc
// @Summary Show user property
// @Description Show user property
// @Tags property
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Property ID"
// @success 200 {object} model.UserProperty
// @Failure 400 {object} controller.CodeMessage
// @Router /user/property/{pid} [get]
func UserPropertyShow(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		property := model.NewUserProperty(id)
		property.UserID = User(ctx).ID
		_, err := model.Get(nil, property)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, property)
	}
}

// UserPropertyDelete godoc
// @Summary Delete user property
// @Description Delete user property
// @Tags property
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Property ID"
// @success 200 {object} model.UserProperty
// @Failure 400 {object} controller.CodeMessage
// @Router /user/property/{pid} [delete]
func UserPropertyDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("pid")
		property := model.NewUserProperty(id)
		property.UserID = User(ctx).ID
		_, err := model.Delete(nil, property)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, property)
	}
}
