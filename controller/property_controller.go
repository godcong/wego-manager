package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

// PropertyList godoc
// @Summary List propertys
// @Description List propertys
// @Tags admin
// @Accept  json
// @Produce  json
// @Param current query string false "paginate:current"
// @Param limit query string false "paginate:limit"
// @Param order query string false "paginate:order"
// @success 200 {array} model.Property
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/property [get]
func PropertyList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var property model.Property
		properties, err := property.Properties()
		if err != nil {
			Error(ctx, err)
			return
		}
		log.Info(properties)
		Success(ctx, properties)
	}
}

// PropertyAdd godoc
// @Summary Add property
// @Description Add property
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param account body Property true "property update info"
// @success 200 {object} model.Property
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/property [post]
func PropertyAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var property model.Property
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

// PropertyUpdate godoc
// @Summary Update property
// @Description Update property
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Property ID"
// @Param account body Property true "property update info"
// @success 200 {object} model.Property
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/property/{id} [post]
func PropertyUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		property := model.NewProperty(id)
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

// PropertyShow godoc
// @Summary Show property
// @Description Show property
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Property ID"
// @success 200 {object} model.Property
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/property/{id} [get]
func PropertyShow(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		property := model.NewProperty(id)
		_, err := model.Get(nil, property)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, property)
	}
}

// PropertyDelete godoc
// @Summary Delete property
// @Description Delete property
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Property ID"
// @success 200 {object} model.Property
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/property/{id} [delete]
func PropertyDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		property := model.NewProperty(id)
		_, err := model.Delete(nil, property)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, property)
	}
}
