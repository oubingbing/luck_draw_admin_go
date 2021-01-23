package controllers

import (
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gin-gonic/gin"
	"luck-admin/enums"
	"luck-admin/models"
	"luck-admin/services"
	"luck-admin/util"
)

func UserView(ctx *gin.Context) (types.Panel, error) {
	return types.Panel{
		Title:       "用户管理",
		Description: "用户管理",
		Content:util.GetHtml("./html/user.html"),
	}, nil
}

func UserList(ctx *gin.Context)  {
	//userId,_:= ctx.Get("user_id")
	var param models.PageParam
	errInfo := &enums.ErrorInfo{}
	if errInfo.Err = ctx.ShouldBind(&param); errInfo.Err != nil {
		util.ResponseJson(ctx,enums.ACTIVITY_PARAM_ERR,errInfo.Err.Error(),nil)
		return
	}

	orm,err := models.Connect()
	defer orm.Close()
	if err != nil {
		util.ResponseJson(ctx,err.Code,err.Err.Error(),nil)
		return
	}

	list,err := services.UserPage(orm,&param)
	if err != nil {
		util.ResponseJson(ctx,err.Code,err.Err.Error(),nil)
		return
	}

	for _,item := range *list{
		services.FormatUserItem(item)
	}

	util.ResponseJson(ctx,enums.SUCCESS,"",list)
	return
}