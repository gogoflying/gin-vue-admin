package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

// @Tags EduProfession
// @Summary 创建EduProfession
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EduProfession true "创建EduProfession"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /et/createEduProfession [post]
func CreateEduProfession(c *gin.Context) {
	var et userJobs.EduProfession
	_ = c.ShouldBindJSON(&et)
	err := et.CreateEduProfession()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags EduProfession
// @Summary 删除EduProfession
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EduProfession true "删除EduProfession"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /et/deleteEduProfession [post]
func DeleteEduProfession(c *gin.Context) {
	var et userJobs.EduProfession
	_ = c.ShouldBindJSON(&et)
	err := et.DeleteEduProfession()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags EduProfession
// @Summary 更新EduProfession
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EduProfession true "更新EduProfession"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /et/updateEduProfession [post]
func UpdateEduProfession(c *gin.Context) {
	var et userJobs.EduProfession
	_ = c.ShouldBindJSON(&et)
	err, reet := et.UpdateEduProfession()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reet": reet,
		})
	}
}

// @Tags EduProfession
// @Summary 用id查询EduProfession
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EduProfession true "用id查询EduProfession"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /et/findEduProfession [post]
func FindEduProfession(c *gin.Context) {
	var et userJobs.EduProfession
	_ = c.ShouldBindJSON(&et)
	err, reet := et.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reet": reet,
		})
	}
}

// @Tags EduProfession
// @Summary 分页获取EduProfession列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取EduProfession列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /et/getEduProfessionList [post]
func GetEduProfessionList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.EduProfession).GetInfoList(pageInfo)
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
