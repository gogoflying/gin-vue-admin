package router

import (
	"gin-vue-admin/controller/api"
	//"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserWorkRouter(Router *gin.RouterGroup) {
	UserWorkRouter := Router.Group("wk")
	{
		UserWorkRouter.POST("createUserWork", api.CreateUserWork)   // 新建UserWork
		UserWorkRouter.POST("deleteUserWork", api.DeleteUserWork)   //删除UserWork
		UserWorkRouter.POST("updateUserWork", api.UpdateUserWork)   //更新UserWork
		UserWorkRouter.POST("findUserWork", api.FindUserWork)       // 根据ID获取UserWork
		UserWorkRouter.POST("getUserWorkList", api.GetUserWorkList) //获取UserWork列表
		UserWorkRouter.POST("getUserWorkListbyOpenid", api.GetUserWorkListByOpenid)
	}
}
