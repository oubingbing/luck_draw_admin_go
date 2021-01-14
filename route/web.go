package route

import (
	ginAdapter "github.com/GoAdminGroup/go-admin/adapter/gin"
	"github.com/gin-gonic/gin"
	"luck-admin/controllers"
	"luck-admin/middleware"
)

func InitRoute(router *gin.Engine)  {
	router.LoadHTMLGlob("html/*")

	//视图路由
	view := router.Group("/admin/view")
	{
		//用户-管理页面
		view.GET("/user",  ginAdapter.Content(controllers.UserView))
		//礼品-管理页面
		view.GET("/gift",  ginAdapter.Content(controllers.GiftView))
	}

	//api路由
	api := router.Group("admin/api")
	api.Use(middleware.CheckAuth())
	{
		//用户-分页
		api.GET("/user",controllers.UserList)

		//新建礼品
		api.POST("/gift/create",controllers.CreateGift)

		//cos token
		api.GET("/cos/token",controllers.GetCosToken)
	}

}
