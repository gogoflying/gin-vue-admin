package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

// @Tags JobCollect
// @Summary 创建JobCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.JobCollect true "创建JobCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jc/createJobCollect [post]
func CreateJobCollect(c *gin.Context) {
	var jc userJobs.JobCollect
	_ = c.ShouldBindJSON(&jc)
	err := jc.CreateJobCollect()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags JobCollect
// @Summary 删除JobCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.JobCollect true "删除JobCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jc/deleteJobCollect [post]
func DeleteJobCollect(c *gin.Context) {
	var jc userJobs.JobCollect
	_ = c.ShouldBindJSON(&jc)
	err := jc.DeleteJobCollect()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags JobCollect
// @Summary 更新JobCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.JobCollect true "更新JobCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /jc/updateJobCollect [post]
func UpdateJobCollect(c *gin.Context) {
	var jc userJobs.JobCollect
	_ = c.ShouldBindJSON(&jc)
	err, rejc := jc.UpdateJobCollect()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"rejc": rejc,
		})
	}
}

// @Tags JobCollect
// @Summary 用id查询JobCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.JobCollect true "用id查询JobCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /jc/findJobCollect [post]
func FindJobCollect(c *gin.Context) {
	var jc userJobs.JobCollect
	_ = c.ShouldBindJSON(&jc)
	err, rejc := jc.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"rejc": rejc,
		})
	}
}

// @Tags JobCollect
// @Summary 分页获取JobCollect列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取JobCollect列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jc/getJobCollectList [post]
func GetJobCollectList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.JobCollect).GetInfoList(pageInfo)
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

type ReqJobCollectList struct {
	Page     int    `json:"Page"`
	PageSize int    `json:"PageSize"`
	OpenId   string `json:"openid"`
}

func GetJobCollectListByOpenid(c *gin.Context) {
	var req ReqJobCollectList
	_ = c.ShouldBindJSON(&req)
	err, list := new(userJobs.JobCollect).GetInfoListByOpenid(req.OpenId, req.Page, req.PageSize)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"userJobsList": list,
			"page":         req.Page,
			"pageSize":     req.PageSize,
		})
	}
}
