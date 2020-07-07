package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"

	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSalaryUsersRouter(Router *gin.RouterGroup) {
	SalaryUsersRouter := Router.Group("salary").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.EnterpriseHandler())
	{
		SalaryUsersRouter.POST("createSalaryUsers", api.CreateSalaryUsers)   // 新建SalaryUsers
		SalaryUsersRouter.POST("deleteSalaryUsers", api.DeleteSalaryUsers)   //删除SalaryUsers
		SalaryUsersRouter.POST("updateSalaryUsers", api.UpdateSalaryUsers)   //更新SalaryUsers
		SalaryUsersRouter.POST("findSalaryUsers", api.FindSalaryUsers)       // 根据ID获取SalaryUsers
		SalaryUsersRouter.POST("getSalaryUsersList", api.GetSalaryUsersList) //获取SalaryUsers列表
		SalaryUsersRouter.POST("importsalaryuser", api.ImportSalaryUsers)
		SalaryUsersRouter.Static("template", "./static/template")
	}
	SalaryUsersRouterWx := Router.Group("un")
	{
		SalaryUsersRouterWx.POST("findSalaryUsersByOpenid", api.FindSalaryUsersByOpenid)
		SalaryUsersRouterWx.POST("updateSalaryUsersLeaveStep", api.UpdateSalaryUsersLeaveStep)
		SalaryUsersRouterWx.POST("updateSalaryUsersEnterStep", api.UpdateSalaryUsersEnterStep)
		SalaryUsersRouterWx.POST("salaryUsersLogin", api.SalaryUserLoginMobile)
		SalaryUsersRouterWx.POST("updateSalaryUsersPassword", api.SalaryUserUpdatePassword)
		SalaryUsersRouterWx.POST("sendemail", api.SalaryUserSendEmail)
		SalaryUsersRouterWx.POST("sendsms", api.SalaryUserSendSMS)
		SalaryUsersRouterWx.POST("findpass", api.SalaryUserFindPass)
	}
}
