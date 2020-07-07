package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"

	//"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUsereducationRouter(Router *gin.RouterGroup) {
	UsereducationRouter := Router.Group("jobuser").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UsereducationRouter.POST("createUsereducation", api.CreateUsereducation)   // 新建Usereducation
		UsereducationRouter.POST("deleteUsereducation", api.DeleteUsereducation)   //删除Usereducation
		UsereducationRouter.POST("updateUsereducation", api.UpdateUsereducation)   //更新Usereducation
		UsereducationRouter.POST("findUsereducation", api.FindUsereducation)       // 根据ID获取Usereducation
		UsereducationRouter.POST("getUsereducationList", api.GetUsereducationList) //获取Usereducation列表
	}
	UsereducationRouterWx := Router.Group("ed")
	{
		UsereducationRouterWx.POST("createUsereducation", api.CreateUsereducation) // 新建Usereducation
		UsereducationRouterWx.POST("deleteUsereducation", api.DeleteUsereducation) //删除Usereducation
		UsereducationRouterWx.POST("updateUsereducation", api.UpdateUsereducation) //更新Usereducation
		UsereducationRouterWx.POST("findUsereducation", api.FindUsereducation)     // 根据ID获取Usereducation
		UsereducationRouterWx.POST("getUsereducationListbyOpenid", api.GetUsereducationListByOpenid)
	}
}
