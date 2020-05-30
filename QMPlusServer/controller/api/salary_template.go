package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"
	"gin-vue-admin/model/userSalary"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Tags SalaryTemplates
// @Summary 创建SalaryTemplates
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryTemplates true "创建SalaryTemplates"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/createSalaryTemplates [post]
func CreateSalaryTemplates(c *gin.Context) {
	var st userSalary.SalaryTemplates
	_ = c.ShouldBindJSON(&st)
	ei, exist := c.Get("enpInfo")
	if exist {
		enpInfo := ei.(*userJobs.EnterpriseInfo)
		st.Enterprise = enpInfo.EnterPriseName
		st.EnterpriseId = int(enpInfo.ID)
	}
	err := st.CreateSalaryTemplates()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SalaryTemplates
// @Summary 删除SalaryTemplates
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryTemplates true "删除SalaryTemplates"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /un/deleteSalaryTemplates [post]
func DeleteSalaryTemplates(c *gin.Context) {
	var st userSalary.SalaryTemplates
	_ = c.ShouldBindJSON(&st)
	err := st.DeleteSalaryTemplates()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SalaryTemplates
// @Summary 更新SalaryTemplates
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryTemplates true "更新SalaryTemplates"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /un/updateSalaryTemplates [post]
func UpdateSalaryTemplates(c *gin.Context) {
	var st userSalary.SalaryTemplates
	_ = c.ShouldBindJSON(&st)
	err, reet := st.UpdateSalaryTemplates()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reet": reet,
		})
	}
}

// @Tags SalaryTemplates
// @Summary 用id查询SalaryTemplates
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryTemplates true "用id查询SalaryTemplates"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /un/findSalaryTemplates [post]
func FindSalaryTemplates(c *gin.Context) {
	var st userSalary.SalaryTemplates
	_ = c.ShouldBindJSON(&st)
	err, rest := st.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"rest": rest,
		})
	}
}

// @Tags SalaryTemplates
// @Summary 分页获取SalaryTemplates列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取SalaryTemplates列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/getSalaryTemplatesList [post]
func GetSalaryTemplatesList(c *gin.Context) {
	type searchParams struct {
		userSalary.SalaryTemplates
		modelInterface.PageInfo
	}
	var sp searchParams
	_ = c.ShouldBindJSON(&sp)
	var enPriseID int
	ei, exist := c.Get("enpInfo")
	if exist {
		enpInfo := ei.(*userJobs.EnterpriseInfo)
		enPriseID = int(enpInfo.ID)
	}
	sp.SalaryTemplates.EnterpriseId = enPriseID
	err, list, total := sp.SalaryTemplates.GetInfoList(sp.PageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"stList":   list,
			"total":    total,
			"page":     sp.PageInfo.Page,
			"pageSize": sp.PageInfo.PageSize,
		})
	}
}

func GetSalaryTemplatesListBySearch(c *gin.Context) {
	query := c.Query("query")
	var enPriseID int
	ei, exist := c.Get("enpInfo")
	if exist {
		enpInfo := ei.(*userJobs.EnterpriseInfo)
		enPriseID = int(enpInfo.ID)
	} else {
		id, _ := strconv.Atoi(c.Query("id"))
		if id != 0 {
			err, _ := new(userJobs.EnterpriseInfo).GeteEpById(id)
			if err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("获取企业信息失败，%v", err), gin.H{})
				return
			}
			enPriseID = id
		}
	}
	st := &userSalary.SalaryTemplates{
		EnterpriseId: enPriseID,
	}
	err, list := st.GetAllInfoList(query)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"result": list,
		})
	}
}
