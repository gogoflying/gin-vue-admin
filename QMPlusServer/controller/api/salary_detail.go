package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userSalary"

	"github.com/gin-gonic/gin"
)

// @Tags Salarys
// @Summary 创建Salarys
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.Salarys true "创建Salarys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/createSalarys [post]
func CreateSalarys(c *gin.Context) {
	var un userSalary.Salarys
	_ = c.ShouldBindJSON(&un)
	err := un.CreateSalarys()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags Salarys
// @Summary 删除Salarys
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.Salarys true "删除Salarys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /un/deleteSalarys [post]
func DeleteSalarys(c *gin.Context) {
	var un userSalary.Salarys
	_ = c.ShouldBindJSON(&un)
	err := un.DeleteSalarys()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags Salarys
// @Summary 更新Salarys
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.Salarys true "更新Salarys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /un/updateSalarys [post]
func UpdateSalarys(c *gin.Context) {
	var un userSalary.Salarys
	_ = c.ShouldBindJSON(&un)
	err, reun := un.UpdateSalarys()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reun": reun,
		})
	}
}

// @Tags Salarys
// @Summary 用id查询Salarys
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.Salarys true "用id查询Salarys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /un/findSalarys [post]
func FindSalarys(c *gin.Context) {
	var un userSalary.Salarys
	_ = c.ShouldBindJSON(&un)
	err, reun := un.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reun": reun,
		})
	}
}

func FindSalarysByIdAndOpenid(c *gin.Context) {
	var un userSalary.Salarys
	_ = c.ShouldBindJSON(&un)
	err, reun := un.FindByIdAndOpenid()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reun": reun,
		})
	}
}

// @Tags Salarys
// @Summary 分页获取Salarys列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取Salarys列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/getSalarysList [post]
func GetSalarysList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userSalary.Salarys).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"userSalaryList": list,
			"total":          total,
			"page":           pageInfo.Page,
			"pageSize":       pageInfo.PageSize,
		})
	}
}

type ReqSalarysListSearch struct {
	Openid   string `json:"openid"`
	Page     int
	PageSize int
	YearL    int `json:"year_low"`
	YearH    int `json:"year_high"`
	MonthL   int `json:"month_low"`
	MonthH   int `json:"month_high"`
}

func GetSalarysListSearch(c *gin.Context) {
	var req ReqSalarysListSearch
	_ = c.ShouldBindJSON(&req)
	err, list := new(userSalary.Salarys).GetInfoListSearch(req.Openid, req.YearL, req.MonthL, req.YearH, req.MonthH, req.Page, req.PageSize)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"userSalaryList": list,
			"page":           req.Page,
			"pageSize":       req.PageSize,
		})
	}
}
