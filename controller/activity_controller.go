package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

// ActivityList godoc
// @Summary List activities
// @Description List activities
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @success 200 {array} model.Activity
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/activity [get]
func ActivityList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var activity model.Activity
		activities, err := activity.Activities()
		if err != nil {
			Error(ctx, err)
			return
		}
		log.Info(activities)
		Success(ctx, activities)
	}
}

// ActivityAdd godoc
// @Summary Add activity
// @Description Add activity
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param account body Activity true "activity update info"
// @success 200 {object} model.Activity
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/activity [post]
func ActivityAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var activity model.Activity
		err := ctx.BindJSON(&activity)
		if err != nil {
			Error(ctx, err)
			return
		}
		_, err = model.Insert(nil, &activity)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, activity)
	}
}

// ActivityUpdate godoc
// @Summary Update activity
// @Description Update activity
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Activity ID"
// @Param account body Activity true "activity update info"
// @success 200 {object} model.Activity
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/activity/{id} [post]
func ActivityUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		activity := model.NewActivity(id)
		b, err := activity.Get()
		if err != nil || !b {
			Error(ctx, xerrors.Errorf("no activities:%w", err))
			return
		}
		err = ctx.BindJSON(activity)
		if err != nil {
			Error(ctx, err)
			return
		}

		_, err = model.Update(nil, id, activity)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, activity)
	}
}

// ActivityShow godoc
// @Summary Show activity
// @Description Show activity
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Activity ID"
// @success 200 {object} model.Activity
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/activity/{id} [get]
func ActivityShow(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		activity := model.NewActivity(id)
		_, err := model.Get(nil, activity)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, activity)
	}
}

// ActivityDelete godoc
// @Summary Delete activity
// @Description Delete activity
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Activity ID"
// @success 200 {object} model.Activity
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/activity/{id} [delete]
func ActivityDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		activity := model.NewActivity(id)
		_, err := model.Delete(nil, activity)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, activity)
	}
}

// ActivityUserList godoc
// @Summary List activity
// @Description List activity
// @Tags admin
// @Accept  json
// @Produce  json
// @Param token header string true "login token"
// @Param id path string true "Activity ID"
// @success 200 {array} model.User
// @Failure 400 {object} controller.CodeMessage
// @Router /admin/activity/{id}/user [get]
func ActivityUserList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		activity := model.NewActivity(id)
		user, err := activity.User()
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, user)
	}
}
