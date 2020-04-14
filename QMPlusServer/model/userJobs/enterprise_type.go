// 自动生成模板EnterpriseType
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type EnterpriseType struct {
	gorm.Model
	Name string `json:"name" gorm:"column:name;comment:'企业类型名称，合资、外企、民营等'"`
}

// 创建EnterpriseType
func (et *EnterpriseType) CreateEnterpriseType() (err error) {
	err = qmsql.DEFAULTDB.Create(et).Error
	return err
}

// 删除EnterpriseType
func (et *EnterpriseType) DeleteEnterpriseType() (err error) {
	err = qmsql.DEFAULTDB.Delete(et).Error
	return err
}

// 更新EnterpriseType
func (et *EnterpriseType) UpdateEnterpriseType() (err error, reet EnterpriseType) {
	err = qmsql.DEFAULTDB.Save(et).Error
	return err, *et
}

// 根据ID查看单条EnterpriseType
func (et *EnterpriseType) FindById() (err error, reet EnterpriseType) {
	err = qmsql.DEFAULTDB.Where("id = ?", et.ID).First(&reet).Error
	return err, reet
}

// 分页获取EnterpriseType
func (et *EnterpriseType) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(et, info)
	if err != nil {
		return
	} else {
		var reEnterpriseTypeList []EnterpriseType
		err = db.Find(&reEnterpriseTypeList).Error
		return err, reEnterpriseTypeList, total
	}
}
