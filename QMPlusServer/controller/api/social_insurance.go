package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/socialInsurance"

	"github.com/gin-gonic/gin"
)

// @Tags SocialInsurance
// @Summary 创建SocialInsurance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialInsurance true "创建SocialInsurance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /si/createSocialInsurance [post]
func CreateSocialInsurance(c *gin.Context) {
	var si socialInsurance.SocialInsurance
	_ = c.ShouldBindJSON(&si)
	err := si.CreateSocialInsurance()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SocialInsurance
// @Summary 删除SocialInsurance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialInsurance true "删除SocialInsurance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /si/deleteSocialInsurance [post]
func DeleteSocialInsurance(c *gin.Context) {
	var si socialInsurance.SocialInsurance
	_ = c.ShouldBindJSON(&si)
	err := si.DeleteSocialInsurance()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SocialInsurance
// @Summary 更新SocialInsurance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialInsurance true "更新SocialInsurance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /si/updateSocialInsurance [post]
func UpdateSocialInsurance(c *gin.Context) {
	var si socialInsurance.SocialInsurance
	_ = c.ShouldBindJSON(&si)
	err, reet := si.UpdateSocialInsurance()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reet": reet,
		})
	}
}

// @Tags SocialInsurance
// @Summary 用id查询SocialInsurance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialInsurance true "用id查询SocialInsurance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /si/findSocialInsurance [post]
func FindSocialInsurance(c *gin.Context) {
	var si socialInsurance.SocialInsurance
	_ = c.ShouldBindJSON(&si)
	err, reet := si.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reet": reet,
		})
	}
}

// @Tags SocialInsurance
// @Summary 分页获取SocialInsurance列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取SocialInsurance列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /si/getSocialInsuranceList [post]
func GetSocialInsuranceList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(socialInsurance.SocialInsurance).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"socialInsuranceList": list,
			"total":               total,
			"page":                pageInfo.Page,
			"pageSize":            pageInfo.PageSize,
		})
	}
}
