package router

import (
	"gin-vue-admin/controller/api"
	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSalaryUsersRouter(Router *gin.RouterGroup) {
	SalaryUsersRouter := Router.Group("un") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SalaryUsersRouter.POST("createSalaryUsers", api.CreateSalaryUsers)   // 新建SalaryUsers
		SalaryUsersRouter.POST("deleteSalaryUsers", api.DeleteSalaryUsers)   //删除SalaryUsers
		SalaryUsersRouter.POST("updateSalaryUsers", api.UpdateSalaryUsers)   //更新SalaryUsers
		SalaryUsersRouter.POST("findSalaryUsers", api.FindSalaryUsers)       // 根据ID获取SalaryUsers
		SalaryUsersRouter.POST("getSalaryUsersList", api.GetSalaryUsersList) //获取SalaryUsers列表
		SalaryUsersRouter.POST("findSalaryUsersByOpenid", api.FindSalaryUsersByOpenid)
		SalaryUsersRouter.POST("importsalaryuser", api.ImportSalaryUsers)
		SalaryUsersRouter.POST("updateSalaryUsersLeaveStep", api.UpdateSalaryUsersLeaveStep)
		SalaryUsersRouter.POST("updateSalaryUsersEnterStep", api.UpdateSalaryUsersEnterStep)
		SalaryUsersRouter.Static("template", "./static/template")
		SalaryUsersRouter.POST("salaryUsersLogin", api.SalaryUserLoginMobile)
		SalaryUsersRouter.POST("updateSalaryUsersPassword", api.SalaryUserUpdatePassword)
	}
}
