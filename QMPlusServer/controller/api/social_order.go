package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/socialInsurance"

	"github.com/gin-gonic/gin"
)

// @Tags SocialOrder
// @Summary 创建SocialOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialOrder true "创建SocialOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /so/createSocialOrder [post]
func CreateSocialOrder(c *gin.Context) {
	var so socialInsurance.SocialOrder
	_ = c.ShouldBindJSON(&so)
	err := so.CreateSocialOrder()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SocialOrder
// @Summary 删除SocialOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialOrder true "删除SocialOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /so/deleteSocialOrder [post]
func DeleteSocialOrder(c *gin.Context) {
	var so socialInsurance.SocialOrder
	_ = c.ShouldBindJSON(&so)
	err := so.DeleteSocialOrder()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SocialOrder
// @Summary 更新SocialOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialOrder true "更新SocialOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /so/updateSocialOrder [post]
func UpdateSocialOrder(c *gin.Context) {
	var so socialInsurance.SocialOrder
	_ = c.ShouldBindJSON(&so)
	err, reet := so.UpdateSocialOrder()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reet": reet,
		})
	}
}

// @Tags SocialOrder
// @Summary 用id查询SocialOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialOrder true "用id查询SocialOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /so/findSocialOrder [post]
func FindSocialOrder(c *gin.Context) {
	var so socialInsurance.SocialOrder
	_ = c.ShouldBindJSON(&so)
	err, reet := so.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reet": reet,
		})
	}
}

// @Tags SocialOrder
// @Summary 分页获取SocialOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取SocialOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /so/getSocialOrderList [post]
func GetSocialOrderList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(socialInsurance.SocialOrder).GetInfoList(pageInfo)
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
