package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserDreamRouter(Router *gin.RouterGroup) {
	UserDreamRouter := Router.Group("jobuser").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UserDreamRouter.POST("createUserDream", api.CreateUserDream)   // 新建UserDream
		UserDreamRouter.POST("deleteUserDream", api.DeleteUserDream)   //删除UserDream
		UserDreamRouter.POST("updateUserDream", api.UpdateUserDream)   //更新UserDream
		UserDreamRouter.POST("findUserDream", api.FindUserDream)       // 根据ID获取UserDream
		UserDreamRouter.POST("getUserDreamList", api.GetUserDreamList) //获取UserDream列表
	}
	UserDreamRouterWx := Router.Group("dm")
	{
		UserDreamRouterWx.POST("createUserDream", api.CreateUserDream) // 新建UserDream
		UserDreamRouterWx.POST("deleteUserDream", api.DeleteUserDream) //删除UserDream
		UserDreamRouterWx.POST("updateUserDream", api.UpdateUserDream) //更新UserDream
		UserDreamRouterWx.POST("findUserDream", api.FindUserDream)     // 根据ID获取UserDream
		UserDreamRouterWx.POST("findUserDreambyOpenid", api.FindUserDreamByOpenid)
	}
}
