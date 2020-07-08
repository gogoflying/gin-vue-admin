package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"

	//"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSalaryContractRouter(Router *gin.RouterGroup) {
	SalaryContractRouter := Router.Group("salary").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.EnterpriseHandler())
	{
		//SalaryContractRouter.POST("createSalaryContract", api.CreateSalaryContract)   // 新建Salary_contract
		//SalaryContractRouter.POST("deleteSalaryContract", api.DeleteSalaryContract)   //删除Salary_contract
		//SalaryContractRouter.POST("updateSalaryContract", api.UpdateSalaryContract)   //更新Salary_contract
		//SalaryContractRouter.POST("findSalaryContract", api.FindSalaryContract)       // 根据ID获取Salary_contract
		//SalaryContractRouter.POST("getSalaryContractList", api.GetSalaryContractList) //获取Salary_contract列表

		SalaryContractRouter.POST("downloadContractList", api.DownloadContractList)   // 新建UserNews
		SalaryContractRouter.POST("downloadLeaveContract", api.DownloadLeaveContract) //更新UserNews
		SalaryContractRouter.POST("uploadLeavingContract", api.UploadLeavingContract)
		SalaryContractRouter.POST("uploadUserContract", api.UploadUserContract) // to admin
		SalaryContractRouter.POST("importUserContract", api.ImportUserContract)
		SalaryContractRouter.POST("getContractJpgList", api.GetContractJpgList)
	}
	SalaryContractRouterWx := Router.Group("un") //.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		SalaryContractRouterWx.POST("downloadContractList", api.DownloadContractList)   // 新建UserNews
		SalaryContractRouterWx.POST("downloadLeaveContract", api.DownloadLeaveContract) //更新UserNews
		SalaryContractRouterWx.POST("writeSignatureJpg", api.WriteSignatureJpg)         //签名文件写入数据库，开始合成jpg
	}
}
