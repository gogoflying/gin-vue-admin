package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userSalary"

	"github.com/gin-gonic/gin"
)

// @Tags SbConf
// @Summary 创建SbConf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SbConf true "创建SbConf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/createSbConf [post]
func CreateSbConf(c *gin.Context) {
	var un userSalary.SbConf
	_ = c.ShouldBindJSON(&un)
	err := un.CreateSbConf()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SbConf
// @Summary 删除SbConf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SbConf true "删除SbConf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /un/deleteSbConf [post]
func DeleteSbConf(c *gin.Context) {
	var un userSalary.SbConf
	_ = c.ShouldBindJSON(&un)
	err := un.DeleteSbConf()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SbConf
// @Summary 更新SbConf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SbConf true "更新SbConf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /un/updateSbConf [post]
func UpdateSbConf(c *gin.Context) {
	var un userSalary.SbConf
	_ = c.ShouldBindJSON(&un)
	err, reun := un.UpdateSbConf()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reun": reun,
		})
	}
}

// @Tags SbConf
// @Summary 用id查询SbConf
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SbConf true "用id查询SbConf"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /un/findSbConf [post]
func FindSbConf(c *gin.Context) {
	var un userSalary.SbConf
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

// @Tags SbConf
// @Summary 分页获取SbConf列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取SbConf列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/getSbConfList [post]
func GetSbConfList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userSalary.SbConf).GetInfoList(pageInfo)
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
