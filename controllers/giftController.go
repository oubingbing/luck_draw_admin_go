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
	if errInfo.Err = ctx.ShouldBindJSON(&param); errInfo.Err != nil {
		util.ResponseJson(ctx,enums.ACTIVITY_PARAM_ERR,errInfo.Err.Error(),nil)
		return
	}

	db,connectErr := models.Connect()
	defer db.Close()
	if connectErr != nil {
		util.ResponseJson(ctx,connectErr.Code,connectErr.Err.Error(),nil)
		return
	}

	userId := 1
	gift,errInfo := services.SaveGift(db,userId,&param)
	if errInfo != nil {
		util.ResponseJson(ctx,errInfo.Code,errInfo.Err.Error(),nil)
		return
	}

	domain,_ 	:= util.GetCosIni("cos_domain")
	giftResp,err := services.FormatGift(gift,domain)
	if err != nil {
		util.ResponseJson(ctx,err.Code,err.Err.Error(),nil)
		return
	}

	util.ResponseJson(ctx,enums.SUCCESS,"新建成功",giftResp)
	return
}

func GiftPage(ctx *gin.Context) {
	var param models.PageParam
	errInfo := &enums.ErrorInfo{}
	if errInfo.Err = ctx.ShouldBind(&param); errInfo.Err != nil {
		util.ResponseJson(ctx,enums.ACTIVITY_PARAM_ERR,errInfo.Err.Error(),nil)
		return
	}

	db,connectErr := models.Connect()
	defer db.Close()
	if connectErr != nil {
		util.ResponseJson(ctx,connectErr.Code,connectErr.Err.Error(),nil)
		return
	}

	list,err := services.PageGift(db,&param)
	if err != nil {
		util.ResponseJson(ctx,err.Code,err.Err.Error(),nil)
		return
	}

	util.ResponseJson(ctx,enums.SUCCESS,"ok",list)
	return
}

func GiftEnableList(ctx *gin.Context)  {
	db,connectErr := models.Connect()
	defer db.Close()
	if connectErr != nil {
		util.ResponseJson(ctx,connectErr.Code,connectErr.Err.Error(),nil)
		return
	}

	data,err := services.FindGiftEnable(db)
	if err != nil {
		util.ResponseJson(ctx,err.Code,err.Err.Error(),nil)
		return
	}

	util.ResponseJson(ctx,enums.SUCCESS,"ok",data)
	return
}
