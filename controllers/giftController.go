package controllers

import (
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gin-gonic/gin"
	"luck-admin/enums"
	"luck-admin/models"
	"luck-admin/services"
	"luck-admin/util"
)

func GiftView(ctx *gin.Context) (types.Panel, error) {
	return types.Panel{
		Title:       "礼品管理",
		Description: "",
		Content:util.GetHtml("./html/gift.html"),
	}, nil
}

func CreateGift(ctx *gin.Context)  {
	var param enums.GiftParam
	errInfo := &enums.ErrorInfo{}
	if errInfo.Err = ctx.ShouldBind(&param); errInfo.Err != nil {
		util.ResponseJson(ctx,enums.ACTIVITY_PARAM_ERR,errInfo.Err.Error(),nil)
		return
	}

	db,connectErr := models.Connect()
	if connectErr != nil {
		util.ResponseJson(ctx,connectErr.Code,connectErr.Err.Error(),nil)
		return
	}

	var effect int64
	userId := 1
	effect,errInfo = services.SaveGift(db,userId,&param)
	if errInfo.Err != nil {
		util.ResponseJson(ctx,errInfo.Code,errInfo.Err.Error(),nil)
		return
	}

	if effect <= 0 {
		util.ResponseJson(ctx,enums.ACTIVITY_SAVE_ERR,enums.GiftSaveErr.Error(),nil)
		return
	}

	util.ResponseJson(ctx,enums.SUCCESS,"",effect)
	return
}
