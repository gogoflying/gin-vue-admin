package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"

	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSocialUserRouter(Router *gin.RouterGroup) {
	SocialUserRouter := Router.Group("social").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SocialUserRouter.POST("createSocialUser", api.CreateSocialUser)   // 新建SocialUser
		SocialUserRouter.POST("deleteSocialUser", api.DeleteSocialUser)   //删除SocialUser
		SocialUserRouter.POST("updateSocialUser", api.UpdateSocialUser)   //更新SocialUser
		SocialUserRouter.POST("findSocialUser", api.FindSocialUser)       // 根据ID获取SocialUser
		SocialUserRouter.POST("getSocialUserList", api.GetSocialUserList) //获取SocialUser列表
	}
	SocialUserRouterWx := Router.Group("si")
	{
		SocialUserRouterWx.POST("createSocialUser", api.CreateSocialUser)   // 新建SocialUser
		SocialUserRouterWx.POST("deleteSocialUser", api.DeleteSocialUser)   //删除SocialUser
		SocialUserRouterWx.POST("updateSocialUser", api.UpdateSocialUser)   //更新SocialUser
		SocialUserRouterWx.POST("findSocialUser", api.FindSocialUser)       // 根据ID获取SocialUser
		SocialUserRouterWx.POST("getSocialUserList", api.GetSocialUserList) //获取SocialUser列表

		SocialUserRouterWx.POST("socialUserLogin", api.SocialUserLogin)
		SocialUserRouterWx.POST("getSocialUserByOpenid", api.FindSocialUserByOpenid)
		SocialUserRouterWx.POST("updateUserByOpenid", api.UpdateUserByOpenid)
	}
}
