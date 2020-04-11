package router

import (
	"gin-vue-admin/controller/api"
	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitJobWorkExpireRouter(Router *gin.RouterGroup) {
	JobWorkExpireRouter := Router.Group("et") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		JobWorkExpireRouter.POST("createJobWorkExpire", api.CreateJobWorkExpire)   // 新建JobWorkExpire
		JobWorkExpireRouter.POST("deleteJobWorkExpire", api.DeleteJobWorkExpire)   //删除JobWorkExpire
		JobWorkExpireRouter.POST("updateJobWorkExpire", api.UpdateJobWorkExpire)   //更新JobWorkExpire
		JobWorkExpireRouter.POST("findJobWorkExpire", api.FindJobWorkExpire)       // 根据ID获取JobWorkExpire
		JobWorkExpireRouter.POST("getJobWorkExpireList", api.GetJobWorkExpireList) //获取JobWorkExpire列表
	}
}
