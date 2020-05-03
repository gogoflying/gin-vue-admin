package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userSalary"

	"github.com/gin-gonic/gin"
)

// @Tags SbUsers
// @Summary 创建SbUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SbUsers true "创建SbUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/createSbUsers [post]
func CreateSbUsers(c *gin.Context) {
	var un userSalary.SbUsers
	_ = c.ShouldBindJSON(&un)
	err := un.CreateSbUsers()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SbUsers
// @Summary 删除SbUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SbUsers true "删除SbUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /un/deleteSbUsers [post]
func DeleteSbUsers(c *gin.Context) {
	var un userSalary.SbUsers
	_ = c.ShouldBindJSON(&un)
	err := un.DeleteSbUsers()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SbUsers
// @Summary 更新SbUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SbUsers true "更新SbUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /un/updateSbUsers [post]
func UpdateSbUsers(c *gin.Context) {
	var un userSalary.SbUsers
	_ = c.ShouldBindJSON(&un)
	err, reun := un.UpdateSbUsers()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reun": reun,
		})
	}
}

// @Tags SbUsers
// @Summary 用id查询SbUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SbUsers true "用id查询SbUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /un/findSbUsers [post]
func FindSbUsers(c *gin.Context) {
	var un userSalary.SbUsers
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

func FindSbUsersByOpenid(c *gin.Context) {
	var un userSalary.SbUsers
	_ = c.ShouldBindJSON(&un)
	err, reun := un.FindByOpenid()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		if reun.ID == 0 {
			var us userSalary.SalaryUsers
			us.Openid = un.Openid
			_, us = us.FindByOpenid()
			reun.Name = us.Name
			reun.Card = us.Card
			reun.Mobile = us.Mobile
		}
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reun": reun,
		})
	}
}

// @Tags SbUsers
// @Summary 分页获取SbUsers列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取SbUsers列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/getSbUsersList [post]
func GetSbUsersList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userSalary.SbUsers).GetInfoList(pageInfo)
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
