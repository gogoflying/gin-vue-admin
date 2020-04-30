package router

import (
	"gin-vue-admin/controller/api"

	"github.com/gin-gonic/gin"
)

func InitEnterpriseInfoRouter(Router *gin.RouterGroup) {
	EnterpriseInfoRouter := Router.Group("enterprise")
	{
		EnterpriseInfoRouter.POST("createEnterpriseInfo", api.CreateEnterpriseInfo)               // 新建EnterpriseInfo
		EnterpriseInfoRouter.POST("deleteEnterpriseInfo", api.DeleteEnterpriseInfo)               //删除EnterpriseInfo
		EnterpriseInfoRouter.POST("updateEnterpriseInfo", api.UpdateEnterpriseInfo)               //更新EnterpriseInfo
		EnterpriseInfoRouter.POST("findEnterpriseInfo", api.FindEnterpriseInfo)                   // 根据ID获取EnterpriseInfo
		EnterpriseInfoRouter.POST("findEnterpriseJobs", api.FindEnterpriseJobs)                   //获取该企业信息以及发布的职位
		EnterpriseInfoRouter.POST("getEnterpriseInfoList", api.GetEnterpriseInfoList)             //获取EnterpriseInfo列表
		EnterpriseInfoRouter.POST("getEnterpriseAllInfo", api.GetEnterpriseAllInfo)               //获取EnterpriseInfo列表
		EnterpriseInfoRouter.POST("getEnterpriseInfoListSearch", api.GetEnterpriseInfoListSearch) //根据cityid，keyword查询公司列表
	}
}
