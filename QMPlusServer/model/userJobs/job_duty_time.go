// 自动生成模板JobDutyTime
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type JobDutyTime struct {
	gorm.Model
	Name string `json:"name" gorm:"column:name;comment:'入职时间，如1个月内，随时到岗等'"`
}

// 创建JobDutyTime
func (et *JobDutyTime) CreateJobDutyTime() (err error) {
	err = qmsql.DEFAULTDB.Create(et).Error
	return err
}

// 删除JobDutyTime
func (et *JobDutyTime) DeleteJobDutyTime() (err error) {
	err = qmsql.DEFAULTDB.Delete(et).Error
	return err
}

// 更新JobDutyTime
func (et *JobDutyTime) UpdateJobDutyTime() (err error, reet JobDutyTime) {
	err = qmsql.DEFAULTDB.Save(et).Error
	return err, *et
}

// 根据ID查看单条JobDutyTime
func (et *JobDutyTime) FindById() (err error, reet JobDutyTime) {
	err = qmsql.DEFAULTDB.Where("id = ?", et.ID).First(&reet).Error
	return err, reet
}

// 分页获取JobDutyTime
func (et *JobDutyTime) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServerAsc(et, info)
	if err != nil {
		return
	} else {
		var reJobDutyTimeList []JobDutyTime
		err = db.Find(&reJobDutyTimeList).Error
		return err, reJobDutyTimeList, total
	}
}
