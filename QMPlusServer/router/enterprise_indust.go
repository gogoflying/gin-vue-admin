package router

import (
	"gin-vue-admin/controller/api"
	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitEnterpriseIndustRouter(Router *gin.RouterGroup) {
	EnterpriseIndustRouter := Router.Group("et") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		EnterpriseIndustRouter.POST("createEnterpriseIndust", api.CreateEnterpriseIndust)   // 新建EnterpriseIndust
		EnterpriseIndustRouter.POST("deleteEnterpriseIndust", api.DeleteEnterpriseIndust)   //删除EnterpriseIndust
		EnterpriseIndustRouter.POST("updateEnterpriseIndust", api.UpdateEnterpriseIndust)   //更新EnterpriseIndust
		EnterpriseIndustRouter.POST("findEnterpriseIndust", api.FindEnterpriseIndust)       // 根据ID获取EnterpriseIndust
		EnterpriseIndustRouter.POST("getEnterpriseIndustList", api.GetEnterpriseIndustList) //获取EnterpriseIndust列表
	}
}
