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
		//用户管理页面
		view.GET("/user",  ginAdapter.Content(controllers.UserView))
		/*view.GET("/user1", func(context *gin.Context) {
			context.HTML(http.StatusOK, "user.html", gin.H{
				"title": "Users",
			})
		})*/
	}

	//api路由
	api := router.Group("admin/api")
	api.Use(middleware.CheckAuth())
	{
		api.GET("/user",controllers.UserList)
	}

}
