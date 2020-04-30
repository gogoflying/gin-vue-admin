// 自动生成模板SalaryUsers
package userSalary

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type SalaryUsers struct {
	gorm.Model
	Openid       string `json:"openid" gorm:"column:openid"`
	Name         string `json:"name" gorm:"column:name"`
	Mobile       string `json:"mobile" gorm:"column:mobile"`
	Card         string `json:"card" gorm:"column:card;comment:'身份证号码'"`
	Email        string `json:"email" gorm:"column:email"`
	Enterprise   string `json:"enterprise" gorm:"column:enterprise";comment:'入职企业'`
	EnterpriseId int    `json:"enterprise_id" gorm:"column:enterprise_id;comment:'入职企业id'"`
	EnterTime    string `json:"enter_time" gorm:"column:enter_time;comment:'入职日期'"`
	LeaveTime    string `json:"leave_time" gorm:"column:LeaveTime;comment:'离职日期'"`
	EnterStep    int    `json:"enter_step" gorm:"column:EnterStep;comment:'入职进度,0待用户确认信息，1已确认信息，2待用户网签合同，3合同签完'"`
	LeaveStep    int    `json:"leave_step" gorm:"column:LeaveStep;comment:'离职进度'"`
}

// 创建SalaryUsers
func (un *SalaryUsers) CreateSalaryUsers() (err error) {
	err = qmsql.DEFAULTDB.Create(un).Error
	return err
}

// 删除SalaryUsers
func (un *SalaryUsers) DeleteSalaryUsers() (err error) {
	err = qmsql.DEFAULTDB.Delete(un).Error
	return err
}

// 更新SalaryUsers
func (un *SalaryUsers) UpdateSalaryUsers() (err error, reun SalaryUsers) {
	err = qmsql.DEFAULTDB.Save(un).Error
	return err, *un
}

// 根据ID查看单条SalaryUsers
func (un *SalaryUsers) FindById() (err error, reun SalaryUsers) {
	err = qmsql.DEFAULTDB.Where("id = ?", un.ID).First(&reun).Error
	return err, reun
}

// 根据ID查看单条SalaryUsers
func (un *SalaryUsers) FindByOpenid() (err error, reun SalaryUsers) {
	err = qmsql.DEFAULTDB.Where("openid = ?", un.Openid).First(&reun).Error
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
		var reSalaryUsersList []SalaryUsers
		err = db.Find(&reSalaryUsersList).Error
		return err, reSalaryUsersList, total
	}
}
