package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

// @Tags UserBaseInfo
// @Summary 创建UserBaseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.UserBaseInfo true "创建UserBaseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bf/createUserBaseInfo [post]
func CreateUserBaseInfo(c *gin.Context) {
	var bf userJobs.UserBaseInfo
	_ = c.ShouldBindJSON(&bf)
	err := bf.CreateUserBaseInfo()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags UserBaseInfo
// @Summary 删除UserBaseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.UserBaseInfo true "删除UserBaseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bf/deleteUserBaseInfo [post]
func DeleteUserBaseInfo(c *gin.Context) {
	var bf userJobs.UserBaseInfo
	_ = c.ShouldBindJSON(&bf)
	err := bf.DeleteUserBaseInfo()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags UserBaseInfo
// @Summary 更新UserBaseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.UserBaseInfo true "更新UserBaseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bf/updateUserBaseInfo [post]
func UpdateUserBaseInfo(c *gin.Context) {
	var bf userJobs.UserBaseInfo
	_ = c.ShouldBindJSON(&bf)
	err, rebf := bf.UpdateUserBaseInfo()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"rebf": rebf,
		})
	}
}

func UpdateUserAvatarUrl(c *gin.Context) {
	var bf userJobs.UserBaseInfo
	_ = c.ShouldBindJSON(&bf)
	err := bf.UpdateUserAvatarUrl()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{})
	}
}

// @Tags UserBaseInfo
// @Summary 用id查询UserBaseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.UserBaseInfo true "用id查询UserBaseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bf/findUserBaseInfo [post]
func FindUserBaseInfo(c *gin.Context) {
	var bf userJobs.UserBaseInfo
	_ = c.ShouldBindJSON(&bf)
	err, rebf := bf.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"rebf": rebf,
		})
	}
}

func FindUserBaseInfoByOpenid(c *gin.Context) {
	var bf userJobs.UserBaseInfo
	_ = c.ShouldBindJSON(&bf)
	err, rebf := bf.FindByOpenid()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"rebf": rebf,
		})
	}
}

// @Tags UserBaseInfo
// @Summary 分页获取UserBaseInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取UserBaseInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bf/getUserBaseInfoList [post]
func GetUserBaseInfoList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.UserBaseInfo).GetInfoList(pageInfo)
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
