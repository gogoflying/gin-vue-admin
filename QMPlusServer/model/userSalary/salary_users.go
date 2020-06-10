// 自动生成模板SalaryUsers
package userSalary

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"strings"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type SalaryUsers struct {
	gorm.Model
	Openid       string   `json:"openid" gorm:"column:openid"`
	Name         string   `json:"name" gorm:"column:name"`
	AvatarUrl    string   `json:"avatarUrl" gorm:"column:avatarUrl;comment:'用户头像url'"`
	Mobile       string   `json:"mobile" gorm:"column:mobile"`
	PassWord     string   `json:"pass_word" gorm:"column:password"`
	NewPassWord  string   `json:"new_pass_word" gorm:"-"`
	JobName      string   `json:"job_name" gorm:"column:job_name;comment:'岗位名称'"`
	Salary       int      `json:"salary" gorm:"column:salary;comment:'工资'"`
	ContractDate int      `json:"contract_date" gorm:"column:contract_date;comment:'合同期限,单位是月'"`
	Card         string   `json:"card" gorm:"column:card;comment:'身份证号码'"`
	CardPhotos   string   `json:"card_photos" gorm:"column:card_photos;comment:'身份证正反面照片地址，分号分割'"`
	Cards        []string `json:"cards" gorm:"-"`
	OtherPhotos  string   `json:"other_photos" gorm:"column:other_photos;comment:'其他照片地址，分号分割'"`
	Others       []string `json:"others" gorm:"-"`
	BankCard     string   `json:"band_card" gorm:"column:band_cardz;comment:'银行卡正反面照片'"`
	Banks        []string `json:"banks" gorm:"-"`
	HuKou        string   `json:"hukou" gorm:"column:hukou;comment:'户口本照片'"`
	HuKous       []string `json:"hukous" gorm:"-"`
	Photo        string   `json:"photo" gorm:"column:photo;comment:'本人照片'"`
	Photos       []string `json:"photos" gorm:"-"`
	Checkup      string   `json:"checkup" gorm:"column:checkup;comment:'体检报告合格照片'"`
	Checkups     []string `json:"checkups" gorm:"-"`
	LeaveProve   string   `json:"leave_prove" gorm:"column:leave_prove;comment:'上家公司离职证明照片'"`
	Leaves       []string `json:"leaves" gorm:"-"`
	CaptchaId    string   `json:"captcha_id" gorm:"-"`
	Captcha      string   `json:"captcha" gorm:"-"`
	Code         string   `json:"code" gorm:"-"`
	OpType       int      `json:"op_type" gorm:"-"`
	Diploma      string   `json:"diploma" gorm:"column:diploma;comment:'毕业证照片'"`
	Diplomas     []string `json:"diplomas" gorm:"-"`
	Email        string   `json:"email" gorm:"column:email"`
	Enterprise   string   `json:"enterprise" gorm:"column:enterprise;comment:'入职企业'"`
	EnterpriseId int      `json:"enterprise_id" gorm:"column:enterprise_id;comment:'入职企业id'"`
	EnterTime    string   `json:"enter_time" gorm:"column:enter_time;comment:'入职日期'"`
	LeaveTime    string   `json:"leave_time" gorm:"column:LeaveTime;comment:'离职日期'"`
	JanYuan      string   `json:"jan_yuan" gorm:"column:JanYuan;comment:'社保减员日期'"`
	Reason       string   `json:"reason" gorm:"column:Reason;comment:'离职原因'"`
	Errors       string   `json:"errors" gorm:"column:errors;comment:'信息错误反馈内容'"`
	EnterStep    int      `json:"enter_step" gorm:"column:EnterStep;comment:'入职进度,0初始状态,待用户确认信息 1 用户基本信息已确认,待补充图片材料 2补充图片材料完毕 3待用户网签合同 4合同签约完毕'"`
	LeaveStep    int      `json:"leave_step" gorm:"column:LeaveStep;comment:'离职进度,0初始状态 1用户提交离职 2管理员审批，补充离职日期 3已离职，补充离职证明和减员年月具体到月就好'"`
}

func (un *SalaryUsers) GetUserByIdcardAndEmail() (err error, u SalaryUsers) {
	var ru SalaryUsers
	err = qmsql.DEFAULTDB.Model(un).Where("card = ? and email = ?", un.Card, un.Email).Find(&ru).Error
	return err, ru
}

func (un *SalaryUsers) GetUserByIdcardAndMobile() (err error, u SalaryUsers) {
	var ru SalaryUsers
	err = qmsql.DEFAULTDB.Model(un).Where("card = ? and mobile = ?", un.Card, un.Mobile).Find(&ru).Error
	return err, ru
}

func (un *SalaryUsers) UpdatePasswdByOpenid() (err error) {
	return qmsql.DEFAULTDB.Model(un).Where("openid = ?", un.Openid).Update("password", MD5PassWord(un.PassWord)).Error
}

