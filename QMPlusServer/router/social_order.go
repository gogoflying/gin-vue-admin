package router

import (
	"gin-vue-admin/controller/api"
	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSocialOrderRouter(Router *gin.RouterGroup) {
	SocialOrderRouter := Router.Group("si") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SocialOrderRouter.POST("createSocialOrder", api.CreateSocialOrder)   // 新建SocialOrder
		SocialOrderRouter.POST("deleteSocialOrder", api.DeleteSocialOrder)   //删除SocialOrder
		SocialOrderRouter.POST("updateSocialOrder", api.UpdateSocialOrder)   //更新SocialOrder
		SocialOrderRouter.POST("findSocialOrder", api.FindSocialOrder)       // 根据ID获取SocialOrder
		SocialOrderRouter.POST("getSocialOrderList", api.GetSocialOrderList) //获取SocialOrder列表

		SocialOrderRouter.POST("paymentReq", api.PaymentReq)         //请求支付
		SocialOrderRouter.POST("confirmPayment", api.ConfirmPayment) //确认支付
		SocialOrderRouter.POST("cancelPayment", api.CancelPayment)   //取消支付
		SocialOrderRouter.POST("notifyResult", api.NotifyResult)     //通知支付结果
	}
}
