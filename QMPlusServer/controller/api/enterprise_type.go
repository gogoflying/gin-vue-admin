package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

// @Tags EnterpriseType
// @Summary 创建EnterpriseType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseType true "创建EnterpriseType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /et/createEnterpriseType [post]
func CreateEnterpriseType(c *gin.Context) {
	var et userJobs.EnterpriseType
	_ = c.ShouldBindJSON(&et)
	err := et.CreateEnterpriseType()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags EnterpriseType
// @Summary 删除EnterpriseType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseType true "删除EnterpriseType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /et/deleteEnterpriseType [post]
func DeleteEnterpriseType(c *gin.Context) {
	var et userJobs.EnterpriseType
	_ = c.ShouldBindJSON(&et)
	err := et.DeleteEnterpriseType()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags EnterpriseType
// @Summary 更新EnterpriseType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseType true "更新EnterpriseType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /et/updateEnterpriseType [post]
func UpdateEnterpriseType(c *gin.Context) {
	var et userJobs.EnterpriseType
	_ = c.ShouldBindJSON(&et)
	err, reet := et.UpdateEnterpriseType()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reet": reet,
		})
	}
}

// @Tags EnterpriseType
// @Summary 用id查询EnterpriseType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseType true "用id查询EnterpriseType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /et/findEnterpriseType [post]
func FindEnterpriseType(c *gin.Context) {
	var et userJobs.EnterpriseType
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

// @Tags EnterpriseType
// @Summary 分页获取EnterpriseType列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取EnterpriseType列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /et/getEnterpriseTypeList [post]
func GetEnterpriseTypeList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.EnterpriseType).GetInfoList(pageInfo)
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
