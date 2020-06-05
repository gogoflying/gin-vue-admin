package api

import (
	"fmt"
	"gin-vue-admin/model/userJobs"

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
	ei, exist := c.Get("enpInfo")
	if exist {
		enpInfo := ei.(*userJobs.EnterpriseInfo)
		un.Enterprise = enpInfo.EnterPriseName
		un.EnterpriseId = int(enpInfo.ID)
	}
	fmt.Printf("checkSalarysUserInTable %v\n", un)
	openid, err := checkSalarysUserInTable(un.Name, un.Enterprise, un.EnterpriseId)
	if len(openid) == 0 {
		fmt.Errorf("checkSalarysUserInTable user:%s not in table salaryUser", un.Name)
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
		return
	}
	fmt.Printf("checkSalarysUserInTable openid:%s\n", openid)
	un.Openid = openid

	err = un.CreateSalarys()
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
	err := un.DeleteSalarysEx()
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
	ei, exist := c.Get("enpInfo")
	if exist {
		enpInfo := ei.(*userJobs.EnterpriseInfo)
		un.Enterprise = enpInfo.EnterPriseName
		un.EnterpriseId = int(enpInfo.ID)
	}
	openid, err := checkSalarysUserInTable(un.Name, un.Enterprise, un.EnterpriseId)
	if len(openid) == 0 {
		fmt.Errorf("checkSalarysUserInTable user:%s not in table salaryUser", un.Name)
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
		return
	}

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

	type searchParams struct {
		userSalary.Salarys
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
	sp.Salarys.EnterpriseId = enPriseID

	err, list, total := sp.Salarys.GetInfoList(sp.PageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"userSalaryList": list,
			"total":          total,
			"page":           sp.PageInfo.Page,
			"pageSize":       sp.PageInfo.PageSize,
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
	var (
		id     int
		epName string
	)

	ei, exist := c.Get("enpInfo")
	if exist {
		enpInfo := ei.(*userJobs.EnterpriseInfo)
		epName = enpInfo.EnterPriseName
		id = int(enpInfo.ID)
	} else {
		id, _ = strconv.Atoi(c.Query("id"))
		err, ep := new(userJobs.EnterpriseInfo).GeteEpById(id)
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("获取企业信息失败，%v", err), gin.H{})
			return
		}
		epName = ep.EnterPriseName
	}
	fmt.Printf("ImportSalarys enter: %v--%v\n", id, epName)
	_, fxlsx, err := c.Request.FormFile("file")
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("上传文件失败，%v", err), gin.H{})
	} else {
		//文件上传后拿到文件路径
		err := UploadFileLocalEx(fxlsx, id, epName)
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("接收返回值失败，%v", err), gin.H{})
		} else {
			servers.ReportFormat(c, true, "上传成功", gin.H{})
		}
	}
}

// 上传文件到本地
func UploadFileLocalEx(file *multipart.FileHeader, id int, ep string) (err error) {
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
	err = batchInsertSalarys(fileTemp, id, ep)
	if err != nil {
		servers.DelLocalFile(tmpFilename)
		return err
	}
	servers.DelLocalFile(tmpFilename)
	fmt.Printf("batchInsertSalarys :%v", tmpFilename)
	return
}

