package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

// @Tags JobDutyTime
// @Summary 创建JobDutyTime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.JobDutyTime true "创建JobDutyTime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /et/createJobDutyTime [post]
func CreateJobDutyTime(c *gin.Context) {
	var et userJobs.JobDutyTime
	_ = c.ShouldBindJSON(&et)
	err := et.CreateJobDutyTime()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags JobDutyTime
// @Summary 删除JobDutyTime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.JobDutyTime true "删除JobDutyTime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /et/deleteJobDutyTime [post]
func DeleteJobDutyTime(c *gin.Context) {
	var et userJobs.JobDutyTime
	_ = c.ShouldBindJSON(&et)
	err := et.DeleteJobDutyTime()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags JobDutyTime
// @Summary 更新JobDutyTime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.JobDutyTime true "更新JobDutyTime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /et/updateJobDutyTime [post]
func UpdateJobDutyTime(c *gin.Context) {
	var et userJobs.JobDutyTime
	_ = c.ShouldBindJSON(&et)
	err, reet := et.UpdateJobDutyTime()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reet": reet,
		})
	}
}

// @Tags JobDutyTime
// @Summary 用id查询JobDutyTime
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.JobDutyTime true "用id查询JobDutyTime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /et/findJobDutyTime [post]
func FindJobDutyTime(c *gin.Context) {
	var et userJobs.JobDutyTime
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

// @Tags JobDutyTime
// @Summary 分页获取JobDutyTime列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取JobDutyTime列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /et/getJobDutyTimeList [post]
func GetJobDutyTimeList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.JobDutyTime).GetInfoList(pageInfo)
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
