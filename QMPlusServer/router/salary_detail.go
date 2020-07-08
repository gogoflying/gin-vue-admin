package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"

	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSalarysRouter(Router *gin.RouterGroup) {
	SalarysRouter := Router.Group("salary").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.EnterpriseHandler())
	{
		SalarysRouter.POST("createSalarys", api.CreateSalarys)           // 新建Salarys
		SalarysRouter.POST("deleteSalarys", api.DeleteSalarys)           //删除Salarys
		SalarysRouter.POST("batchdeleteSalarys", api.BatchDeleteSalarys) //批量删除Salarys
		SalarysRouter.POST("updateSalarys", api.UpdateSalarys)           //更新Salarys
		SalarysRouter.POST("findSalarys", api.FindSalarys)               // 根据ID获取Salarys
		SalarysRouter.POST("getSalarysList", api.GetSalarysList)         //获取Salarys列表
		SalarysRouter.POST("importsalarys", api.ImportSalarys)
	}
	SalarysRouterWx := Router.Group("un")
	{
		// 根据ID获取Salarys
		SalarysRouterWx.POST("findSalarysByIdAndOpenid", api.FindSalarysByIdAndOpenid)
		SalarysRouterWx.POST("getSalarysListSearch", api.GetSalarysListSearch)
	}
}
