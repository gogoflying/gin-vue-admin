package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitJoblistRouter(Router *gin.RouterGroup) {
	JoblistRouter := Router.Group("joblist").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.EnterpriseHandler())
	{
		JoblistRouter.POST("createJoblist", api.CreateJoblist) // 新建Joblist
		JoblistRouter.POST("deleteJoblist", api.DeleteJoblist) //删除Joblist
		JoblistRouter.POST("updateJoblist", api.UpdateJoblist) //更新Joblist
		JoblistRouter.POST("findJoblist", api.FindJoblist)
		JoblistRouter.POST("getjoblistOptions", api.GetJoblistOptions)
		JoblistRouter.POST("getJoblistList", api.GetJoblistListBackend) // 根据ID获取Joblist
	}
	JoblistRouterWx := Router.Group("jl")
	{
		JoblistRouterWx.POST("findJoblist", api.FindJoblist)
		JoblistRouterWx.POST("getJoblistList", api.GetJoblistList)
		JoblistRouterWx.POST("updateJoblistView", api.UpdateJoblistView)
		JoblistRouterWx.POST("getJoblistSearch", api.GetJoblistListSearch)
	}
}
