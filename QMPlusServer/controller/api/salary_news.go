package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userSalary"

	"github.com/gin-gonic/gin"
)

// @Tags UserNews
// @Summary 创建UserNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.UserNews true "创建UserNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/createUserNews [post]
func CreateUserNews(c *gin.Context) {
	var un userSalary.UserNews
	_ = c.ShouldBindJSON(&un)
	err := un.CreateUserNews()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags UserNews
// @Summary 删除UserNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.UserNews true "删除UserNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /un/deleteUserNews [post]
func DeleteUserNews(c *gin.Context) {
	var un userSalary.UserNews
	_ = c.ShouldBindJSON(&un)
	err := un.DeleteUserNews()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags UserNews
// @Summary 更新UserNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.UserNews true "更新UserNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /un/updateUserNews [post]
func UpdateUserNews(c *gin.Context) {
	var un userSalary.UserNews
	_ = c.ShouldBindJSON(&un)
	err, reun := un.UpdateUserNews()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reun": reun,
		})
	}
}

func UpdateUserNewsCount(c *gin.Context) {
	var un userSalary.UserNews
	_ = c.ShouldBindJSON(&un)
	err := un.UpdateUserNewsCount()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{})
	}
}

// @Tags UserNews
// @Summary 用id查询UserNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.UserNews true "用id查询UserNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /un/findUserNews [post]
func FindUserNews(c *gin.Context) {
	var un userSalary.UserNews
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

// @Tags UserNews
// @Summary 分页获取UserNews列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取UserNews列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/getUserNewsList [post]
func GetUserNewsList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userSalary.UserNews).GetInfoList(pageInfo)
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

type ReqUserNewsListByType struct {
	Page     int
	PageSize int
	Type     int `json:"news_type"`
}

func GetUserNewsListByType(c *gin.Context) {
	var req ReqUserNewsListByType
	_ = c.ShouldBindJSON(&req)
	err, list := new(userSalary.UserNews).GetInfoListByTyp(req.Type, req.Page, req.PageSize)
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
