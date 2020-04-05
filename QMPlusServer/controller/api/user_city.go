package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userCity"

	"github.com/gin-gonic/gin"
)

// @Tags Cityname
// @Summary 创建Cityname
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userCity.Cityname true "创建Cityname"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cityname/createCityname [post]
func CreateCityname(c *gin.Context) {
	var cityname userCity.Cityname
	_ = c.ShouldBindJSON(&cityname)
	err := cityname.CreateCityname()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags Cityname
// @Summary 删除Cityname
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userCity.Cityname true "删除Cityname"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cityname/deleteCityname [post]
func DeleteCityname(c *gin.Context) {
	var cityname userCity.Cityname
	_ = c.ShouldBindJSON(&cityname)
	err := cityname.DeleteCityname()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags Cityname
// @Summary 更新Cityname
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userCity.Cityname true "更新Cityname"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cityname/updateCityname [post]
func UpdateCityname(c *gin.Context) {
	var cityname userCity.Cityname
	_ = c.ShouldBindJSON(&cityname)
	err, recityname := cityname.UpdateCityname()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"recityname": recityname,
		})
	}
}

// @Tags Cityname
// @Summary 用id查询Cityname
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userCity.Cityname true "用id查询Cityname"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cityname/findCityname [post]
func FindCityname(c *gin.Context) {
	var cityname userCity.Cityname
	_ = c.ShouldBindJSON(&cityname)
	err, recityname := cityname.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"recityname": recityname,
		})
	}
}

// @Tags Cityname
// @Summary 分页获取Cityname列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取Cityname列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cityname/getCitynameList [post]
func GetCitynameList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userCity.Cityname).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"userCityList": list,
			"total":        total,
			"page":         pageInfo.Page,
			"pageSize":     pageInfo.PageSize,
		})
	}
}
