// 自动生成模板Salary_contract
package userSalary

import (
	//"time"
	//"gin-vue-admin/global"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
	//"github.com/pkg/errors"
)

type UserContractInfo struct {
	OpenId          string `json:"openid"`
	TmpContractPath string `json:"filepath"`
}

type SalaryContract struct {
	gorm.Model
	Openid                    string `json:"openid" gorm:"column:openid"`
	Enter_contract_source_url string `json:"enter_contract_source_url" gorm:"column:enter_contract_source_url"`
	Leave_contract_source_url string `json:"leave_contract_source_url" gorm:"column:leave_contract_source_url"`
	Enter_contract_target_url string `json:"enter_contract_target_url" gorm:"column:enter_contract_target_url"`
	Leave_contract_target_url string `json:"leave_contract_target_url" gorm:"column:leave_contract_target_url"`
	Status                    int    `json:"status" gorm:"column:status"`
	Version                   string `json:"version" gorm:"column:version"`
	//Created_at  time.Time `json:"created_at" gorm:"column:created_at"`
	//Updated_at  time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// 创建SalaryContract
func (un *SalaryContract) CreateSalaryContract() (err error) {
	err = qmsql.DEFAULTDB.Create(un).Error
	return err
}

// 删除SalaryContract
func (un *SalaryContract) DeleteSalaryContract() (err error) {
	err = qmsql.DEFAULTDB.Delete(un).Error
	return err
}

// 更新SalaryContract
func (un *SalaryContract) UpdateSalaryContract() (err error, reun SalaryContract) {
	//err = qmsql.DEFAULTDB.Save(un).Error
	err = qmsql.DEFAULTDB.Model(un).Where("openid = ? ", un.Openid).Updates(un).Error
	return err, *un
}

// 根据ID查看单条SalaryContract
func (un *SalaryContract) FindById() (err error, reun SalaryContract) {
	err = qmsql.DEFAULTDB.Where("openid = ?", un.Openid).First(&reun).Error
	//err = qmsql.DEFAULTDB.Where("id = ?",un.ID).First(&reun).Error
	return err, reun
}

// 分页获取SalaryContract
func (un *SalaryContract) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := qmsql.DEFAULTDB
	if err != nil {
		return
	} else {
		var reSalaryContractList []SalaryContract
		err = db.Limit(limit).Offset(offset).Find(&reSalaryContractList).Error
		return err, reSalaryContractList, total
	}
}
