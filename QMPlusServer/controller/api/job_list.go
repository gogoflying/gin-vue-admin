package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Tags Joblist
// @Summary 创建Joblist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.Joblist true "创建Joblist"
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
// @Param data body userJobs.Joblist true "删除Joblist"
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
// @Param data body userJobs.Joblist true "更新Joblist"
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
// @Param data body userJobs.Joblist true "用id查询Joblist"
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
func GetJoblistList(c *gin.Context) {
	var (
		req   ReqJoblist
		low   int64
		hight int64
		err   error
		list  interface{}
		total int
	)
	_ = c.ShouldBindJSON(&req)
	if strings.Contains(req.KeyWord, "-") {
		array := strings.Split(req.KeyWord, "-")
		if len(array) == 2 {
			low, _ = strconv.ParseInt(array[0], 10, 32)
			hight, _ = strconv.ParseInt(array[1], 10, 32)
		}
		err, list, total = new(userJobs.Joblist).GetInfoListSearch("", req.CityId, int(low), int(hight), modelInterface.PageInfo{Page: req.Page, PageSize: req.PageSize})
	} else {
		err, list, total = new(userJobs.Joblist).GetInfoListSearch(req.KeyWord, req.CityId, 0, 0, modelInterface.PageInfo{Page: req.Page, PageSize: req.PageSize})
	}
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"jobs":     list,
			"total":    total,
			"page":     req.PageInfo.Page,
			"pageSize": req.PageInfo.PageSize,
		})
	}
}

type ReqJoblist struct {
	modelInterface.PageInfo
	CityId  int    `json:"city_id"`
	KeyWord string `json:"keyword"`
}
