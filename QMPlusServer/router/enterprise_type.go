package router

import (
	"gin-vue-admin/controller/api"
	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitEnterpriseTypeRouter(Router *gin.RouterGroup) {
	EnterpriseTypeRouter := Router.Group("et") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		EnterpriseTypeRouter.POST("createEnterpriseType", api.CreateEnterpriseType)   // 新建EnterpriseType
		EnterpriseTypeRouter.POST("deleteEnterpriseType", api.DeleteEnterpriseType)   //删除EnterpriseType
		EnterpriseTypeRouter.POST("updateEnterpriseType", api.UpdateEnterpriseType)   //更新EnterpriseType
		EnterpriseTypeRouter.POST("findEnterpriseType", api.FindEnterpriseType)       // 根据ID获取EnterpriseType
		EnterpriseTypeRouter.POST("getEnterpriseTypeList", api.GetEnterpriseTypeList) //获取EnterpriseType列表
	}
}
