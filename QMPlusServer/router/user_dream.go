package router

import (
	"gin-vue-admin/controller/api"

	"github.com/gin-gonic/gin"
)

func InitUserDreamRouter(Router *gin.RouterGroup) {
	UserDreamRouter := Router.Group("dm")
	{
		UserDreamRouter.POST("createUserDream", api.CreateUserDream)   // 新建UserDream
		UserDreamRouter.POST("deleteUserDream", api.DeleteUserDream)   //删除UserDream
		UserDreamRouter.POST("updateUserDream", api.UpdateUserDream)   //更新UserDream
		UserDreamRouter.POST("findUserDream", api.FindUserDream)       // 根据ID获取UserDream
		UserDreamRouter.POST("getUserDreamList", api.GetUserDreamList) //获取UserDream列表
		UserDreamRouter.POST("findUserDreambyOpenid", api.FindUserDreamByOpenid)
	}
}
