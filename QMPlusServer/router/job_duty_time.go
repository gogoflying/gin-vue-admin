package router

import (
	"gin-vue-admin/controller/api"
	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitJobDutyTimeRouter(Router *gin.RouterGroup) {
	JobDutyTimeRouter := Router.Group("et") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		JobDutyTimeRouter.POST("createJobDutyTime", api.CreateJobDutyTime)   // 新建JobDutyTime
		JobDutyTimeRouter.POST("deleteJobDutyTime", api.DeleteJobDutyTime)   //删除JobDutyTime
		JobDutyTimeRouter.POST("updateJobDutyTime", api.UpdateJobDutyTime)   //更新JobDutyTime
		JobDutyTimeRouter.POST("findJobDutyTime", api.FindJobDutyTime)       // 根据ID获取JobDutyTime
		JobDutyTimeRouter.POST("getJobDutyTimeList", api.GetJobDutyTimeList) //获取JobDutyTime列表
	}
}
