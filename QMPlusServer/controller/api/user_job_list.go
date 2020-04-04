package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

// @Tags Joblist
// @Summary 创建Joblist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userjobs.Joblist true "创建Joblist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jl/createJoblist [post]
func CreateJoblist(c *gin.Context) {
	var jl userJobs.Joblist
	_ = c.ShouldBindJSON(&jl)
	err := jl.CreateJoblist()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags Joblist
// @Summary 删除Joblist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userjobs.Joblist true "删除Joblist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jl/deleteJoblist [post]
func DeleteJoblist(c *gin.Context) {
	var jl userJobs.Joblist
	_ = c.ShouldBindJSON(&jl)
	err := jl.DeleteJoblist()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags Joblist
// @Summary 更新Joblist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userjobs.Joblist true "更新Joblist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /jl/updateJoblist [post]
func UpdateJoblist(c *gin.Context) {
	var jl userJobs.Joblist
	_ = c.ShouldBindJSON(&jl)
	err, rejl := jl.UpdateJoblist()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"rejl": rejl,
		})
	}
}

// @Tags Joblist
// @Summary 用id查询Joblist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userjobs.Joblist true "用id查询Joblist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /jl/findJoblist [post]
func FindJoblist(c *gin.Context) {
	var jl userJobs.Joblist
	_ = c.ShouldBindJSON(&jl)
	err, rejl := jl.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"rejl": rejl,
		})
	}
}

// @Tags Joblist
// @Summary 分页获取Joblist列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取Joblist列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jl/getJoblistList [post]
type ReqJoblist struct {
	modelInterface.PageInfo
	KeyWord     string `json:"keyword"`
	JobsalaHigh int    `json:"job_salary_high"`
	JobsalaLow  int    `json:"job_salary_low"`
}

func GetJoblistList(c *gin.Context) {
	var req ReqJoblist
	_ = c.ShouldBindJSON(&req)
	err, list, total := new(userJobs.Joblist).GetInfoListSearch(req.KeyWord, req.JobsalaLow, req.JobsalaHigh, req.PageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"RspJoblistList": list,
			"total":          total,
			"page":           req.PageInfo.Page,
			"pageSize":       req.PageInfo.PageSize,
		})
	}
}
