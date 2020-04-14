// 自动生成模板JobWorkExpire
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type JobWorkExpire struct {
	gorm.Model
	Name string `json:"name" gorm:"column:name;comment:'工作资历名称，如 3-5年，10年以上'"`
}

// 创建JobWorkExpire
func (et *JobWorkExpire) CreateJobWorkExpire() (err error) {
	err = qmsql.DEFAULTDB.Create(et).Error
	return err
}

// 删除JobWorkExpire
func (et *JobWorkExpire) DeleteJobWorkExpire() (err error) {
	err = qmsql.DEFAULTDB.Delete(et).Error
	return err
}

// 更新JobWorkExpire
func (et *JobWorkExpire) UpdateJobWorkExpire() (err error, reet JobWorkExpire) {
	err = qmsql.DEFAULTDB.Save(et).Error
	return err, *et
}

// 根据ID查看单条JobWorkExpire
func (et *JobWorkExpire) FindById() (err error, reet JobWorkExpire) {
	err = qmsql.DEFAULTDB.Where("id = ?", et.ID).First(&reet).Error
	return err, reet
}

// 分页获取JobWorkExpire
func (et *JobWorkExpire) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(et, info)
	if err != nil {
		return
	} else {
		var reJobWorkExpireList []JobWorkExpire
		err = db.Find(&reJobWorkExpireList).Error
		return err, reJobWorkExpireList, total
	}
}
