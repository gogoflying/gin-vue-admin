package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitResumeStatusRouter(Router *gin.RouterGroup) {
	ResumeStatusRouter := Router.Group("jst").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.EnterpriseHandler())
	{
		ResumeStatusRouter.POST("createResumeStatus", api.CreateResumeStatus)   // 新建ResumeStatus
		ResumeStatusRouter.POST("deleteResumeStatus", api.DeleteResumeStatus)   //删除ResumeStatus
		ResumeStatusRouter.POST("updateResumeStatus", api.UpdateResumeStatus)   //更新ResumeStatus
		ResumeStatusRouter.POST("findResumeStatus", api.FindResumeStatus)       // 根据ID获取ResumeStatus
		ResumeStatusRouter.POST("getResumeStatusList", api.GetResumeStatusList) //获取ResumeStatus列表
		ResumeStatusRouter.POST("findUsersByOpenid", api.FindUsersByOpenid)     // 根据ID获取Users
		ResumeStatusRouter.POST("updateResumeMemo", api.UpdateResumeMemo)
		ResumeStatusRouter.POST("updateResumeRemark", api.UpdateResumeRemark)
	}
	ResumeStatusRouterWx := Router.Group("jobstatus")
	{
		ResumeStatusRouterWx.POST("createResumeStatus", api.CreateResumeStatus)
		ResumeStatusRouterWx.POST("getResumeStatus", api.GetResumeByStatus)
		ResumeStatusRouterWx.POST("jobEnterStatus", api.JobEnterStatus)
		ResumeStatusRouterWx.POST("badgeStatus", api.BadgeStatus)
	}
}
