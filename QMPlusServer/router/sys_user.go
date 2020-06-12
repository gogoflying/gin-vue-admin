package router

import (
	"gin-vue-admin/controller/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	//.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UserRouter.POST("changePassword", api.ChangePassword)   // 修改密码
		UserRouter.POST("changeEmail", api.ChangeEmail)         // 修改邮箱
		UserRouter.POST("uploadHeaderImg", api.UploadHeaderImg) //上传头像
		UserRouter.POST("resetPassword", api.ResetPassword)
		UserRouter.POST("findSysUserById", api.FindUserById)
		UserRouter.POST("deleteSysUser", api.DeleteUser)
		UserRouter.POST("getUserList", api.GetUserList)           // 分页获取用户列表
		UserRouter.POST("setUserAuthority", api.SetUserAuthority) //设置用户权限
		UserRouter.POST("forgetSendEmail", api.SendEmailForget)
		UserRouter.POST("forgetVerify", api.VerifyEmailCode)
	}
}
