package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

// @Tags ResumeStatus
// @Summary 创建ResumeStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.ResumeStatus true "创建ResumeStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rs/createResumeStatus [post]
func CreateResumeStatus(c *gin.Context) {
	var rs userJobs.ResumeStatus
	_ = c.ShouldBindJSON(&rs)
	list, err := rs.CreateResumeStatus()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{
			"results": list,
		})
	}
}

// @Tags ResumeStatus
// @Summary 删除ResumeStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.ResumeStatus true "删除ResumeStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rs/deleteResumeStatus [post]
func DeleteResumeStatus(c *gin.Context) {
	var rs userJobs.ResumeStatus
	_ = c.ShouldBindJSON(&rs)
	err := rs.DeleteResumeStatus()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags ResumeStatus
// @Summary 更新ResumeStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.ResumeStatus true "更新ResumeStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rs/updateResumeStatus [post]
func UpdateResumeStatus(c *gin.Context) {
	var rs userJobs.ResumeStatus
	_ = c.ShouldBindJSON(&rs)
	err, rers := rs.UpdateResumeStatus()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"rers": rers,
		})
	}
}

func UpdateResumeMemo(c *gin.Context) {
	var rs userJobs.ResumeStatus
	_ = c.ShouldBindJSON(&rs)
	err := rs.UpdateMemo()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{})
	}
}

func UpdateResumeRemark(c *gin.Context) {
	var rs userJobs.ResumeStatus
	_ = c.ShouldBindJSON(&rs)
	err := rs.UpdateRemark()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{})
	}
}

// @Tags ResumeStatus
// @Summary 用id查询ResumeStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.ResumeStatus true "用id查询ResumeStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rs/findResumeStatus [post]
func FindResumeStatus(c *gin.Context) {
	var rs userJobs.ResumeStatus
	_ = c.ShouldBindJSON(&rs)
	err, rers := rs.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"rers": rers,
		})
	}
}

// @Tags ResumeStatus
// @Summary 分页获取ResumeStatus列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取ResumeStatus列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rs/getResumeStatusList [post]
func GetResumeStatusList(c *gin.Context) {
	type searchParams struct {
		userJobs.ResumeStatus
		modelInterface.PageInfo
	}
	var sp searchParams
	_ = c.ShouldBindJSON(&sp)
	var enPriseID int
	ei, exist := c.Get("enpInfo")
	if exist {
		enpInfo := ei.(*userJobs.EnterpriseInfo)
		enPriseID = int(enpInfo.ID)
	} else {
		//if c.Query("id") != "" {
		//	id, _ := strconv.Atoi(c.Query("id"))
		//	if id != 0 {
		//		err, _ := new(userJobs.EnterpriseInfo).GeteEpById(id)
		//		if err != nil {
		//			servers.ReportFormat(c, false, fmt.Sprintf("获取企业信息失败，%v", err), gin.H{})
		//			return
		//		}
		//		enPriseID = id
		//	}
		//}
	}
	sp.ResumeStatus.CompanyId = enPriseID
	err, list, total := sp.ResumeStatus.GetInfoList(sp.PageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"resumeStatusList": list,
			"total":            total,
			"page":             sp.PageInfo.Page,
			"pageSize":         sp.PageInfo.PageSize,
		})
	}
}

type ReqResumeByStatus struct {
	OpenId   string `json:"openid"`
	Status   int    `json:"status"`
	Page     int    `json:"Page"`
	PageSize int    `json:"PageSize"`
}

func GetResumeByStatus(c *gin.Context) {
	var req ReqResumeByStatus
	_ = c.ShouldBindJSON(&req)
	err, rs, total := new(userJobs.ResumeStatus).GetResumeStatus(req.OpenId, req.Status, req.Page, req.PageSize)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"results":  rs.Results,
			"total":    total,
			"page":     req.Page,
			"pageSize": req.PageSize,
		})
	}
}

func JobEnterStatus(c *gin.Context) {
	var req userJobs.ResumeStatus
	_ = c.ShouldBindJSON(&req)
	err := req.EnterJob()
	if err != nil {
		servers.ReportFormat(c, false, err.Error(), gin.H{})
	} else {
		servers.ReportFormat(c, true, "", gin.H{})
	}
}
