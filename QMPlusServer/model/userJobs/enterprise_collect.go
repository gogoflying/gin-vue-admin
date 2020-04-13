// 自动生成模板EnterpriseCollect
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type EnterpriseCollect struct {
	gorm.Model
	Openid       string `json:"openid" gorm:"column:openid"`
	EnterpriseId int    `json:"enterprise_id" gorm:"column:EnterpriseId"`
}

// 创建EnterpriseCollect
func (jc *EnterpriseCollect) CreateEnterpriseCollect() (err error) {
	err = qmsql.DEFAULTDB.Create(jc).Error
	return err
}

// 删除EnterpriseCollect
func (jc *EnterpriseCollect) DeleteEnterpriseCollect() (err error) {
	err = qmsql.DEFAULTDB.Unscoped().Where("openid = ? and EnterpriseId = ?", jc.Openid, jc.EnterpriseId).Delete(jc).Error
	return err
}

// 更新EnterpriseCollect
func (jc *EnterpriseCollect) UpdateEnterpriseCollect() (err error, rejc EnterpriseCollect) {
	err = qmsql.DEFAULTDB.Save(jc).Error
	return err, *jc
}

// 根据ID查看单条EnterpriseCollect
func (jc *EnterpriseCollect) FindById() (err error, rejc EnterpriseCollect) {
	err = qmsql.DEFAULTDB.Where("id = ?", jc.ID).First(&rejc).Error
	return err, rejc
}

// 分页获取EnterpriseCollect
func (jc *EnterpriseCollect) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(jc, info)
	if err != nil {
		return
	} else {
		var reEnterpriseCollectList []EnterpriseCollect
		err = db.Find(&reEnterpriseCollectList).Error
		return err, reEnterpriseCollectList, total
	}
}

// 分页获取JobCollect
func (jc *EnterpriseCollect) GetInfoListByOpenid(openid string, page, pageSize int) (err error, list interface{}) {
	offset := pageSize * (page - 1)
	var jcs []EnterpriseCollect
	db := qmsql.DEFAULTDB.Model(jc)
	err = db.Where("openid = ?", openid).Limit(pageSize).Offset(offset).Find(&jcs).Error
	if err != nil || len(jcs) == 0 {
		return nil, nil
	}
	var enpIds []int
	for _, j := range jcs {
		enpIds = append(enpIds, j.EnterpriseId)
	}
	var eps []EnterpriseInfo
	db.Model(new(EnterpriseInfo)).Where("id IN (?)", enpIds).Find(&eps)
	return nil, eps
}
