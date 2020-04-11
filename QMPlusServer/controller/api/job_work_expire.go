package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

// @Tags JobWorkExpire
// @Summary 创建JobWorkExpire
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.JobWorkExpire true "创建JobWorkExpire"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /et/createJobWorkExpire [post]
func CreateJobWorkExpire(c *gin.Context) {
	var et userJobs.JobWorkExpire
	_ = c.ShouldBindJSON(&et)
	err := et.CreateJobWorkExpire()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags JobWorkExpire
// @Summary 删除JobWorkExpire
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.JobWorkExpire true "删除JobWorkExpire"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /et/deleteJobWorkExpire [post]
func DeleteJobWorkExpire(c *gin.Context) {
	var et userJobs.JobWorkExpire
	_ = c.ShouldBindJSON(&et)
	err := et.DeleteJobWorkExpire()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags JobWorkExpire
// @Summary 更新JobWorkExpire
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.JobWorkExpire true "更新JobWorkExpire"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /et/updateJobWorkExpire [post]
func UpdateJobWorkExpire(c *gin.Context) {
	var et userJobs.JobWorkExpire
	_ = c.ShouldBindJSON(&et)
	err, reet := et.UpdateJobWorkExpire()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reet": reet,
		})
	}
}

// @Tags JobWorkExpire
// @Summary 用id查询JobWorkExpire
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.JobWorkExpire true "用id查询JobWorkExpire"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /et/findJobWorkExpire [post]
func FindJobWorkExpire(c *gin.Context) {
	var et userJobs.JobWorkExpire
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

// @Tags JobWorkExpire
// @Summary 分页获取JobWorkExpire列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取JobWorkExpire列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /et/getJobWorkExpireList [post]
func GetJobWorkExpireList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.JobWorkExpire).GetInfoList(pageInfo)
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
