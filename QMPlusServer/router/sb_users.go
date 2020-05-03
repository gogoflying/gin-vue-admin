package router

import (
	"gin-vue-admin/controller/api"
	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSbUsersRouter(Router *gin.RouterGroup) {
	SbUsersRouter := Router.Group("un") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SbUsersRouter.POST("createSbUsers", api.CreateSbUsers)   // 新建SbUsers
		SbUsersRouter.POST("deleteSbUsers", api.DeleteSbUsers)   //删除SbUsers
		SbUsersRouter.POST("updateSbUsers", api.UpdateSbUsers)   //更新SbUsers
		SbUsersRouter.POST("findSbUsers", api.FindSbUsers)       // 根据ID获取SbUsers
		SbUsersRouter.POST("getSbUsersList", api.GetSbUsersList) //获取SbUsers列表
		SbUsersRouter.POST("findSbUsersByOpenid", api.FindSbUsersByOpenid)
	}
}
