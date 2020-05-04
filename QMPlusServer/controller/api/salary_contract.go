package api

import (
	"fmt"
	"io"
	"os"
	"time"
	"strings"

	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/initlog"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userSalary"

	"github.com/gin-gonic/gin"
)

type StatusInfo int32

const (
	STATUS_UPLOAD_SOURCE_FILE StatusInfo = iota + 1
	STATUS_COMPOSE_SUCCESS
	STATUS_WRITE_NAME
	STATUS_MERGED
)

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

func UpdateSalaryContract(c *gin.Context) {
	var un userSalary.SalaryContract
	_ = c.ShouldBindJSON(&un)
	err, reun := un.UpdateSalaryContract()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reun": reun,
		})
	}
}

func FindSalaryContract(c *gin.Context) {
	var un userSalary.SalaryContract
	_ = c.ShouldBindJSON(&un)
	err, reun := un.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reun": reun,
		})
	}
}

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

	}
}

func DownloadContractList(c *gin.Context) {
	var uci userSalary.UserContractInfo
	_ = c.ShouldBindJSON(&uci)
	log.L.Info("UploadUserContract recv info:", uci)
	sc := userSalary.SalaryContract{
		Openid: uci.OpenId,
	}
	//var un userSalary.SalaryContract
	_ = c.ShouldBindJSON(&sc)
	err, resultSc := sc.FindById()
	jpgPathList := resultSc.Enter_contract_source_url
	fmt.Printf("get jpgPathList :%v", jpgPathList)

	jpgPathArr := strings.Split(jpgPathList, ";")
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"status":  "ok",
			"jpgList": jpgPathArr,
		})
	}
}

func WriteSignatureJpg(c *gin.Context) {
	var usi userSalary.UserSignatureInfo
	_ = c.ShouldBindJSON(&usi)
	log.L.Info("WriteSignatureJpg recv info:", usi)
	if len(usi.OpenId) == 0{
		servers.ReportFormat(c, false, fmt.Sprintf("WriteSignatureJpg faild "), gin.H{})
		return
	}
	sc := userSalary.SalaryContract{
		Openid: usi.OpenId,
		User_signature_url:usi.SignatureFileUrl,
		Status:int(STATUS_WRITE_NAME),
	}
	_ = c.ShouldBindJSON(&sc)
	log.L.Info("WriteSignatureJpg new json:", sc)
	err, _ := sc.UpdateSalaryContract()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("contract 更新失败 :%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, fmt.Sprintf("contract 更新成功 :%v", err), gin.H{})
	}

	go runMergeImg(c,usi.SignatureFileUrl,usi.OpenId)

}

func DownloadLeaveContract(c *gin.Context) {
	var usi userSalary.UserSignatureInfo
	_ = c.ShouldBindJSON(&usi)
	log.L.Info("DownloadLeaveContract recv info:", usi)
	if len(usi.OpenId) == 0{
		servers.ReportFormat(c, false, fmt.Sprintf("WriteSignatureJpg faild "), gin.H{})
		return
	}
	sc := userSalary.SalaryContract{
		Openid: usi.OpenId,
	}
	_ = c.ShouldBindJSON(&sc)
	log.L.Info("WriteSignatureJpg new json:", sc)
	err, resultPath := sc.FindById()
	if len(resultPath.Leave_contract_target_url) <0{
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败：%v", err), gin.H{})
		return
	}
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("query DownloadLeaveContract openid:%s  :%v",usi.OpenId, err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"status":  "ok",
			"leaveJpg": resultPath.Leave_contract_target_url,
		})
	}
}

func UploadLeavingContract(c *gin.Context) {
	var uci userSalary.UserContractInfo
	_ = c.ShouldBindJSON(&uci)
	log.L.Info("UploadUserContract recv info:", uci)
	//write db
	sc := userSalary.SalaryContract{
		Openid: uci.OpenId,
		Leave_contract_target_url: uci.TmpContractPath,
	}
	//var un userSalary.SalaryContract
	_ = c.ShouldBindJSON(&sc)
	log.L.Info("UploadUserContract new json:", sc)
	err ,_:= sc.UpdateSalaryContract()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

func UploadUserContract(c *gin.Context) {
	//,sysUser,localTmpPath string
	var uci userSalary.UserContractInfo
	_ = c.ShouldBindJSON(&uci)
	log.L.Info("UploadUserContract recv info:", uci)
	//write db
	sc := userSalary.SalaryContract{
		Openid: uci.OpenId,
		Status: int(STATUS_UPLOAD_SOURCE_FILE),
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

	go runPDFConvert(c, uci.TmpContractPath, uci.OpenId)
	//return nil
}

func runPDFConvert(c *gin.Context, localPath, openId string) {
	//uci.tmpContractPath
	fmt.Printf("runPDFConvert :%s--%s\n", localPath,openId)
	jpgPathList, err := servers.SplitPdf(localPath, openId, "tmp")
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("split contract pdf err ：%v", err), gin.H{})
	}
	fmt.Printf("SplitPdf result :%v", jpgPathList)

	var cloudContractPath string
	for idx, jpgPath := range jpgPathList {
		err, filePath, _ := servers.UploadLocalFile(jpgPath, USER_HEADER_BUCKET, USER_HEADER_IMG_PATH)
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("接收返回值失败，%v", err), gin.H{})
		} else {
			//cloudContractPath = append(cloudContractPath,filePath)
			cloudContractPath += filePath
			if idx != len(jpgPathList) -1  {
				cloudContractPath += ";"
			}

			servers.DelLocalFile(jpgPath)
		}
	}
	//start upload and write db
	sc := userSalary.SalaryContract{
		Openid:                    openId,
		Status:                    int(STATUS_COMPOSE_SUCCESS),
		Enter_contract_source_url: cloudContractPath,
	}
	//var un userSalary.SalaryContract
	_ = c.ShouldBindJSON(&sc)
	log.L.Info("UploadUserContract new json:", sc)
	err, _ = sc.UpdateSalaryContract()
	if err != nil {
		fmt.Printf("contract 更新失败：%v", err)
	} else {
		fmt.Printf("contract 更新成功")
	}
}

