// 自动生成模板EnterpriseIndust
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type EnterpriseIndust struct {
	gorm.Model
	Name string `json:"name" gorm:"column:name;comment:'企业所属行业名称，互联网、医疗等'"`
}

// 创建EnterpriseIndust
func (et *EnterpriseIndust) CreateEnterpriseIndust() (err error) {
	err = qmsql.DEFAULTDB.Create(et).Error
	return err
}

// 删除EnterpriseIndust
func (et *EnterpriseIndust) DeleteEnterpriseIndust() (err error) {
	err = qmsql.DEFAULTDB.Delete(et).Error
	return err
}

// 更新EnterpriseIndust
func (et *EnterpriseIndust) UpdateEnterpriseIndust() (err error, reet EnterpriseIndust) {
	err = qmsql.DEFAULTDB.Save(et).Error
	return err, *et
}

// 根据ID查看单条EnterpriseIndust
func (et *EnterpriseIndust) FindById() (err error, reet EnterpriseIndust) {
	err = qmsql.DEFAULTDB.Where("id = ?", et.ID).First(&reet).Error
	return err, reet
}

// 分页获取EnterpriseIndust
func (et *EnterpriseIndust) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(et, info)
	if err != nil {
		return
	} else {
		var reEnterpriseIndustList []EnterpriseIndust
		err = db.Find(&reEnterpriseIndustList).Error
		return err, reEnterpriseIndustList, total
	}
}
