package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

// @Tags JobSalary
// @Summary 创建JobSalary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.JobSalary true "创建JobSalary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /et/createJobSalary [post]
func CreateJobSalary(c *gin.Context) {
	var et userJobs.JobSalary
	_ = c.ShouldBindJSON(&et)
	err := et.CreateJobSalary()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags JobSalary
// @Summary 删除JobSalary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.JobSalary true "删除JobSalary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /et/deleteJobSalary [post]
func DeleteJobSalary(c *gin.Context) {
	var et userJobs.JobSalary
	_ = c.ShouldBindJSON(&et)
	err := et.DeleteJobSalary()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags JobSalary
// @Summary 更新JobSalary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.JobSalary true "更新JobSalary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /et/updateJobSalary [post]
func UpdateJobSalary(c *gin.Context) {
	var et userJobs.JobSalary
	_ = c.ShouldBindJSON(&et)
	err, reet := et.UpdateJobSalary()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reet": reet,
		})
	}
}

// @Tags JobSalary
// @Summary 用id查询JobSalary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.JobSalary true "用id查询JobSalary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /et/findJobSalary [post]
func FindJobSalary(c *gin.Context) {
	var et userJobs.JobSalary
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

// @Tags JobSalary
// @Summary 分页获取JobSalary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取JobSalary列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /et/getJobSalaryList [post]
func GetJobSalaryList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.JobSalary).GetInfoList(pageInfo)
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
