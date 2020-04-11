package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

// @Tags EnterpriseIndust
// @Summary 创建EnterpriseIndust
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseIndust true "创建EnterpriseIndust"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /et/createEnterpriseIndust [post]
func CreateEnterpriseIndust(c *gin.Context) {
	var et userJobs.EnterpriseIndust
	_ = c.ShouldBindJSON(&et)
	err := et.CreateEnterpriseIndust()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags EnterpriseIndust
// @Summary 删除EnterpriseIndust
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseIndust true "删除EnterpriseIndust"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /et/deleteEnterpriseIndust [post]
func DeleteEnterpriseIndust(c *gin.Context) {
	var et userJobs.EnterpriseIndust
	_ = c.ShouldBindJSON(&et)
	err := et.DeleteEnterpriseIndust()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags EnterpriseIndust
// @Summary 更新EnterpriseIndust
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseIndust true "更新EnterpriseIndust"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /et/updateEnterpriseIndust [post]
func UpdateEnterpriseIndust(c *gin.Context) {
	var et userJobs.EnterpriseIndust
	_ = c.ShouldBindJSON(&et)
	err, reet := et.UpdateEnterpriseIndust()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reet": reet,
		})
	}
}

// @Tags EnterpriseIndust
// @Summary 用id查询EnterpriseIndust
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseIndust true "用id查询EnterpriseIndust"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /et/findEnterpriseIndust [post]
func FindEnterpriseIndust(c *gin.Context) {
	var et userJobs.EnterpriseIndust
	_ = c.ShouldBindJSON(&et)
	err, reet := et.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reet": reet,
		})
	}
}

// @Tags EnterpriseIndust
// @Summary 分页获取EnterpriseIndust列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取EnterpriseIndust列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /et/getEnterpriseIndustList [post]
func GetEnterpriseIndustList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.EnterpriseIndust).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"userJobsList": list,
			"total":        total,
			"page":         pageInfo.Page,
			"pageSize":     pageInfo.PageSize,
		})
	}
}
