package router

import (
	"gin-vue-admin/controller/api"

	"github.com/gin-gonic/gin"
)

func InitCitynameRouter(Router *gin.RouterGroup) {
	CitynameRouter := Router.Group("cityname")
	{
		CitynameRouter.POST("createCityname", api.CreateCityname)   // 新建Cityname
		CitynameRouter.POST("deleteCityname", api.DeleteCityname)   //删除Cityname
		CitynameRouter.POST("updateCityname", api.UpdateCityname)   //更新Cityname
		CitynameRouter.POST("findCityname", api.FindCityname)       // 根据ID获取Cityname
		CitynameRouter.POST("getCitynameList", api.GetCitynameList) //获取Cityname列表
	}
}
