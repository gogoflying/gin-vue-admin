package router

import (
	"gin-vue-admin/controller/api"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("register", api.Register)
		BaseRouter.POST("login", api.Login)
		BaseRouter.POST("captcha", api.Captcha)
		BaseRouter.GET("captcha/:captchaId", api.CaptchaImg)
		BaseRouter.POST("forgetSendEmail", api.SendEmailForget)
		BaseRouter.POST("forgetVerify", api.VerifyEmailCode)
	}
	return BaseRouter
}
