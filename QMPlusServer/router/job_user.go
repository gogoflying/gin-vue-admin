package router

import (
	"gin-vue-admin/controller/api"

	"github.com/gin-gonic/gin"
)

func InitUsersRouter(Router *gin.RouterGroup) {
	UsersRouter := Router.Group("users")
	{
		UsersRouter.POST("createUsers", api.CreateUsers)             // 新建Users
		UsersRouter.POST("deleteUsers", api.DeleteUsers)             //删除Users
		UsersRouter.POST("updateUsers", api.UpdateUsers)             //更新Users
		UsersRouter.POST("findUsers", api.FindUsers)                 // 根据ID获取Users
		UsersRouter.POST("findUsersByOpenid", api.FindUsersByOpenid) // 根据ID获取Users
		UsersRouter.POST("getUsersList", api.GetUsersList)           //获取Users列表
		UsersRouter.POST("getUserOptions", api.GetUserOptions)
		UsersRouter.POST("jobUserLogin", api.JobUserLogin) //用户登陆
		//UsersRouter.POST("sendUserMobile", api.SendUserMobile)           //用户登陆
		UsersRouter.POST("decodeMobile", api.DecodeMobile)
		UsersRouter.POST("getEduLevels", api.GetEduLevels)
		UsersRouter.POST("getWorkTypes", api.GetWorkTypes)
	}
}
