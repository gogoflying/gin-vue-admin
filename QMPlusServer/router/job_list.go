package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitJoblistRouter(Router *gin.RouterGroup) {
	JoblistRouter := Router.Group("jl")
	{
		JoblistRouter.POST("createJoblist", middleware.JWTAuth(), middleware.CasbinHandler(), middleware.EnterpriseHandler(), api.CreateJoblist) // 新建Joblist
		JoblistRouter.POST("deleteJoblist", middleware.JWTAuth(), middleware.EnterpriseHandler(), api.DeleteJoblist)                             //删除Joblist
		JoblistRouter.POST("updateJoblist", middleware.JWTAuth(), middleware.EnterpriseHandler(), api.UpdateJoblist)                             //更新Joblist
		JoblistRouter.POST("findJoblist", api.FindJoblist)
		JoblistRouter.POST("getjoblistOptions", api.GetJoblistOptions)
		JoblistRouter.POST("getJoblistListBackend", middleware.JWTAuth(), middleware.EnterpriseHandler(), api.GetJoblistListBackend) // 根据ID获取Joblist
		JoblistRouter.POST("getJoblistList", api.GetJoblistList)                                                                     //获取Joblist列表
	}

}
