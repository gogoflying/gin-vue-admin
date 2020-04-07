package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

// @Tags UserWork
// @Summary 创建UserWork
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.UserWork true "创建UserWork"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wk/createUserWork [post]
func CreateUserWork(c *gin.Context) {
	var wk userJobs.UserWork
	_ = c.ShouldBindJSON(&wk)
	err, wks := wk.CreateUserWork()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{
			"userJobsList": wks,
		})
	}
}

// @Tags UserWork
// @Summary 删除UserWork
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.UserWork true "删除UserWork"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wk/deleteUserWork [post]
func DeleteUserWork(c *gin.Context) {
	var wk userJobs.UserWork
	_ = c.ShouldBindJSON(&wk)
	err, wks := wk.DeleteUserWork()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{
			"userJobsList": wks,
		})
	}
}

// @Tags UserWork
// @Summary 更新UserWork
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.UserWork true "更新UserWork"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wk/updateUserWork [post]
func UpdateUserWork(c *gin.Context) {
	var wk userJobs.UserWork
	_ = c.ShouldBindJSON(&wk)
	err, rewk := wk.UpdateUserWork()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"rewk": rewk,
		})
	}
}

// @Tags UserWork
// @Summary 用id查询UserWork
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.UserWork true "用id查询UserWork"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wk/findUserWork [post]
func FindUserWork(c *gin.Context) {
	var wk userJobs.UserWork
	_ = c.ShouldBindJSON(&wk)
	err, rewk := wk.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"rewk": rewk,
		})
	}
}

// @Tags UserWork
// @Summary 分页获取UserWork列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取UserWork列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wk/getUserWorkList [post]
func GetUserWorkList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.UserWork).GetInfoList(pageInfo)
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

type ReqWorkList struct {
	OpenId   string `json:"openid"`
	Page     int
	PageSize int
}

func GetUserWorkListByOpenid(c *gin.Context) {
	var pageInfo ReqWorkList
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.UserWork).GetInfoListOpenId(pageInfo.OpenId, pageInfo.Page, pageInfo.PageSize)
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
