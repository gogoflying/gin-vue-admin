package api

import (
	"fmt"
	//"bytes"
	//"io"
	"io/ioutil"
	"mime/multipart"

	//"os"
	"strconv"
	"strings"
	"time"

	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userSalary"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

// @Tags Salarys
// @Summary 创建Salarys
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.Salarys true "创建Salarys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/createSalarys [post]
func CreateSalarys(c *gin.Context) {
	var un userSalary.Salarys
	_ = c.ShouldBindJSON(&un)
	err := un.CreateSalarys()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags Salarys
// @Summary 删除Salarys
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.Salarys true "删除Salarys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /un/deleteSalarys [post]
func DeleteSalarys(c *gin.Context) {
	var un userSalary.Salarys
	_ = c.ShouldBindJSON(&un)
	err := un.DeleteSalarys()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags Salarys
// @Summary 更新Salarys
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.Salarys true "更新Salarys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /un/updateSalarys [post]
func UpdateSalarys(c *gin.Context) {
	var un userSalary.Salarys
	_ = c.ShouldBindJSON(&un)
	err, reun := un.UpdateSalarys()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reun": reun,
		})
	}
}

// @Tags Salarys
// @Summary 用id查询Salarys
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.Salarys true "用id查询Salarys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /un/findSalarys [post]
func FindSalarys(c *gin.Context) {
	var un userSalary.Salarys
	_ = c.ShouldBindJSON(&un)
	err, reun := un.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reun": reun,
		})
	}
}

func FindSalarysByIdAndOpenid(c *gin.Context) {
	var un userSalary.Salarys
	_ = c.ShouldBindJSON(&un)
	err, reun := un.FindByIdAndOpenid()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reun": reun,
		})
	}
}

// @Tags Salarys
// @Summary 分页获取Salarys列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取Salarys列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/getSalarysList [post]
func GetSalarysList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userSalary.Salarys).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"userSalaryList": list,
			"total":          total,
			"page":           pageInfo.Page,
			"pageSize":       pageInfo.PageSize,
		})
	}
}

type ReqSalarysListSearch struct {
	Openid   string `json:"openid"`
	Page     int
	PageSize int
	YearL    int `json:"year_low"`
	YearH    int `json:"year_high"`
	MonthL   int `json:"month_low"`
	MonthH   int `json:"month_high"`
}

func GetSalarysListSearch(c *gin.Context) {
	var req ReqSalarysListSearch
	_ = c.ShouldBindJSON(&req)
	err, list := new(userSalary.Salarys).GetInfoListSearch(req.Openid, req.YearL, req.MonthL, req.YearH, req.MonthH, req.Page, req.PageSize)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"userSalaryList": list,
			"page":           req.Page,
			"pageSize":       req.PageSize,
		})
	}
}

func ImportSalarys(c *gin.Context) {

	_, fxlsx, err := c.Request.FormFile("file")

	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("上传文件失败，%v", err), gin.H{})
	} else {
		//文件上传后拿到文件路径
		err := UploadFileLocalEx(fxlsx)
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("接收返回值失败，%v", err), gin.H{})
		} else {
			servers.ReportFormat(c, true, "上传成功", gin.H{})
		}
	}
}

// 上传文件到本地
func UploadFileLocalEx(file *multipart.FileHeader) (err error) {
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()
	xlsxBuf, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	filename := "./tmp/salarys"
	timestring := time.Now().Format("20060102150405")
	err = saveXlsxFile(filename, timestring, xlsxBuf)
	if err != nil {
		return err
	}
	tmpFilename := filename + timestring + ".xlsx"
	fileTemp, err := xlsx.OpenFile(tmpFilename)
	if err != nil {
		return err
	}
	err = batchInsertSalarys(fileTemp)
	if err != nil {
		return err
	}
	servers.DelLocalFile(tmpFilename)
	fmt.Printf("batchInsertSalarys :%v", tmpFilename)
	return
}

func batchInsertSalarys(file *xlsx.File) error {
	for _, sheet := range file.Sheets {
		fmt.Printf("Sheet Name: %s\n", sheet.Name)
		for rowIndex, row := range sheet.Rows {
			if rowIndex == 0 {
				continue
			}
			name := strings.Trim(row.Cells[1].Value, " ")
			year, _ := strconv.Atoi(strings.Trim(row.Cells[2].Value, " "))
			month, _ := strconv.Atoi(strings.Trim(row.Cells[3].Value, " "))
			gangwei := strings.Trim(row.Cells[4].Value, " ")
			xzhj := strings.Trim(row.Cells[5].Value, " ")
			yjtc := strings.Trim(row.Cells[6].Value, " ")
			jjjs := strings.Trim(row.Cells[7].Value, " ")
			fdxs := strings.Trim(row.Cells[8].Value, " ")
			ydjj := strings.Trim(row.Cells[9].Value, " ")
			gzts := strings.Trim(row.Cells[10].Value, " ")
			jbf := strings.Trim(row.Cells[11].Value, " ")
			txbt := strings.Trim(row.Cells[12].Value, " ")
			csbt := strings.Trim(row.Cells[13].Value, " ")
			jtbt := strings.Trim(row.Cells[14].Value, " ")
			qtbt := strings.Trim(row.Cells[15].Value, " ")
			bthj := strings.Trim(row.Cells[16].Value, " ")
			qtjq := strings.Trim(row.Cells[17].Value, " ")
			njts := strings.Trim(row.Cells[18].Value, " ")
			cdkk := strings.Trim(row.Cells[19].Value, " ")
			bjts := strings.Trim(row.Cells[20].Value, " ")
			bjkk := strings.Trim(row.Cells[21].Value, " ")
			sjts := strings.Trim(row.Cells[22].Value, " ")
			sjkk := strings.Trim(row.Cells[23].Value, " ")
			kkhj := strings.Trim(row.Cells[24].Value, " ")
			yftz := strings.Trim(row.Cells[25].Value, " ")
			byyf := strings.Trim(row.Cells[26].Value, " ")
			dkwx := strings.Trim(row.Cells[27].Value, " ")
			gjj := strings.Trim(row.Cells[28].Value, " ")
			dkgs := strings.Trim(row.Cells[29].Value, " ")
			sfgz := strings.Trim(row.Cells[30].Value, " ")
			//salary, _ := strconv.Atoi(strings.Trim(row.Cells[4].Value, " "))
			//contract_date, _ := strconv.Atoi(strings.Trim(row.Cells[5].Value, " "))
			un := userSalary.Salarys{
				Name:    name,
				Year:    year,
				Month:   month,
				Gangwei: gangwei,
				Xzhj:    xzhj,
				Yjtc:    yjtc,
				Jjjs:    jjjs,
				Fdxs:    fdxs,
				Ydjj:    ydjj,
				Gzts:    gzts,
				Jbf:     jbf,
				Txbt:    txbt,
				Csbt:    csbt,
				Jtbt:    jtbt,
				Qtbt:    qtbt,
				Bthj:    bthj,
				Qtjq:    qtjq,
				Njts:    njts,
				Cdkk:    cdkk,
				Bjts:    bjts,
				Bjkk:    bjkk,
				Sjts:    sjts,
				Sjkk:    sjkk,
				Kkhj:    kkhj,
				Yftz:    yftz,
				Byyf:    byyf,
				Dkwx:    dkwx,
				Gjj:     gjj,
				Dkgs:    dkgs,
				Sfgz:    sfgz,
			}
			err := un.CreateSalarys()
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
	return nil
}
