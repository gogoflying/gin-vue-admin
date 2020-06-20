package router

import (
	"gin-vue-admin/controller/api"
	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSocialUserRouter(Router *gin.RouterGroup) {
	SocialUserRouter := Router.Group("si") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SocialUserRouter.POST("createSocialUser", api.CreateSocialUser)   // 新建SocialUser
		SocialUserRouter.POST("deleteSocialUser", api.DeleteSocialUser)   //删除SocialUser
		SocialUserRouter.POST("updateSocialUser", api.UpdateSocialUser)   //更新SocialUser
		SocialUserRouter.POST("findSocialUser", api.FindSocialUser)       // 根据ID获取SocialUser
		SocialUserRouter.POST("getSocialUserList", api.GetSocialUserList) //获取SocialUser列表
	}
}
