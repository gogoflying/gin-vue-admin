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
	}
}
