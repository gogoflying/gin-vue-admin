package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

// @Tags EnterpriseCollect
// @Summary 创建EnterpriseCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseCollect true "创建EnterpriseCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jc/createEnterpriseCollect [post]
func CreateEnterpriseCollect(c *gin.Context) {
	var jc userJobs.EnterpriseCollect
	_ = c.ShouldBindJSON(&jc)
	err := jc.CreateEnterpriseCollect()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags EnterpriseCollect
// @Summary 删除EnterpriseCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseCollect true "删除EnterpriseCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jc/deleteEnterpriseCollect [post]
func DeleteEnterpriseCollect(c *gin.Context) {
	var jc userJobs.EnterpriseCollect
	_ = c.ShouldBindJSON(&jc)
	err := jc.DeleteEnterpriseCollect()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags EnterpriseCollect
// @Summary 更新EnterpriseCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseCollect true "更新EnterpriseCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /jc/updateEnterpriseCollect [post]
func UpdateEnterpriseCollect(c *gin.Context) {
	var jc userJobs.EnterpriseCollect
	_ = c.ShouldBindJSON(&jc)
	err, rejc := jc.UpdateEnterpriseCollect()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"rejc": rejc,
		})
	}
}

// @Tags EnterpriseCollect
// @Summary 用id查询EnterpriseCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseCollect true "用id查询EnterpriseCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /jc/findEnterpriseCollect [post]
func FindEnterpriseCollect(c *gin.Context) {
	var jc userJobs.EnterpriseCollect
	_ = c.ShouldBindJSON(&jc)
	err, rejc := jc.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"rejc": rejc,
		})
	}
}

// @Tags EnterpriseCollect
// @Summary 分页获取EnterpriseCollect列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取EnterpriseCollect列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jc/getEnterpriseCollectList [post]
func GetEnterpriseCollectList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.EnterpriseCollect).GetInfoList(pageInfo)
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

type ReqEnterpriseCollectList struct {
	Page     int    `json:"Page"`
	PageSize int    `json:"PageSize"`
	OpenId   string `json:"openid"`
}

func GetEnterpriseCollectListByOpenid(c *gin.Context) {
	var req ReqEnterpriseCollectList
	_ = c.ShouldBindJSON(&req)
	err, list := new(userJobs.EnterpriseCollect).GetInfoListByOpenid(req.OpenId, req.Page, req.PageSize)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"userJobsList": list,
			"page":         req.Page,
			"pageSize":     req.PageSize,
		})
	}
}
