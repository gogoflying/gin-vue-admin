package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"

	//"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserWorkRouter(Router *gin.RouterGroup) {
	UserWorkRouter := Router.Group("jobuser").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UserWorkRouter.POST("createUserWork", api.CreateUserWork)   // 新建UserWork
		UserWorkRouter.POST("deleteUserWork", api.DeleteUserWork)   //删除UserWork
		UserWorkRouter.POST("updateUserWork", api.UpdateUserWork)   //更新UserWork
		UserWorkRouter.POST("findUserWork", api.FindUserWork)       // 根据ID获取UserWork
		UserWorkRouter.POST("getUserWorkList", api.GetUserWorkList) //获取UserWork列表
	}
	UserWorkRouterWx := Router.Group("wk")
	{
		UserWorkRouterWx.POST("createUserWork", api.CreateUserWork) // 新建UserWork
		UserWorkRouterWx.POST("deleteUserWork", api.DeleteUserWork) //删除UserWork
		UserWorkRouterWx.POST("updateUserWork", api.UpdateUserWork) //更新UserWork
		UserWorkRouterWx.POST("findUserWork", api.FindUserWork)     // 根据ID获取UserWork
		UserWorkRouterWx.POST("getUserWorkListbyOpenid", api.GetUserWorkListByOpenid)
	}
}
