package route

import (
	"github.com/gin-gonic/gin"
	"luck-admin/middleware"
	"luck-admin/pages"
	ginAdapter "github.com/GoAdminGroup/go-admin/adapter/gin"
)

func InitRoute(router *gin.Engine)  {
	//视图路由
	view := router.Group("/admin/view")
	{
		//用户管理页面
		view.GET("/user",  ginAdapter.Content(pages.UserView))
	}

	//api路由
	api := router.Group("admin/api")
	api.Use(middleware.CheckAuth())
	{
		api.GET("/user",ginAdapter.Content(pages.UserList))
	}

}
