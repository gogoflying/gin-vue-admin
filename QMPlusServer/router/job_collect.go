package router

import (
	"gin-vue-admin/controller/api"
	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitJobCollectRouter(Router *gin.RouterGroup) {
	JobCollectRouter := Router.Group("jc") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		JobCollectRouter.POST("createJobCollect", api.CreateJobCollect)                   // 新建JobCollect
		JobCollectRouter.POST("deleteJobCollect", api.DeleteJobCollect)                   //删除JobCollect
		JobCollectRouter.POST("updateJobCollect", api.UpdateJobCollect)                   //更新JobCollect
		JobCollectRouter.POST("findJobCollect", api.FindJobCollect)                       // 根据ID获取JobCollect
		JobCollectRouter.POST("getJobCollectList", api.GetJobCollectList)                 //获取JobCollect列表
		JobCollectRouter.POST("getJobCollectListbyOpenid", api.GetJobCollectListByOpenid) //获取JobCollect列表
	}
}
