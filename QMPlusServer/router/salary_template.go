package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSalaryTemplatesRouter(Router *gin.RouterGroup) {
	SalaryTemplatesRouter := Router.Group("salary").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.EnterpriseHandler())
	{
		SalaryTemplatesRouter.POST("createSalaryTemplates", api.CreateSalaryTemplates)                   // 新建SalaryTemplates
		SalaryTemplatesRouter.POST("deleteSalaryTemplates", api.DeleteSalaryTemplates)                   //删除SalaryTemplates
		SalaryTemplatesRouter.POST("updateSalaryTemplates", api.UpdateSalaryTemplates)                   //更新SalaryTemplates
		SalaryTemplatesRouter.POST("findSalaryTemplates", api.FindSalaryTemplates)                       // 根据ID获取SalaryTemplates
		SalaryTemplatesRouter.POST("getSalaryTemplatesList", api.GetSalaryTemplatesList)                 //获取SalaryTemplates列表
		SalaryTemplatesRouter.POST("getSalaryTemplatesListBySearch", api.GetSalaryTemplatesListBySearch) //获取SalaryTemplates列表
	}
}
