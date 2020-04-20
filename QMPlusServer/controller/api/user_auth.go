package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

// @Tags UserAuth
// @Summary 创建UserAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAuth true "创建UserAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userauth/createUserAuth [post]
func CreateUserAuth(c *gin.Context) {
	var userauth userJobs.UserAuth
	_ = c.ShouldBindJSON(&userauth)
	err := userauth.CreateUserAuth()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags UserAuth
// @Summary 删除UserAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAuth true "删除UserAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userauth/deleteUserAuth [post]
func DeleteUserAuth(c *gin.Context) {
	var userauth userJobs.UserAuth
	_ = c.ShouldBindJSON(&userauth)
	err := userauth.DeleteUserAuth()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags UserAuth
// @Summary 更新UserAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAuth true "更新UserAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userauth/updateUserAuth [post]
func UpdateUserAuth(c *gin.Context) {
	var userauth userJobs.UserAuth
	_ = c.ShouldBindJSON(&userauth)
	err, reuserauth := userauth.UpdateUserAuth()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reua": reuserauth,
		})
	}
}

// @Tags UserAuth
// @Summary 用id查询UserAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAuth true "用id查询UserAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userauth/findUserAuth [post]
func FindUserAuth(c *gin.Context) {
	var userauth userJobs.UserAuth
	_ = c.ShouldBindJSON(&userauth)
	err, reuserauth := userauth.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reua": reuserauth,
		})
	}
}

// @Tags UserAuth
// @Summary 分页获取UserAuth列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取UserAuth列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userauth/getUserAuthList [post]
func GetUserAuthList(c *gin.Context) {
	// 此结构体仅本方法使用
	type searchParams struct {
		userJobs.UserAuth
		modelInterface.PageInfo
	}
	var sp searchParams
	_ = c.ShouldBindJSON(&sp)
	err, list, total := sp.UserAuth.GetInfoList(sp.PageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"userAuthsList": list,
			"total":         total,
			"page":          sp.PageInfo.Page,
			"pageSize":      sp.PageInfo.PageSize,
		})
	}
}
