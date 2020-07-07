package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserBaseInfoRouter(Router *gin.RouterGroup) {
	UserBaseInfoRouter := Router.Group("jobuser").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UserBaseInfoRouter.POST("createUserBaseInfo", api.CreateUserBaseInfo)   // 新建UserBaseInfo
		UserBaseInfoRouter.POST("deleteUserBaseInfo", api.DeleteUserBaseInfo)   //删除UserBaseInfo
		UserBaseInfoRouter.POST("updateUserBaseInfo", api.UpdateUserBaseInfo)   //更新UserBaseInfo
		UserBaseInfoRouter.POST("findUserBaseInfo", api.FindUserBaseInfo)       // 根据ID获取UserBaseInfo
		UserBaseInfoRouter.POST("getUserBaseInfoList", api.GetUserBaseInfoList) //获取UserBaseInfo列表
	}
	UserBaseInfoRouterWx := Router.Group("bf")
	{
		UserBaseInfoRouterWx.POST("createUserBaseInfo", api.CreateUserBaseInfo) // 新建UserBaseInfo
		UserBaseInfoRouterWx.POST("updateUserBaseInfo", api.UpdateUserBaseInfo) //更新UserBaseInfo
		UserBaseInfoRouterWx.POST("findUserBaseInfo", api.FindUserBaseInfo)     // 根据ID获取UserBaseInfo
		UserBaseInfoRouterWx.POST("findUserBaseInfobyOpenid", api.FindUserBaseInfoByOpenid)
		UserBaseInfoRouterWx.POST("updateUserAvatarUrlbyOpenid", api.UpdateUserAvatarUrl)
	}
}
