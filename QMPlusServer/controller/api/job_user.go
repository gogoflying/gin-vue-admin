package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

// @Tags Users
// @Summary 创建Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.Users true "创建Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /users/createUsers [post]
func CreateUsers(c *gin.Context) {
	var users userJobs.Users
	_ = c.ShouldBindJSON(&users)
	err := users.CreateUsers()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags Users
// @Summary 删除Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.Users true "删除Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /users/deleteUsers [post]
func DeleteUsers(c *gin.Context) {
	var users userJobs.Users
	_ = c.ShouldBindJSON(&users)
	err := users.DeleteUsers()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags Users
// @Summary 更新Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.Users true "更新Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /users/updateUsers [post]
func UpdateUsers(c *gin.Context) {
	var users userJobs.Users
	_ = c.ShouldBindJSON(&users)
	err, reusers := users.UpdateUsers()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reusers": reusers,
		})
	}
}

// @Tags Users
// @Summary 用id查询Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.Users true "用id查询Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /users/findUsers [post]
func FindUsers(c *gin.Context) {
	var users userJobs.Users
	_ = c.ShouldBindJSON(&users)
	err, reusers := users.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reusers": reusers,
		})
	}
}

func FindUsersByOpenid(c *gin.Context) {
	var users userJobs.Users
	_ = c.ShouldBindJSON(&users)
	err, user, edus, works, base, dream := users.FindByOpenId()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"user":  user,
			"edus":  edus,
			"works": works,
			"base":  base,
			"dream": dream,
		})
	}
}

// @Tags Users
// @Summary 分页获取Users列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取Users列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /users/getUsersList [post]
func GetUsersList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.Users).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"userList": list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})
	}
}
