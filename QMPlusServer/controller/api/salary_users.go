package api

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"
	"gin-vue-admin/model/userSalary"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dchest/captcha"
	uuid "github.com/satori/go.uuid"

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
	ei, exist := c.Get("enpInfo")
	if exist {
		enpInfo := ei.(*userJobs.EnterpriseInfo)
		un.Enterprise = enpInfo.EnterPriseName
		un.EnterpriseId = int(enpInfo.ID)
	}
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
	err := un.DeleteSalaryUsersEx()
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

func SalaryUserSendEmail(c *gin.Context) {
	var req userSalary.SalaryUsers
	_ = c.ShouldBindJSON(&req)
	if captcha.VerifyString(req.CaptchaId, req.Captcha) {
		err, _ := req.GetUserByIdcardAndEmail()
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("用户不存在"), gin.H{})
			return
		}
		code := GenerateCode(4)
		servers.AddMap(req.Email, code)
		if err := servers.SendEmail(req.Email, "密码找回", "你正在使用邮箱找回密码功能，本次邮箱验证码为:"+code); err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("发送邮箱验证码失败", err), gin.H{})
			return
		}
		servers.ReportFormat(c, true, "已发送邮箱验证码", gin.H{})
	} else {
		servers.ReportFormat(c, false, "图片验证码错误", gin.H{})
	}
}

func SalaryUserSendSMS(c *gin.Context) {
	var req userSalary.SalaryUsers

	_ = c.ShouldBindJSON(&req)
	if captcha.VerifyString(req.CaptchaId, req.Captcha) {
		err, _ := req.GetUserByIdcardAndMobile()
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("用户不存在"), gin.H{})
			return
		}
		code := GenerateCode(4)
		servers.AddMap(req.Mobile, code)
		if err := servers.SendSms(req.Mobile, code, servers.SMSMSGTYPE_VERIFY); err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("发送短信失败"), gin.H{})
			return
		}
		servers.ReportFormat(c, true, "已发送验证码", gin.H{})
	} else {
		servers.ReportFormat(c, false, "图片验证码错误", gin.H{})
	}
}

func SalaryUserFindPass(c *gin.Context) {
	var req userSalary.SalaryUsers

	_ = c.ShouldBindJSON(&req)
	if req.OpType == 0 {
		if req.Code != "" && req.Code == servers.GetMap(req.Email).(string) {
			err, reun := req.GetUserByIdcardAndEmail()
			if err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("用户不存在"), gin.H{})
				return
			}
			req.Openid = reun.Openid
			if err = req.UpdatePasswdByOpenid(); err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("修改密码失败"), gin.H{})
				return
			}
			servers.DelMap(req.Email)
			servers.ReportFormat(c, true, "密码重置成功", gin.H{})
		} else {
			servers.ReportFormat(c, false, fmt.Sprintf("邮箱验证码错误"), gin.H{})
			return
		}

	} else {
		if req.Code != "" && req.Code == servers.GetMap(req.Mobile).(string) {
			err, reun := req.GetUserByIdcardAndMobile()
			if err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("用户不存在"), gin.H{})
				return
			}
			req.Openid = reun.Openid
			if err = req.UpdatePasswdByOpenid(); err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("修改密码失败"), gin.H{})
				return
			}
			servers.DelMap(req.Mobile)
			servers.ReportFormat(c, true, "密码重置成功", gin.H{})
		} else {
			servers.ReportFormat(c, false, fmt.Sprintf("手机验证码错误"), gin.H{})
			return
		}
	}
}

func GenerateCode(n int) string {
	var s string
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		s = s + fmt.Sprintf("%v", rand.Intn(10))
	}
	return s
}

func SalaryUserUpdatePassword(c *gin.Context) {
	var un userSalary.SalaryUsers
	var err error
	_ = c.ShouldBindJSON(&un)
	bMd5 := md5.Sum([]byte(un.PassWord))
	un.PassWord = hex.EncodeToString(bMd5[:])
	if len(un.NewPassWord) >= 6 {
		bMd51 := md5.Sum([]byte(un.NewPassWord))
		un.NewPassWord = hex.EncodeToString(bMd51[:])
		err = un.UpdatePassword()
	} else {
		err = fmt.Errorf("新密码长度小于6位")
	}

	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{})
	}
}

func SalaryUserLoginMobile(c *gin.Context) {
	var un userSalary.SalaryUsers
	_ = c.ShouldBindJSON(&un)
	bMd5 := md5.Sum([]byte(un.PassWord))
	un.PassWord = hex.EncodeToString(bMd5[:])
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
	type searchParams struct {
		userSalary.SalaryUsers
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
	sp.SalaryUsers.EnterpriseId = enPriseID
	err, list, total := sp.SalaryUsers.GetInfoList(sp.PageInfo)
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

func ImportSalaryUsers(c *gin.Context) {
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

	_, fxlsx, err := c.Request.FormFile("file")
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("上传文件失败，%v", err), gin.H{})
	} else {
		//文件上传后拿到文件路径
		err := UploadFileLocal(fxlsx, id, epName)
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("接收返回值失败，%v", err), gin.H{})
		} else {
			servers.ReportFormat(c, true, "上传成功", gin.H{})
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
	var un userSalary.SalaryUsers
	tx := un.GetUserTx().Begin()
	defer tx.Rollback()

	for _, sheet := range file.Sheets {
		fmt.Printf("Sheet Name: %s\n", sheet.Name)
		for rowIndex, row := range sheet.Rows {
			if rowIndex == 0 {
				continue
			}
			if ep != strings.Trim(row.Cells[1].Value, " ") {
				fmt.Errorf("batchInsertSalaryUser excel name:%s,import target name:%s", strings.Trim(row.Cells[1].Value, " "), ep)
				return fmt.Errorf("batchInsertSalaryUser Improt enterprise nmae not eq")
			}
			name := strings.Trim(row.Cells[2].Value, " ")
			mobile := strings.Trim(row.Cells[3].Value, " ")
			jobname := strings.Trim(row.Cells[4].Value, " ")
			salary, _ := strconv.Atoi(strings.Trim(row.Cells[5].Value, " "))
			contract_date, _ := strconv.Atoi(strings.Trim(row.Cells[6].Value, " "))
			card := strings.Trim(row.Cells[7].Value, " ")
			email := strings.Trim(row.Cells[8].Value, " ")
			openid := NewOpenID()
			bMd5 := md5.Sum([]byte(card[len(card)-6:]))
			passwd := hex.EncodeToString(bMd5[:])
			un := userSalary.SalaryUsers{
				Openid:       openid,
				Name:         name,
				Mobile:       mobile,
				JobName:      jobname,
				Salary:       salary,
				ContractDate: contract_date,
				Card:         card,
				Email:        email,
				EnterpriseId: id,
				Enterprise:   ep,
				PassWord:     passwd,
			}

			err := tx.Create(un).Error
			if err != nil {
				return fmt.Errorf("import insert salaryuser err")
			}
		}
	}
	tx.Commit()
	return nil
}

func NewOpenID() string {
	buf := bytes.Buffer{}
	u := uuid.NewV4()

	buf.Write(u.Bytes())

	dst := md5.Sum(buf.Bytes())
	return hex.EncodeToString(dst[:])
}