func ImportUserContract(c *gin.Context) {
	openid := c.Query("openid")
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("上传文件失败，%v", err), gin.H{})
	} else {
		fileName := time.Now().Local().Format("20060102150405") + header.Filename
		dst, err := os.Create(fileName)
		if err != nil {
			// ignore
		}
		_, err = io.Copy(dst, file)
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("接收返回值失败，%v", err), gin.H{})
		} else {
			servers.ReportFormat(c, true, "上传成功", gin.H{
				"filepath": fileName,
				"openid":   openid,
			})
		}
	}
}

func GetContractJpgList(c *gin.Context) {
	//,sysUser,localTmpPath string
	var uci userSalary.UserContractInfo
	_ = c.ShouldBindJSON(&uci)
	log.L.Info("UploadUserContract recv info:", uci)

	sc := userSalary.SalaryContract{
		Openid: uci.OpenId,
	}
	//var un userSalary.SalaryContract
	_ = c.ShouldBindJSON(&sc)
	err, reun := sc.FindById()
	//log.L.Info("UploadUserContract new json:", reun.Enter_contract_source_url)
	if len(reun.Enter_contract_source_url) <0{
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败：%v", err), gin.H{})
		return
	}
	resultJpg := strings.Split(reun.Enter_contract_source_url, ";")
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"status":  "ok",
			"contartList": resultJpg,
		})
	}
}

func runMergeImg(c *gin.Context, imgPath,openId string) {
	//uci.tmpContractPath
	fmt.Printf("runMergeImg :%s\n", imgPath)
	err,localFile := servers.DownLoadLocalFile(USER_HEADER_BUCKET,imgPath,"tmp")
	if err != nil {
		fmt.Printf("DownLoadLocalFile  file :%s err:%v\n", imgPath,err)
		servers.ReportFormat(c, false, fmt.Sprintf("DownLoadLocalFile  file  err"), gin.H{})
		return
	} 
	sourcePdf := getContractLastFile(c,openId)
	if len(sourcePdf) == 0 {
		fmt.Printf("DownLoadLocalFile  file :%s err\n", sourcePdf)
		servers.ReportFormat(c, false, fmt.Sprintf("DownLoadLocalFile  file  err"), gin.H{})
		return
	} 

	signatureShrinkNmaePng,err := servers.ImgShrink(localFile,openId,"tmp")
	if err != nil{
		servers.ReportFormat(c, false, fmt.Sprintf("ImgShrink  file  err:%v",err), gin.H{})
		return
	}

	//fmt.Printf("DownLoadLocalFile  file003 :%s err\n", signatureShrinkNmaePng)
	mergedFile,err := servers.MergeImage(sourcePdf, signatureShrinkNmaePng,openId, "tmp")
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("MergeImage contract pdf err :%v",err), gin.H{})
		return
	}
	fmt.Printf("SplitPdf result :%v", mergedFile)

	err, cloudContractPath, _ := servers.UploadLocalFile(mergedFile, USER_HEADER_BUCKET, USER_HEADER_IMG_PATH)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("UploadLocalFile %v", err), gin.H{})
	} else {
		servers.DelLocalFile(imgPath)
		servers.DelLocalFile(signatureShrinkNmaePng)
		servers.DelLocalFile(sourcePdf)
		servers.DelLocalFile(mergedFile)
	}
	fmt.Printf("DownLoadLocalFile  file004 :%s err\n", cloudContractPath)
	//start upload and write db
	sc := userSalary.SalaryContract{
		Openid:                    openId,
		Status:                    int(STATUS_MERGED),
		Enter_contract_target_url: cloudContractPath,
	}
	//var un userSalary.SalaryContract
	_ = c.ShouldBindJSON(&sc)
	log.L.Info("UploadUserContract new json:", sc)
	err, _ = sc.UpdateSalaryContract()
	if err != nil {
		fmt.Printf("contract 更新失败：%v", err)
	} else {
		fmt.Printf("contract 更新成功")
	}
}

func getContractLastFile(c *gin.Context, openId string) string{
	sc := userSalary.SalaryContract{
		Openid: openId,
	}
	//var un userSalary.SalaryContract
	_ = c.ShouldBindJSON(&sc)
	err, reun := sc.FindById()
	//log.L.Info("UploadUserContract new json:", reun.Enter_contract_source_url)
	if len(reun.Enter_contract_source_url) <0{
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败：%v", err), gin.H{})
		return ""
	}
	resultJpg := strings.Split(reun.Enter_contract_source_url, ";")
	err,localFile := servers.DownLoadLocalFile(USER_HEADER_BUCKET,resultJpg[len(resultJpg) -1],"tmp")
	if err != nil {
		fmt.Printf("DownLoadLocalFile  file :%s err:%v\n", localFile,err)
		return ""
	} 
	fmt.Printf("getContractLastFile   down success :%s err\n", localFile)

	return localFile
}