package router

import (
	"gin-vue-admin/controller/api"

	"github.com/gin-gonic/gin"
)

func InitUserBaseInfoRouter(Router *gin.RouterGroup) {
	UserBaseInfoRouter := Router.Group("bf")
	{
		UserBaseInfoRouter.POST("createUserBaseInfo", api.CreateUserBaseInfo)   // 新建UserBaseInfo
		UserBaseInfoRouter.POST("deleteUserBaseInfo", api.DeleteUserBaseInfo)   //删除UserBaseInfo
		UserBaseInfoRouter.POST("updateUserBaseInfo", api.UpdateUserBaseInfo)   //更新UserBaseInfo
		UserBaseInfoRouter.POST("findUserBaseInfo", api.FindUserBaseInfo)       // 根据ID获取UserBaseInfo
		UserBaseInfoRouter.POST("getUserBaseInfoList", api.GetUserBaseInfoList) //获取UserBaseInfo列表
		UserBaseInfoRouter.POST("findUserBaseInfobyOpenid", api.FindUserBaseInfoByOpenid)
		UserBaseInfoRouter.POST("updateUserAvatarUrlbyOpenid", api.UpdateUserAvatarUrl)
	}
}
