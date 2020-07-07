package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"

	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserNewsRouter(Router *gin.RouterGroup) {
	UserNewsRouter := Router.Group("news").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UserNewsRouter.POST("createUserNews", api.CreateUserNews)   // 新建UserNews
		UserNewsRouter.POST("deleteUserNews", api.DeleteUserNews)   //删除UserNews
		UserNewsRouter.POST("updateUserNews", api.UpdateUserNews)   //更新UserNews
		UserNewsRouter.POST("findUserNews", api.FindUserNews)       // 根据ID获取UserNews
		UserNewsRouter.POST("getUserNewsList", api.GetUserNewsList) //获取UserNews列表
	}
	UserNewsRouterWx := Router.Group("un")
	{
		UserNewsRouterWx.POST("updateUserNewsCount", api.UpdateUserNewsCount)
		UserNewsRouterWx.POST("findUserNews", api.FindUserNews) // 根据ID获取UserNews
		UserNewsRouterWx.POST("getUserNewsListByType", api.GetUserNewsListByType)
	}
}
