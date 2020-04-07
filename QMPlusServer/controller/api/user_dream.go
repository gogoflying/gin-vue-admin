package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

// @Tags UserDream
// @Summary 创建UserDream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.UserDream true "创建UserDream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dm/createUserDream [post]
func CreateUserDream(c *gin.Context) {
	var dm userJobs.UserDream
	_ = c.ShouldBindJSON(&dm)
	err := dm.CreateUserDream()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags UserDream
// @Summary 删除UserDream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.UserDream true "删除UserDream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dm/deleteUserDream [post]
func DeleteUserDream(c *gin.Context) {
	var dm userJobs.UserDream
	_ = c.ShouldBindJSON(&dm)
	err := dm.DeleteUserDream()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags UserDream
// @Summary 更新UserDream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.UserDream true "更新UserDream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dm/updateUserDream [post]
func UpdateUserDream(c *gin.Context) {
	var dm userJobs.UserDream
	_ = c.ShouldBindJSON(&dm)
	err, redm := dm.UpdateUserDream()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"redm": redm,
		})
	}
}

// @Tags UserDream
// @Summary 用id查询UserDream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.UserDream true "用id查询UserDream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dm/findUserDream [post]
func FindUserDream(c *gin.Context) {
	var dm userJobs.UserDream
	_ = c.ShouldBindJSON(&dm)
	err, redm := dm.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"redm": redm,
		})
	}
}

func FindUserDreamByOpenid(c *gin.Context) {
	var dm userJobs.UserDream
	_ = c.ShouldBindJSON(&dm)
	err, redm := dm.FindByOpenid()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"redm": redm,
		})
	}
}

// @Tags UserDream
// @Summary 分页获取UserDream列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取UserDream列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dm/getUserDreamList [post]
func GetUserDreamList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.UserDream).GetInfoList(pageInfo)
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
