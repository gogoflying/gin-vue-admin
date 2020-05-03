package router

import (
	"gin-vue-admin/controller/api"
	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSbConfRouter(Router *gin.RouterGroup) {
	SbConfRouter := Router.Group("un") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SbConfRouter.POST("createSbConf", api.CreateSbConf)   // 新建SbConf
		SbConfRouter.POST("deleteSbConf", api.DeleteSbConf)   //删除SbConf
		SbConfRouter.POST("updateSbConf", api.UpdateSbConf)   //更新SbConf
		SbConfRouter.POST("findSbConf", api.FindSbConf)       // 根据ID获取SbConf
		SbConfRouter.POST("getSbConfList", api.GetSbConfList) //获取SbConf列表
	}
}