func batchInsertSalarys(file *xlsx.File, id int, ep string) error {
	fmt.Printf("batchInsertSalarys: %v--%v\n", id, ep)
	var un userSalary.Salarys

	tx := un.GetDbTx().Begin()
	defer tx.Rollback()
	for _, sheet := range file.Sheets {
		fmt.Printf("Sheet Name: %s\n", sheet.Name)

		for rowIndex, row := range sheet.Rows {
			if rowIndex == 0 {
				continue
			}
			if ep != strings.Trim(row.Cells[1].Value, " ") {
				fmt.Errorf("batchInsertSalarys excel name:%s,import target name:%s", strings.Trim(row.Cells[1].Value, " "), ep)
				return fmt.Errorf("Improt enterprise name not eq")
			}

			name := strings.Trim(row.Cells[2].Value, " ")
			openid, _ := checkSalarysUserInTable(name, ep, id)
			if len(openid) == 0 {
				fmt.Errorf("batchInsertSalarys user:%s not in table salaryUser", name)
				return fmt.Errorf("user not in table salaryUser")
			}
			year, _ := strconv.Atoi(strings.Trim(row.Cells[3].Value, " "))
			month, _ := strconv.Atoi(strings.Trim(row.Cells[4].Value, " "))
			base := strings.Trim(row.Cells[5].Value, " ")
			gangwei := strings.Trim(row.Cells[6].Value, " ")
			xzhj := strings.Trim(row.Cells[7].Value, " ")
			yjtc := strings.Trim(row.Cells[8].Value, " ")
			jjjs := strings.Trim(row.Cells[9].Value, " ")
			fdxs := strings.Trim(row.Cells[10].Value, " ")
			ydjj := strings.Trim(row.Cells[11].Value, " ")
			gzts := strings.Trim(row.Cells[12].Value, " ")
			jbf := strings.Trim(row.Cells[13].Value, " ")
			txbt := strings.Trim(row.Cells[14].Value, " ")
			csbt := strings.Trim(row.Cells[15].Value, " ")
			jtbt := strings.Trim(row.Cells[16].Value, " ")
			qtbt := strings.Trim(row.Cells[17].Value, " ")
			bthj := strings.Trim(row.Cells[18].Value, " ")
			qtjq := strings.Trim(row.Cells[19].Value, " ")
			njts := strings.Trim(row.Cells[20].Value, " ")
			cdkk := strings.Trim(row.Cells[21].Value, " ")
			bjts := strings.Trim(row.Cells[22].Value, " ")
			bjkk := strings.Trim(row.Cells[23].Value, " ")
			sjts := strings.Trim(row.Cells[24].Value, " ")
			sjkk := strings.Trim(row.Cells[25].Value, " ")
			kkhj := strings.Trim(row.Cells[26].Value, " ")
			yftz := strings.Trim(row.Cells[27].Value, " ")
			byyf := strings.Trim(row.Cells[28].Value, " ")
			dkwx := strings.Trim(row.Cells[29].Value, " ")
			gjj := strings.Trim(row.Cells[30].Value, " ")
			dkgs := strings.Trim(row.Cells[31].Value, " ")
			sfgz := strings.Trim(row.Cells[32].Value, " ")

			ext1 := strings.Trim(row.Cells[33].Value, " ")
			ext2 := strings.Trim(row.Cells[34].Value, " ")
			ext3 := strings.Trim(row.Cells[35].Value, " ")
			ext4 := strings.Trim(row.Cells[36].Value, " ")
			ext5 := strings.Trim(row.Cells[37].Value, " ")
			ext6 := strings.Trim(row.Cells[38].Value, " ")
			ext7 := strings.Trim(row.Cells[39].Value, " ")
			ext8 := strings.Trim(row.Cells[40].Value, " ")
			ext9 := strings.Trim(row.Cells[41].Value, " ")
			ext10 := strings.Trim(row.Cells[42].Value, " ")
			//salary, _ := strconv.Atoi(strings.Trim(row.Cells[4].Value, " "))
			//contract_date, _ := strconv.Atoi(strings.Trim(row.Cells[5].Value, " "))
			un = userSalary.Salarys{
				Openid:       openid,
				Enterprise:   ep,
				EnterpriseId: id,
				Name:         name,
				Year:         year,
				Month:        month,
				Base:         base,
				Gangwei:      gangwei,
				Xzhj:         xzhj,
				Yjtc:         yjtc,
				Jjjs:         jjjs,
				Fdxs:         fdxs,
				Ydjj:         ydjj,
				Gzts:         gzts,
				Jbf:          jbf,
				Txbt:         txbt,
				Csbt:         csbt,
				Jtbt:         jtbt,
				Qtbt:         qtbt,
				Bthj:         bthj,
				Qtjq:         qtjq,
				Njts:         njts,
				Cdkk:         cdkk,
				Bjts:         bjts,
				Bjkk:         bjkk,
				Sjts:         sjts,
				Sjkk:         sjkk,
				Kkhj:         kkhj,
				Yftz:         yftz,
				Byyf:         byyf,
				Dkwx:         dkwx,
				Gjj:          gjj,
				Dkgs:         dkgs,
				Sfgz:         sfgz,
				Ext1:         ext1,
				Ext2:         ext2,
				Ext3:         ext3,
				Ext4:         ext4,
				Ext5:         ext5,
				Ext6:         ext6,
				Ext7:         ext7,
				Ext8:         ext8,
				Ext9:         ext9,
				Ext10:        ext10,
			}
			//err := un.CreateSalarys()
			err := tx.Create(un)
			if err != nil {
				return fmt.Errorf("insert import err")
			}
		}
	}
	tx.Commit()

	return nil
}

func checkSalarysUserInTable(name, enterprise string, enterpriseID int) (string, error) {
	if name == "" || enterprise == "" {
		return "", fmt.Errorf("param empty")
	}
	//var un userSalary.SalaryUsers
	salaryUser := userSalary.SalaryUsers{
		Name:         name,
		Enterprise:   enterprise,
		EnterpriseId: enterpriseID,
	}
	_, newSalary := salaryUser.FindByEnterprise()
	if newSalary.Openid != "" {
		return newSalary.Openid, nil
	} else {
		return "", fmt.Errorf("salary user no this user")
	}
}
