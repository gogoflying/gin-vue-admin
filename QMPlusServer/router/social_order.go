package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"

	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSocialOrderRouter(Router *gin.RouterGroup) {
	SocialOrderRouter := Router.Group("social").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SocialOrderRouter.POST("createSocialOrder", api.CreateSocialOrder)   // 新建SocialOrder
		SocialOrderRouter.POST("deleteSocialOrder", api.DeleteSocialOrder)   //删除SocialOrder
		SocialOrderRouter.POST("updateSocialOrder", api.UpdateSocialOrder)   //更新SocialOrder
		SocialOrderRouter.POST("findSocialOrder", api.FindSocialOrder)       // 根据ID获取SocialOrder
		SocialOrderRouter.POST("getSocialOrderList", api.GetSocialOrderList) //获取SocialOrder列表
	}
	SocialOrderRouterWx := Router.Group("si")
	{
		SocialOrderRouterWx.POST("createSocialOrder", api.CreateSocialOrder)   // 新建SocialOrder
		SocialOrderRouterWx.POST("deleteSocialOrder", api.DeleteSocialOrder)   //删除SocialOrder
		SocialOrderRouterWx.POST("updateSocialOrder", api.UpdateSocialOrder)   //更新SocialOrder
		SocialOrderRouterWx.POST("findSocialOrder", api.FindSocialOrder)       // 根据ID获取SocialOrder
		SocialOrderRouterWx.POST("getSocialOrderList", api.GetSocialOrderList) //获取SocialOrder列表

		SocialOrderRouterWx.POST("paymentReq", api.PaymentReq)         //请求支付
		SocialOrderRouterWx.POST("confirmPayment", api.ConfirmPayment) //确认支付
		SocialOrderRouterWx.POST("cancelPayment", api.CancelPayment)   //取消支付
		SocialOrderRouterWx.POST("notifyResult", api.NotifyResult)     //通知支付结果
		SocialOrderRouterWx.POST("addSocialOrder", api.AddSocialOrder)
		SocialOrderRouterWx.POST("getOrders", api.GetOrders)
	}
}
