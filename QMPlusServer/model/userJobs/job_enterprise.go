// 自动生成模板EnterpriseInfo
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type EnterpriseInfo struct {
	gorm.Model
	EnterPriseName    string    `json:"enterprise_name" gorm:"column:enterprise_name"`
	EnterpriseLogo    string    `json:"enterprise_logo" gorm:"column:enterprise_logo"`
	EnterpriseImg     string    `json:"enterprise_img" gorm:"column:enterprise_img"`
	EnterpriseAddress string    `json:"enterprise_address" gorm:"column:enterprise_address"`
	EnterpriseDesc    string    `json:"enterprise_desc" gorm:"column:enterprise_desc"`
	EnterpriseCityId  int       `json:"city_id" gorm:"column:city_id"`
	Results           []Joblist `json:"result" gorm:"-"`
}

// 创建EnterpriseInfo
func (info *EnterpriseInfo) CreateEnterpriseInfo() (err error) {
	err = qmsql.DEFAULTDB.Create(info).Error
	return err
}

// 删除EnterpriseInfo
func (info *EnterpriseInfo) DeleteEnterpriseInfo() (err error) {
	err = qmsql.DEFAULTDB.Delete(info).Error
	return err
}

// 更新EnterpriseInfo
func (info *EnterpriseInfo) UpdateEnterpriseInfo() (err error, reinfo EnterpriseInfo) {
	err = qmsql.DEFAULTDB.Save(info).Error
	return err, *info
}

// 根据ID查看单条EnterpriseInfo
func (info *EnterpriseInfo) FindById() (err error, reinfo EnterpriseInfo) {
	err = qmsql.DEFAULTDB.Where("id = ?", info.ID).First(&reinfo).Error

	return err, reinfo
}

// 根据ID查看单条EnterpriseInfo
func (info *EnterpriseInfo) FindPositionsByCompId(page, pageSize int) (err error, reinfo EnterpriseInfo) {
	err = qmsql.DEFAULTDB.Where("id = ?", info.ID).First(&reinfo).Error
	if err != nil {
		return
	}
	limit := pageSize
	offset := pageSize * (page - 1)
	var jobList []Joblist
	err = qmsql.DEFAULTDB.Model(new(Joblist)).Where("company_id = ?", info.ID).Limit(limit).Offset(offset).Order("id desc").Find(&jobList).Error
	if err != nil {
		return
	}
	reinfo.Results = append(reinfo.Results, jobList...)
	return err, reinfo
}

// 分页获取EnterpriseInfo
func (info *EnterpriseInfo) GetInfoList(pinfo modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(info, pinfo)
	if err != nil {
		return
	} else {
		var reEnterpriseInfoList []EnterpriseInfo
		err = db.Find(&reEnterpriseInfoList).Error
		return err, reEnterpriseInfoList, total
	}
}

// 分页获取Joblist
func (info *EnterpriseInfo) GetInfoListSearch(keyword string, cityId, page, pageSize int) (err error, list interface{}, total int) {
	limit := pageSize
	offset := pageSize * (page - 1)
	err = qmsql.DEFAULTDB.Model(info).Count(&total).Error
	if err != nil {
		return
	}
	var infoList []EnterpriseInfo
	db := qmsql.DEFAULTDB.Model(info)
	if cityId > 0 {
		db = db.Where("city_id = ?", cityId)
	}
	if len(keyword) > 0 {
		db = db.Where("enterprise_name LIKE ? ", "%"+keyword+"%")
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&infoList).Error
	return err, infoList, total
}
