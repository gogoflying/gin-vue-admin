package router

import (
	"gin-vue-admin/controller/api"
	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitJobSalaryRouter(Router *gin.RouterGroup) {
	JobSalaryRouter := Router.Group("et") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		JobSalaryRouter.POST("createJobSalary", api.CreateJobSalary)   // 新建JobSalary
		JobSalaryRouter.POST("deleteJobSalary", api.DeleteJobSalary)   //删除JobSalary
		JobSalaryRouter.POST("updateJobSalary", api.UpdateJobSalary)   //更新JobSalary
		JobSalaryRouter.POST("findJobSalary", api.FindJobSalary)       // 根据ID获取JobSalary
		JobSalaryRouter.POST("getJobSalaryList", api.GetJobSalaryList) //获取JobSalary列表
	}
}
