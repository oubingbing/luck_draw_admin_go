package middleware

import (
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/gin-gonic/gin"
	"luck-admin/enums"
	"luck-admin/util"
)

func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user,err := engine.User(ctx)
		if !err {
			util.ResponseJson(ctx ,enums.AUTH_NOT_LOGIN,enums.LoginRequestSessionErr.Error(),"")
			ctx.Abort()
			return
		}

		ctx.Set("user_id",user.Id)
		ctx.Set("user_name",user.Name)
		ctx.Set("user_avatar",user.Avatar)

		ctx.Next()
	}
}