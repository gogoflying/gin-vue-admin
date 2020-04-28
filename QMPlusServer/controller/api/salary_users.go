package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userSalary"

	"github.com/gin-gonic/gin"
)

// @Tags SalaryUsers
// @Summary 创建SalaryUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryUsers true "创建SalaryUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/createSalaryUsers [post]
func CreateSalaryUsers(c *gin.Context) {
	var un userSalary.SalaryUsers
	_ = c.ShouldBindJSON(&un)
	err := un.CreateSalaryUsers()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SalaryUsers
// @Summary 删除SalaryUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryUsers true "删除SalaryUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /un/deleteSalaryUsers [post]
func DeleteSalaryUsers(c *gin.Context) {
	var un userSalary.SalaryUsers
	_ = c.ShouldBindJSON(&un)
	err := un.DeleteSalaryUsers()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SalaryUsers
// @Summary 更新SalaryUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryUsers true "更新SalaryUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /un/updateSalaryUsers [post]
func UpdateSalaryUsers(c *gin.Context) {
	var un userSalary.SalaryUsers
	_ = c.ShouldBindJSON(&un)
	err, reun := un.UpdateSalaryUsers()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reun": reun,
		})
	}
}

// @Tags SalaryUsers
// @Summary 用id查询SalaryUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryUsers true "用id查询SalaryUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /un/findSalaryUsers [post]
func FindSalaryUsers(c *gin.Context) {
	var un userSalary.SalaryUsers
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

func FindSalaryUsersByOpenid(c *gin.Context) {
	var un userSalary.SalaryUsers
	_ = c.ShouldBindJSON(&un)
	err, reun := un.FindByOpenid()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reun": reun,
		})
	}
}

// @Tags SalaryUsers
// @Summary 分页获取SalaryUsers列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取SalaryUsers列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/getSalaryUsersList [post]
func GetSalaryUsersList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userSalary.SalaryUsers).GetInfoList(pageInfo)
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