// 创建SalaryUsers
func (un *SalaryUsers) CreateSalaryUsers() (err error) {
	bMd5 := md5.Sum([]byte(un.Card[len(un.Card)-6:]))
	un.PassWord = hex.EncodeToString(bMd5[:])
	un.Openid = NewOpenID()
	err = qmsql.DEFAULTDB.Create(un).Error
	return err
}

// 删除SalaryUsers
func (un *SalaryUsers) DeleteSalaryUsers() (err error) {
	err = qmsql.DEFAULTDB.Delete(un).Error
	return err
}
func (un *SalaryUsers) DeleteSalaryUsersEx() (err error) {
	err = qmsql.DEFAULTDB.Unscoped().Delete(un).Error
	return err
}
func (un *SalaryUsers) GetUserTx() *gorm.DB {
	return qmsql.DEFAULTDB
}

// 更新SalaryUsers
func (un *SalaryUsers) UpdateSalaryUsers() (err error, reun SalaryUsers) {
	//err = qmsql.DEFAULTDB.Save(un).Error
	//return err, *un

	db := qmsql.DEFAULTDB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var tmp SalaryUsers
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&tmp, un.ID).Error; err != nil {
		tx.Rollback()
		return err, *un
	}
	if err := tx.Save(un).Error; err != nil {
		tx.Rollback()
		return err, *un
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err, *un
	}
	return err, *un
}

func (un *SalaryUsers) UpdateLeaveStep() error {
	return qmsql.DEFAULTDB.Model(new(SalaryUsers)).Where("id = ?", un.ID).Update("LeaveStep", un.LeaveStep).Error
}

func (un *SalaryUsers) UpdateEnterStep() error {
	return qmsql.DEFAULTDB.Model(new(SalaryUsers)).Where("openid = ?", un.Openid).Update("EnterStep", un.EnterStep).Error
}

// 根据ID查看单条SalaryUsers
func (un *SalaryUsers) FindById() (err error, reun SalaryUsers) {
	err = qmsql.DEFAULTDB.Where("id = ?", un.ID).First(&reun).Error
	return err, reun
}

func (un *SalaryUsers) LoginMobile() (err error, reun SalaryUsers) {
	err = qmsql.DEFAULTDB.Model(un).Where("mobile = ? and password = ?", un.Mobile, un.PassWord).Find(&reun).Error
	return
}

func (un *SalaryUsers) UpdatePassword() (err error) {
	err = qmsql.DEFAULTDB.Model(un).Where("openid = ? and password = ?", un.Openid, un.PassWord).Update("password", un.NewPassWord).Error
	return
}

func MD5PassWord(ps string) string {
	bMd5 := md5.Sum([]byte(ps))
	return hex.EncodeToString(bMd5[:])
}

// 根据ID查看单条SalaryUsers
func (un *SalaryUsers) FindByOpenid() (err error, reun SalaryUsers) {
	err = qmsql.DEFAULTDB.Where("openid = ?", un.Openid).First(&reun).Error
	return err, reun
}

// 根据ID查看单条SalaryUsers
func (un *SalaryUsers) FindByEnterprise() (err error, reun SalaryUsers) {
	err = qmsql.DEFAULTDB.Where("enterprise = ? AND enterprise_id = ? AND name = ? ", un.Enterprise, un.EnterpriseId, un.Name).First(&reun).Error
	return err, reun
}

// 根据ID查看单条SalaryUsers
func (un *SalaryUsers) FindByMobile() (err error, reun SalaryUsers) {
	err = qmsql.DEFAULTDB.Where("mobile = ?", un.Mobile).First(&reun).Error
	return err, reun
}

// 分页获取SalaryUsers
func (un *SalaryUsers) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(un, info)
	if err != nil {
		return
	} else {
		reSalaryUsersList := make([]*SalaryUsers, 0)
		model := qmsql.DEFAULTDB.Model(info)
		if un.EnterpriseId != 0 {
			model = model.Where("enterprise_id = ?", un.EnterpriseId)
			db = db.Where("enterprise_id = ?", un.EnterpriseId)
		}
		if un.Name != "" {
			model = model.Where("name LIKE ?", "%"+un.Name+"%")
			db = db.Where("name LIKE ?", "%"+un.Name+"%")
		}
		if un.Card != "" {
			model = model.Where("card LIKE ?", "%"+un.Card+"%")
			db = db.Where("card LIKE ?", "%"+un.Card+"%")
		}
		if un.Mobile != "" {
			model = model.Where("mobile LIKE ?", "%"+un.Mobile+"%")
			db = db.Where("mobile LIKE ?", "%"+un.Mobile+"%")
		}

		err = model.Find(&reSalaryUsersList).Count(&total).Error
		if err != nil {
			return err, reSalaryUsersList, total
		} else {
			err = db.Find(&reSalaryUsersList).Error
		}
		if err == nil {
			for _, reun := range reSalaryUsersList {
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
			}
		}
		return err, reSalaryUsersList, total
	}
}

func NewOpenID() string {
	buf := bytes.Buffer{}
	u := uuid.NewV4()

	buf.Write(u.Bytes())

	dst := md5.Sum(buf.Bytes())
	return hex.EncodeToString(dst[:])
}
