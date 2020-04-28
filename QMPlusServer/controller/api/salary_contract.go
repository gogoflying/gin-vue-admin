package api

import (
	"fmt"
	//"time"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userSalary"
	"gin-vue-admin/init/initlog"

	"github.com/gin-gonic/gin"
)

type StatusInfo int32
const (
	STATUS_UPLOAD_SOURCE_FILE StatusInfo = iota + 1
	STATUS_COMPOSE_SUCCESS
	STATUS_WRITE_NAME
	STATUS_DOWNLOAD
)
// @Tags SalaryContract
// @Summary 创建SalaryContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SalaryContract true "创建SalaryContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/createSalaryContract [post]
func CreateSalaryContract(c *gin.Context) {
	var un userSalary.SalaryContract
	_ = c.ShouldBindJSON(&un)
	err := un.CreateSalaryContract()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}


// @Tags SalaryContract
// @Summary 删除SalaryContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SalaryContract true "删除SalaryContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /un/deleteSalaryContract [post]
func DeleteSalaryContract(c *gin.Context) {
	var un userSalary.SalaryContract
	_ = c.ShouldBindJSON(&un)
	err := un.DeleteSalaryContract()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}


// @Tags SalaryContract
// @Summary 更新SalaryContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SalaryContract true "更新SalaryContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /un/updateSalaryContract [post]
func UpdateSalaryContract(c *gin.Context) {
	var un userSalary.SalaryContract
	_ = c.ShouldBindJSON(&un)
	err,reun := un.UpdateSalaryContract()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reun": reun,
		})
	}
}


// @Tags SalaryContract
// @Summary 用id查询SalaryContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SalaryContract true "用id查询SalaryContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /un/findSalaryContract [post]
func FindSalaryContract(c *gin.Context) {
	var un userSalary.SalaryContract
	_ = c.ShouldBindJSON(&un)
	err,reun := un.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reun": reun,
		})
	}
}


// @Tags SalaryContract
// @Summary 分页获取SalaryContract列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取SalaryContract列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/getSalaryContractList [post]
func GetSalaryContractList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userSalary.SalaryContract).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"userSalaryList": list,
			"total":          total,
			"page":           pageInfo.Page,
			"pageSize":       pageInfo.PageSize,
		})
	   /* response.Result(response.SUCCESS, gin.H{
            "userSalaryList": list,
            "total":    total,
            "page":     pageInfo.Page,
            "pageSize": pageInfo.PageSize,
        }, "获取数据成功", c)*/
	}
}

func DownloadContractList(c *gin.Context) {
	var uci userSalary.UserContractInfo
	_ = c.ShouldBindJSON(&uci)
	log.L.Info("UploadUserContract recv info:", uci)
	sc := userSalary.SalaryContract{
		Openid:uci.OpenId,
	}
	//var un userSalary.SalaryContract
	_ = c.ShouldBindJSON(&sc)
	err,resultSc := sc.FindById()
	jpgPathList := resultSc.Enter_contract_source_url
	fmt.Printf("get jpgPathList :%v",jpgPathList)
	
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"status": "ok",
			"jpgList":jpgPathList,
		})
	}
}

/*func Entersignature(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.UserWork).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"status": "ok",
		})
	}
}


func DownloadLeaveContract(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.UserWork).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"status": "ok",
		})
	}
}


func Leavingsignature(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.UserWork).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"status": "ok",
		})
	}
}*/

/*{
	"openId":"abc",
	"tmpContractPath":"./tmp/aa.pdf"
}*/
func UploadUserContract(c *gin.Context) {
	//,sysUser,localTmpPath string
	var uci userSalary.UserContractInfo
	_ = c.ShouldBindJSON(&uci)
	log.L.Info("UploadUserContract recv info:", uci)
	//write db
	sc := userSalary.SalaryContract{
		Openid:uci.OpenId,
		Status:int(STATUS_UPLOAD_SOURCE_FILE),
	}
	//var un userSalary.SalaryContract
	_ = c.ShouldBindJSON(&sc)
	log.L.Info("UploadUserContract new json:", sc)
	err := sc.CreateSalaryContract()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
	
	go runPDFConvert(c,uci.TmpContractPath,uci.OpenId)
	//return nil
}

func runPDFConvert(c *gin.Context,localPath,openId string) {
	//uci.tmpContractPath
	jpgPathList,err := servers.SplitPdf(localPath,openId,"tmp")
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("split contract pdf err ：%v", err), gin.H{})
	} 
	fmt.Printf("SplitPdf result :%v",jpgPathList)

	var  cloudContractPath string
	for _,jpgPath := range jpgPathList{
		err, filePath, _ := servers.UploadLocalFile(jpgPath, USER_HEADER_BUCKET, USER_HEADER_IMG_PATH)
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("接收返回值失败，%v", err), gin.H{})
		} else {
			//cloudContractPath = append(cloudContractPath,filePath)
			cloudContractPath += filePath
			cloudContractPath +=";"
		}
	}
	//start upload and write db
	sc := userSalary.SalaryContract{
		Openid:openId,
		Status:int(STATUS_COMPOSE_SUCCESS),
		Enter_contract_source_url:cloudContractPath,
	}
	//var un userSalary.SalaryContract
	_ = c.ShouldBindJSON(&sc)
	log.L.Info("UploadUserContract new json:", sc)
	err,_ = sc.UpdateSalaryContract()
	if err != nil {
		fmt.Printf("contract 更新失败：%v", err)
	} else {
		fmt.Printf("contract 更新成功")
	}
}
