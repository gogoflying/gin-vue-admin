package router

import (
	"gin-vue-admin/controller/api"
	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitEduProfessionRouter(Router *gin.RouterGroup) {
	EduProfessionRouter := Router.Group("et") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		EduProfessionRouter.POST("createEduProfession", api.CreateEduProfession)   // 新建EduProfession
		EduProfessionRouter.POST("deleteEduProfession", api.DeleteEduProfession)   //删除EduProfession
		EduProfessionRouter.POST("updateEduProfession", api.UpdateEduProfession)   //更新EduProfession
		EduProfessionRouter.POST("findEduProfession", api.FindEduProfession)       // 根据ID获取EduProfession
		EduProfessionRouter.POST("getEduProfessionList", api.GetEduProfessionList) //获取EduProfession列表
	}
}
