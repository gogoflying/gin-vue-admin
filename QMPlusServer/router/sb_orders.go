package router

import (
	"gin-vue-admin/controller/api"
	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSbOrdersRouter(Router *gin.RouterGroup) {
	SbOrdersRouter := Router.Group("un") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SbOrdersRouter.POST("createSbOrders", api.CreateSbOrders)   // 新建SbOrders
		SbOrdersRouter.POST("deleteSbOrders", api.DeleteSbOrders)   //删除SbOrders
		SbOrdersRouter.POST("updateSbOrders", api.UpdateSbOrders)   //更新SbOrders
		SbOrdersRouter.POST("findSbOrders", api.FindSbOrders)       // 根据ID获取SbOrders
		SbOrdersRouter.POST("getSbOrdersList", api.GetSbOrdersList) //获取SbOrders列表
	}
}
