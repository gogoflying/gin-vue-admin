package router

import (
	"gin-vue-admin/controller/api"
	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitEnterpriseCollectRouter(Router *gin.RouterGroup) {
	EnterpriseCollectRouter := Router.Group("jc") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		EnterpriseCollectRouter.POST("createEnterpriseCollect", api.CreateEnterpriseCollect)                   // 新建EnterpriseCollect
		EnterpriseCollectRouter.POST("deleteEnterpriseCollect", api.DeleteEnterpriseCollect)                   //删除EnterpriseCollect
		EnterpriseCollectRouter.POST("updateEnterpriseCollect", api.UpdateEnterpriseCollect)                   //更新EnterpriseCollect
		EnterpriseCollectRouter.POST("findEnterpriseCollect", api.FindEnterpriseCollect)                       // 根据ID获取EnterpriseCollect
		EnterpriseCollectRouter.POST("getEnterpriseCollectList", api.GetEnterpriseCollectList)                 //获取EnterpriseCollect列表
		EnterpriseCollectRouter.POST("getEnterpriseCollectListByOpenid", api.GetEnterpriseCollectListByOpenid) //获取EnterpriseCollect列表
	}
}
