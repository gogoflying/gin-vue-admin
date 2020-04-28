package router

import (
	"gin-vue-admin/controller/api"
	"github.com/gin-gonic/gin"
)

func InitUserAuthRouter(Router *gin.RouterGroup) {
	UserAuthRouter := Router.Group("userauth")
	{
		UserAuthRouter.POST("createUserAuth", api.CreateUserAuth)   // 新建UserAuth
		UserAuthRouter.POST("deleteUserAuth", api.DeleteUserAuth)   //删除UserAuth
		UserAuthRouter.POST("updateUserAuth", api.UpdateUserAuth)   //更新UserAuth
		UserAuthRouter.POST("findUserAuth", api.FindUserAuth)       // 根据ID获取UserAuth
		UserAuthRouter.POST("getUserAuthList", api.GetUserAuthList) //获取UserAuth列表
	}
}
