package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"

	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSocialInsuranceRouter(Router *gin.RouterGroup) {
	SocialInsuranceRouter := Router.Group("social").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SocialInsuranceRouter.POST("createSocialInsurance", api.CreateSocialInsurance)   // 新建SocialInsurance
		SocialInsuranceRouter.POST("deleteSocialInsurance", api.DeleteSocialInsurance)   //删除SocialInsurance
		SocialInsuranceRouter.POST("updateSocialInsurance", api.UpdateSocialInsurance)   //更新SocialInsurance
		SocialInsuranceRouter.POST("findSocialInsurance", api.FindSocialInsurance)       // 根据ID获取SocialInsurance
		SocialInsuranceRouter.POST("getSocialInsuranceList", api.GetSocialInsuranceList) //获取SocialInsurance列表
		SocialInsuranceRouter.POST("getSocialInsuranceOptions", api.GetSocialInsuranceOptions)
	}
	SocialInsuranceRouterWx := Router.Group("si")
	{
		SocialInsuranceRouterWx.POST("createSocialInsurance", api.CreateSocialInsurance)   // 新建SocialInsurance
		SocialInsuranceRouterWx.POST("deleteSocialInsurance", api.DeleteSocialInsurance)   //删除SocialInsurance
		SocialInsuranceRouterWx.POST("updateSocialInsurance", api.UpdateSocialInsurance)   //更新SocialInsurance
		SocialInsuranceRouterWx.POST("findSocialInsurance", api.FindSocialInsurance)       // 根据ID获取SocialInsurance
		SocialInsuranceRouterWx.POST("getSocialInsuranceList", api.GetSocialInsuranceList) //获取SocialInsurance列表
		SocialInsuranceRouterWx.POST("getSocialInsuranceOptions", api.GetSocialInsuranceOptions)
	}
}
