package controllers

import (
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gin-gonic/gin"
	"luck-admin/enums"
	"luck-admin/models"
	"luck-admin/services"
	"luck-admin/util"
)

func ActivityView(ctx *gin.Context) (types.Panel, error) {
	return types.Panel{
		Title:       "活动管理",
		Description: "",
		Content:util.GetHtml("./html/activity.html"),
	}, nil
}

/**
 * 新增活动
 */
func CreateActivity(ctx *gin.Context)  {
	var param enums.ActivityCreateParam
	errInfo := &enums.ErrorInfo{}
	if errInfo.Err = ctx.ShouldBind(&param); errInfo.Err != nil {
		util.ResponseJson(ctx,enums.ACTIVITY_PARAM_ERR,errInfo.Err.Error(),nil)
		return
	}

	var effect int64
	db,connectErr := models.Connect()
	defer db.Close()
	if connectErr != nil {
		util.ResponseJson(ctx,connectErr.Code,connectErr.Err.Error(),nil)
		return
	}

	effect,errInfo = services.SaveActivity(db,&param)
	if errInfo.Err != nil {
		util.ResponseJson(ctx,errInfo.Code,errInfo.Err.Error(),nil)
		return
	}

	if effect <= 0 {
		util.ResponseJson(ctx,enums.ACTIVITY_SAVE_ERR,enums.CreateLDFail.Error(),nil)
		return
	}

	util.ResponseJson(ctx,enums.SUCCESS,"创建成功",effect)
	return
}

func ActivityPage(ctx *gin.Context)  {
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

	activities,err := services.ActivityPage(db,&param)
	if err != nil {
		util.ResponseJson(ctx,err.Code,err.Err.Error(),nil)
		return
	}

	util.ResponseJson(ctx,enums.SUCCESS,"",activities)
	return
}

/**
 * 获取详情
 */
func GetActivityDetail(ctx *gin.Context)  {
	id,ok := ctx.GetQuery("id")
	if !ok {
		util.ResponseJson(ctx,enums.ACTIVITY_DETAIL_PARAM_ERR,"参数不能为空",nil)
		return
	}

	db,connectErr := models.Connect()
	defer db.Close()
	if connectErr != nil {
		util.ResponseJson(ctx,connectErr.Code,connectErr.Err.Error(),nil)
		return
	}

	activity,err := services.ActivityDetail(db ,id)
	if err != nil {
		util.ResponseJson(ctx,err.Code,err.Err.Error(),nil)
		return
	}

	util.ResponseJson(ctx,enums.SUCCESS,"",activity)
	return
}

func DeleteActivity(ctx *gin.Context)  {
	id,ok := util.Input(ctx,"id")
	if !ok {
		util.ResponseJson(ctx,enums.ACTIVITY_DETAIL_PARAM_ERR,"参数不能为空",nil)
		return
	}

	float64Id,float64ok := id.(float64)
	if !float64ok {
		util.ResponseJson(ctx,enums.ACTIVITY_DETAIL_PARAM_ERR,"系统异常",nil)
		return
	}

	db,connectErr := models.Connect()
	defer db.Close()
	if connectErr != nil {
		util.ResponseJson(ctx,connectErr.Code,connectErr.Err.Error(),nil)
		return
	}

	deleteResult := services.ActivityDelete(db,uint(float64Id))
	if deleteResult != nil {
		util.ResponseJson(ctx,deleteResult.Code,deleteResult.Err.Error(),nil)
		return
	}

	util.ResponseJson(ctx,enums.SUCCESS,"删除成功",nil)
	return
}

func ChangeActivityStatus(ctx *gin.Context)  {
	data := util.InputAll(ctx)
	id,ok := data["id"]
	if !ok {
		util.ResponseJson(ctx,enums.ACTIVITY_DETAIL_PARAM_ERR,"id不能为空",nil)
		return
	}

	status,ok := data["status"]
	if !ok {
		util.ResponseJson(ctx,enums.ACTIVITY_DETAIL_PARAM_ERR,"status不能为空",nil)
		return
	}

	db,connectErr := models.Connect()
	defer db.Close()
	if connectErr != nil {
		util.ResponseJson(ctx,connectErr.Code,connectErr.Err.Error(),nil)
		return
	}

	err := services.ActivityUpdateStatus(db,id,status)
	if err != nil {
		util.ResponseJson(ctx,err.Code,err.Err.Error(),nil)
		return
	}

	util.ResponseJson(ctx,enums.SUCCESS,"操作成功",nil)
	return
}

func UpdateActivity(ctx *gin.Context)  {
	var param enums.ActivityUpdateParam
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

	activity,err := services.ActivityUpdate(db,&param)
	if err != nil {
		util.ResponseJson(ctx,err.Code,err.Err.Error(),nil)
		return
	}

	domain,_ 	:= util.GetCosIni("cos_domain")
	services.ActivityFormat(domain,activity)

	util.ResponseJson(ctx,enums.SUCCESS,"编辑成功",activity)
	return
}
