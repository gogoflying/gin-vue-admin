package router

import (
	"gin-vue-admin/controller/api"
	//"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUsereducationRouter(Router *gin.RouterGroup) {
	UsereducationRouter := Router.Group("ed")
	{
		UsereducationRouter.POST("createUsereducation", api.CreateUsereducation)   // 新建Usereducation
		UsereducationRouter.POST("deleteUsereducation", api.DeleteUsereducation)   //删除Usereducation
		UsereducationRouter.POST("updateUsereducation", api.UpdateUsereducation)   //更新Usereducation
		UsereducationRouter.POST("findUsereducation", api.FindUsereducation)       // 根据ID获取Usereducation
		UsereducationRouter.POST("getUsereducationList", api.GetUsereducationList) //获取Usereducation列表
		UsereducationRouter.POST("getUsereducationListbyOpenid", api.GetUsereducationListByOpenid)
	}
}
