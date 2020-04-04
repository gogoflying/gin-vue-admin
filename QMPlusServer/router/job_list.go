package router

import (
	"gin-vue-admin/controller/api"

	"github.com/gin-gonic/gin"
)

func InitJoblistRouter(Router *gin.RouterGroup) {
	JoblistRouter := Router.Group("joblist")
	{
		JoblistRouter.POST("createJoblist", api.CreateJoblist)   // 新建Joblist
		JoblistRouter.POST("deleteJoblist", api.DeleteJoblist)   //删除Joblist
		JoblistRouter.POST("updateJoblist", api.UpdateJoblist)   //更新Joblist
		JoblistRouter.POST("findJoblist", api.FindJoblist)       // 根据ID获取Joblist
		JoblistRouter.POST("getJoblistList", api.GetJoblistList) //获取Joblist列表
	}

}
