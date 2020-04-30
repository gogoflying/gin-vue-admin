package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

// @Tags EnterpriseInfo
// @Summary 创建EnterpriseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseInfo true "创建EnterpriseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /info/createEnterpriseInfo [post]
func CreateEnterpriseInfo(c *gin.Context) {
	var info userJobs.EnterpriseInfo
	_ = c.ShouldBindJSON(&info)
	err := info.CreateEnterpriseInfo()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags EnterpriseInfo
// @Summary 删除EnterpriseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseInfo true "删除EnterpriseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /info/deleteEnterpriseInfo [post]
func DeleteEnterpriseInfo(c *gin.Context) {
	var info userJobs.EnterpriseInfo
	_ = c.ShouldBindJSON(&info)
	err := info.DeleteEnterpriseInfo()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags EnterpriseInfo
// @Summary 更新EnterpriseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseInfo true "更新EnterpriseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /info/updateEnterpriseInfo [post]
func UpdateEnterpriseInfo(c *gin.Context) {
	var info userJobs.EnterpriseInfo
	_ = c.ShouldBindJSON(&info)
	err, reinfo := info.UpdateEnterpriseInfo()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reinfo": reinfo,
		})
	}
}

// @Tags EnterpriseInfo
// @Summary 用id查询EnterpriseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseInfo true "用id查询EnterpriseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /info/findEnterpriseInfo [post]
func FindEnterpriseInfo(c *gin.Context) {
	var info userJobs.EnterpriseInfo
	_ = c.ShouldBindJSON(&info)
	err, reinfo := info.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reinfo": reinfo,
		})
	}
}

type ReqEnterpriseJobs struct {
	ID       uint
	Page     int
	PageSize int
}

func FindEnterpriseJobs(c *gin.Context) {
	var req ReqEnterpriseJobs
	_ = c.ShouldBindJSON(&req)
	var info userJobs.EnterpriseInfo
	info.ID = req.ID
	err, reinfo := info.FindPositionsByCompId(req.Page, req.PageSize)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reinfo": reinfo,
		})
	}
}

// @Tags EnterpriseInfo
// @Summary 分页获取EnterpriseInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取EnterpriseInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /info/getEnterpriseInfoList [post]
func GetEnterpriseInfoList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.EnterpriseInfo).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"result":   list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})
	}
}

type ReqEnterpriseSearch struct {
	Page     int
	PageSize int
	CityId   int    `json:"city_id"`
	KeyWord  string `json:"keyword"`
}

func GetEnterpriseInfoListSearch(c *gin.Context) {
	var req ReqEnterpriseSearch
	_ = c.ShouldBindJSON(&req)
	err, list, total := new(userJobs.EnterpriseInfo).GetInfoListSearch(req.KeyWord, req.CityId, req.Page, req.PageSize)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"enterprise_list": list,
			"total":           total,
			"page":            req.Page,
			"pageSize":        req.PageSize,
		})
	}
}

func GetEnterpriseAllInfo(c *gin.Context) {
	err, list := new(userJobs.EnterpriseInfo).GetAllInfoList()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"result": list,
		})
	}
}
