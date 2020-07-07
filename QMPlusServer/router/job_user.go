package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUsersRouter(Router *gin.RouterGroup) {
	UsersRouter := Router.Group("jobuser").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UsersRouter.POST("createUsers", api.CreateUsers) // 新建Users
		UsersRouter.POST("deleteUsers", api.DeleteUsers) //删除Users
		UsersRouter.POST("updateUsers", api.UpdateUsers) //更新Users
		UsersRouter.POST("findUsers", api.FindUsers)
		UsersRouter.POST("findJobUserById", api.FindJobUserById)         // 根据ID获取Users
		UsersRouter.POST("findUsersResumeInfo", api.FindUsersResumeInfo) // 根据ID获取简历信息
		UsersRouter.POST("getUsersList", api.GetUsersList)               //获取Users列表
		UsersRouter.POST("getUserOptions", api.GetUserOptions)
		UsersRouter.POST("getUserResumeList", api.GetUserResumeList) //获取User投递记录
		UsersRouter.GET("getContactInfo", api.GetContactInfo)
	}
	UsersRouterWx := Router.Group("users")
	{
		UsersRouterWx.POST("jobUserLogin", api.JobUserLogin) //用户登陆
		//UsersRouterWx.POST("sendUserMobile", api.SendUserMobile)           //用户登陆
		UsersRouterWx.POST("decodeMobile", api.DecodeMobile)
		UsersRouterWx.POST("updateUsers", api.UpdateUsers) //更新Users
		UsersRouterWx.POST("findUsersByOpenid", api.FindUsersByOpenid)
		UsersRouterWx.POST("getEduLevels", api.GetEduLevels)
		UsersRouterWx.POST("getWorkTypes", api.GetWorkTypes)
		UsersRouterWx.GET("getContactInfo", api.GetContactInfo)
	}
}
