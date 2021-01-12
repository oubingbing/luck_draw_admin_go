package pages

import (
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gin-gonic/gin"
	"luck-admin/util"
)

func UserView(ctx *gin.Context) (types.Panel, error) {
	return types.Panel{
		Title:       "用户管理",
		Description: "用户管理",
		Content:util.GetHtml("./html/user.html"),
	}, nil
}

func UserList(ctx *gin.Context) (types.Panel, error)  {
	//userId,_:= ctx.Get("user_id")
	//cfg := config.ReadFromJson("./config.json")


	util.ResponseJson(ctx,200,"成功",nil)
	ctx.Abort()

	return types.Panel{}, nil
}