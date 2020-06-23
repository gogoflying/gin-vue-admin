package router

import (
	"gin-vue-admin/controller/api"
	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSocialInsuranceRouter(Router *gin.RouterGroup) {
	SocialInsuranceRouter := Router.Group("si") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SocialInsuranceRouter.POST("createSocialInsurance", api.CreateSocialInsurance)   // 新建SocialInsurance
		SocialInsuranceRouter.POST("deleteSocialInsurance", api.DeleteSocialInsurance)   //删除SocialInsurance
		SocialInsuranceRouter.POST("updateSocialInsurance", api.UpdateSocialInsurance)   //更新SocialInsurance
		SocialInsuranceRouter.POST("findSocialInsurance", api.FindSocialInsurance)       // 根据ID获取SocialInsurance
		SocialInsuranceRouter.POST("getSocialInsuranceList", api.GetSocialInsuranceList) //获取SocialInsurance列表
		SocialInsuranceRouter.POST("getSocialInsuranceOptions", api.GetSocialInsuranceOptions)
	}
}
