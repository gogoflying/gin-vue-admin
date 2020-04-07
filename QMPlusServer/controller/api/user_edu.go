package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

// @Tags Usereducation
// @Summary 创建Usereducation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.Usereducation true "创建Usereducation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ed/createUsereducation [post]
func CreateUsereducation(c *gin.Context) {
	var ed userJobs.Usereducation
	_ = c.ShouldBindJSON(&ed)
	err, reed := ed.CreateUsereducation()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{
			"userJobsList": reed,
		})
	}
}

// @Tags Usereducation
// @Summary 删除Usereducation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.Usereducation true "删除Usereducation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ed/deleteUsereducation [post]
func DeleteUsereducation(c *gin.Context) {
	var ed userJobs.Usereducation
	_ = c.ShouldBindJSON(&ed)
	err := ed.DeleteUsereducation()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags Usereducation
// @Summary 更新Usereducation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.Usereducation true "更新Usereducation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ed/updateUsereducation [post]
func UpdateUsereducation(c *gin.Context) {
	var ed userJobs.Usereducation
	_ = c.ShouldBindJSON(&ed)
	err, reed := ed.UpdateUsereducation()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"userJobsList": reed,
		})
	}
}

// @Tags Usereducation
// @Summary 用id查询Usereducation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.Usereducation true "用id查询Usereducation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ed/findUsereducation [post]
func FindUsereducation(c *gin.Context) {
	var ed userJobs.Usereducation
	_ = c.ShouldBindJSON(&ed)
	err, reed := ed.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reed": reed,
		})
	}
}

// @Tags Usereducation
// @Summary 分页获取Usereducation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取Usereducation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ed/getUsereducationList [post]
func GetUsereducationList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.Usereducation).GetInfoList(pageInfo)
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

type ReqEduList struct {
	OpenId   string `json:"openid"`
	Page     int
	PageSize int
}

func GetUsereducationListByOpenid(c *gin.Context) {
	var req ReqEduList
	_ = c.ShouldBindJSON(&req)
	err, list, total := new(userJobs.Usereducation).GetInfoListByOpenid(req.OpenId, req.Page, req.PageSize)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"userJobsList": list,
			"total":        total,
			"page":         req.Page,
			"pageSize":     req.PageSize,
		})
	}
}
