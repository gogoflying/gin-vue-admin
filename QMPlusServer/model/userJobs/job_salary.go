// 自动生成模板JobSalary
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type JobSalary struct {
	gorm.Model
	Name string `json:"name" gorm:"column:name;comment:'工资范围名称，如10k-20k，则下面对应填10000,20000'"`
	Low  int    `json:"low" gorm:"column:low;comment:''"`
	High int    `json:"high" gorm:"column:high;comment:''"`
}

// 创建JobSalary
func (et *JobSalary) CreateJobSalary() (err error) {
	err = qmsql.DEFAULTDB.Create(et).Error
	return err
}

// 删除JobSalary
func (et *JobSalary) DeleteJobSalary() (err error) {
	err = qmsql.DEFAULTDB.Delete(et).Error
	return err
}

// 更新JobSalary
func (et *JobSalary) UpdateJobSalary() (err error, reet JobSalary) {
	err = qmsql.DEFAULTDB.Save(et).Error
	return err, *et
}

// 根据ID查看单条JobSalary
func (et *JobSalary) FindById() (err error, reet JobSalary) {
	err = qmsql.DEFAULTDB.Where("id = ?", et.ID).First(&reet).Error
	return err, reet
}

// 分页获取JobSalary
func (et *JobSalary) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(et, info)
	if err != nil {
		return
	} else {
		var reJobSalaryList []JobSalary
		err = db.Find(&reJobSalaryList).Error
		return err, reJobSalaryList, total
	}
}
