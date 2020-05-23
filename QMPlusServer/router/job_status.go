package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitResumeStatusRouter(Router *gin.RouterGroup) {
	ResumeStatusRouter := Router.Group("jobstatus")
	{
		ResumeStatusRouter.POST("createResumeStatus", api.CreateResumeStatus)                                                         // 新建ResumeStatus
		ResumeStatusRouter.POST("deleteResumeStatus", api.DeleteResumeStatus)                                                         //删除ResumeStatus
		ResumeStatusRouter.POST("updateResumeStatus", api.UpdateResumeStatus)                                                         //更新ResumeStatus
		ResumeStatusRouter.POST("findResumeStatus", api.FindResumeStatus)                                                             // 根据ID获取ResumeStatus
		ResumeStatusRouter.POST("getResumeStatusList", middleware.JWTAuth(), middleware.EnterpriseHandler(), api.GetResumeStatusList) //获取ResumeStatus列表
		ResumeStatusRouter.POST("getResumeStatus", api.GetResumeByStatus)
		ResumeStatusRouter.POST("updateResumeMemo", api.UpdateResumeMemo)
	}
}
