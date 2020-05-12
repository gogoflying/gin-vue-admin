package api

import (
	"bytes"
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"
	"gin-vue-admin/model/userSalary"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

// @Tags SalaryUsers
// @Summary 创建SalaryUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryUsers true "创建SalaryUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/createSalaryUsers [post]
func CreateSalaryUsers(c *gin.Context) {
	var un userSalary.SalaryUsers
	_ = c.ShouldBindJSON(&un)
	err := un.CreateSalaryUsers()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SalaryUsers
// @Summary 删除SalaryUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryUsers true "删除SalaryUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /un/deleteSalaryUsers [post]
func DeleteSalaryUsers(c *gin.Context) {
	var un userSalary.SalaryUsers
	_ = c.ShouldBindJSON(&un)
	err := un.DeleteSalaryUsers()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SalaryUsers
// @Summary 更新SalaryUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryUsers true "更新SalaryUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /un/updateSalaryUsers [post]
func UpdateSalaryUsers(c *gin.Context) {
	var un userSalary.SalaryUsers
	_ = c.ShouldBindJSON(&un)
	err, reun := un.UpdateSalaryUsers()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reun": reun,
		})
	}
}

func UpdateSalaryUsersLeaveStep(c *gin.Context) {
	var un userSalary.SalaryUsers
	_ = c.ShouldBindJSON(&un)
	err := un.UpdateLeaveStep()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{})
	}
}

func UpdateSalaryUsersEnterStep(c *gin.Context) {
	var un userSalary.SalaryUsers
	_ = c.ShouldBindJSON(&un)
	err := un.UpdateEnterStep()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{})
	}
}

// @Tags SalaryUsers
// @Summary 用id查询SalaryUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryUsers true "用id查询SalaryUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /un/findSalaryUsers [post]
func FindSalaryUsers(c *gin.Context) {
	var un userSalary.SalaryUsers
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

func SalaryUserUpdatePassword(c *gin.Context) {
	var un userSalary.SalaryUsers
	_ = c.ShouldBindJSON(&un)
	err := un.UpdatePassword()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{})
	}
}

func SalaryUserLoginMobile(c *gin.Context) {
	var un userSalary.SalaryUsers
	_ = c.ShouldBindJSON(&un)
	err, reun := un.LoginMobile()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"openid":   reun.Openid,
			"username": reun.Name,
		})
	}
}

func FindSalaryUsersByOpenid(c *gin.Context) {
	var un userSalary.SalaryUsers
	_ = c.ShouldBindJSON(&un)
	err, reun := un.FindByOpenid()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		if len(reun.CardPhotos) > 0 {
			reun.Cards = strings.Split(reun.CardPhotos, ";")
		}
		if len(reun.BankCard) > 0 {
			reun.Banks = strings.Split(reun.BankCard, ";")
		}
		if len(reun.Photo) > 0 {
			reun.Photos = strings.Split(reun.Photo, ";")
		}
		if len(reun.Checkup) > 0 {
			reun.Checkups = strings.Split(reun.Checkup, ";")
		}
		if len(reun.Diploma) > 0 {
			reun.Diplomas = strings.Split(reun.Diploma, ";")
		}
		if len(reun.LeaveProve) > 0 {
			reun.Leaves = strings.Split(reun.LeaveProve, ";")
		}
		if len(reun.HuKou) > 0 {
			reun.HuKous = strings.Split(reun.HuKou, ";")
		}
		if len(reun.OtherPhotos) > 0 {
			reun.Others = strings.Split(reun.OtherPhotos, ";")
		}
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reun": reun,
		})
	}
}

// @Tags SalaryUsers
// @Summary 分页获取SalaryUsers列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取SalaryUsers列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/getSalaryUsersList [post]
func GetSalaryUsersList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userSalary.SalaryUsers).GetInfoList(pageInfo)
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

func ImportSalaryUsers(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	err, ep := new(userJobs.EnterpriseInfo).GeteEpById(id)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("上传文件失败，%v", err), gin.H{})
	} else {
		_, fxlsx, err := c.Request.FormFile("file")
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("上传文件失败，%v", err), gin.H{})
		} else {
			//文件上传后拿到文件路径
			err := UploadFileLocal(fxlsx, id, ep.EnterPriseName)
			if err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("接收返回值失败，%v", err), gin.H{})
			} else {
				servers.ReportFormat(c, true, "上传成功", gin.H{})
			}
		}
	}
}

// 上传文件到本地
func UploadFileLocal(file *multipart.FileHeader, id int, ep string) (err error) {
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()
	xlsxBuf, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	filename := "./tmp/salaryuser"
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
	err = batchInsertSalaryUser(fileTemp, id, ep)
	if err != nil {
		servers.DelLocalFile(tmpFilename)
		return err
	}
	servers.DelLocalFile(tmpFilename)
	return
}

func saveXlsxFile(filename, time string, src []byte) error {
	out, err := os.Create(filename + time + ".xlsx")
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, bytes.NewBuffer(src))
	if err != nil {
		return err
	}
	return nil
}

func batchInsertSalaryUser(file *xlsx.File, id int, ep string) error {

	for _, sheet := range file.Sheets {
		fmt.Printf("Sheet Name: %s\n", sheet.Name)
		for rowIndex, row := range sheet.Rows {
			if rowIndex == 0 {
				continue
			}
			name := strings.Trim(row.Cells[1].Value, " ")
			mobile := strings.Trim(row.Cells[2].Value, " ")
			jobname := strings.Trim(row.Cells[3].Value, " ")
			salary, _ := strconv.Atoi(strings.Trim(row.Cells[4].Value, " "))
			contract_date, _ := strconv.Atoi(strings.Trim(row.Cells[5].Value, " "))
			card := strings.Trim(row.Cells[6].Value, " ")
			email := strings.Trim(row.Cells[7].Value, " ")
			un := userSalary.SalaryUsers{
				Name:         name,
				Mobile:       mobile,
				JobName:      jobname,
				Salary:       salary,
				ContractDate: contract_date,
				Card:         card,
				Email:        email,
				EnterpriseId: id,
				Enterprise:   ep,
			}
			err := un.CreateSalaryUsersTx()
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
	return nil
}
