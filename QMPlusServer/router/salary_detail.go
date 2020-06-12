package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"

	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSalarysRouter(Router *gin.RouterGroup) {
	SalarysRouter := Router.Group("un") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SalarysRouter.POST("createSalarys", middleware.JWTAuth(), middleware.EnterpriseHandler(), api.CreateSalarys) // 新建Salarys
		SalarysRouter.POST("deleteSalarys", api.DeleteSalarys)                                                       //删除Salarys
		SalarysRouter.POST("batchdeleteSalarys", api.BatchDeleteSalarys)                                             //批量删除Salarys
		SalarysRouter.POST("updateSalarys", api.UpdateSalarys)                                                       //更新Salarys
		SalarysRouter.POST("findSalarys", api.FindSalarys)                                                           // 根据ID获取Salarys
		SalarysRouter.POST("findSalarysByIdAndOpenid", api.FindSalarysByIdAndOpenid)
		SalarysRouter.POST("getSalarysList", middleware.JWTAuth(), middleware.EnterpriseHandler(), api.GetSalarysList) //获取Salarys列表
		SalarysRouter.POST("getSalarysListSearch", api.GetSalarysListSearch)
		//SalarysRouter.Static("template", "./static/template")
		SalarysRouter.POST("importsalarys", middleware.JWTAuth(), middleware.EnterpriseHandler(), api.ImportSalarys)
	}
}
