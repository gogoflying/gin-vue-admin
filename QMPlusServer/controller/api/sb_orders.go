package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userSalary"

	"github.com/gin-gonic/gin"
)

// @Tags SbOrders
// @Summary 创建SbOrders
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SbOrders true "创建SbOrders"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/createSbOrders [post]
func CreateSbOrders(c *gin.Context) {
	var un userSalary.SbOrders
	_ = c.ShouldBindJSON(&un)
	err := un.CreateSbOrders()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SbOrders
// @Summary 删除SbOrders
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SbOrders true "删除SbOrders"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /un/deleteSbOrders [post]
func DeleteSbOrders(c *gin.Context) {
	var un userSalary.SbOrders
	_ = c.ShouldBindJSON(&un)
	err := un.DeleteSbOrders()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SbOrders
// @Summary 更新SbOrders
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SbOrders true "更新SbOrders"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /un/updateSbOrders [post]
func UpdateSbOrders(c *gin.Context) {
	var un userSalary.SbOrders
	_ = c.ShouldBindJSON(&un)
	err, reun := un.UpdateSbOrders()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reun": reun,
		})
	}
}

// @Tags SbOrders
// @Summary 用id查询SbOrders
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SbOrders true "用id查询SbOrders"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /un/findSbOrders [post]
func FindSbOrders(c *gin.Context) {
	var un userSalary.SbOrders
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

// @Tags SbOrders
// @Summary 分页获取SbOrders列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取SbOrders列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/getSbOrdersList [post]
func GetSbOrdersList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userSalary.SbOrders).GetInfoList(pageInfo)
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
