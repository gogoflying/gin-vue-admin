package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSalaryTemplatesRouter(Router *gin.RouterGroup) {
	SalaryTemplatesRouter := Router.Group("un") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SalaryTemplatesRouter.POST("createSalaryTemplates", api.CreateSalaryTemplates)                                                                         // 新建SalaryTemplates
		SalaryTemplatesRouter.POST("deleteSalaryTemplates", api.DeleteSalaryTemplates)                                                                         //删除SalaryTemplates
		SalaryTemplatesRouter.POST("updateSalaryTemplates", api.UpdateSalaryTemplates)                                                                         //更新SalaryTemplates
		SalaryTemplatesRouter.POST("findSalaryTemplates", api.FindSalaryTemplates)                                                                             // 根据ID获取SalaryTemplates
		SalaryTemplatesRouter.POST("getSalaryTemplatesList", middleware.JWTAuth(), middleware.EnterpriseHandler(), api.GetSalaryTemplatesList)                 //获取SalaryTemplates列表
		SalaryTemplatesRouter.POST("getSalaryTemplatesListBySearch", middleware.JWTAuth(), middleware.EnterpriseHandler(), api.GetSalaryTemplatesListBySearch) //获取SalaryTemplates列表
	}
}
