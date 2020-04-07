// 自动生成模板Usereducation
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type Usereducation struct {
	gorm.Model
	Graduation     string `json:"graduation" gorm:"column:graduation"`
	Schoolname     string `json:"schoolname" gorm:"column:schoolname"`
	Edulevelindex  int    `json:"edulevelindex" gorm:"column:edulevelindex"`
	Profession     string `json:"profession" gorm:"column:profession"`
	GraduationTime string `json:"graduationTime" gorm:"column:graduationTime"`
	Openid         string `json:"openid" gorm:"column:openid"`
}

// 创建Usereducation
func (ed *Usereducation) CreateUsereducation() (err error, list interface{}) {
	qmsql.DEFAULTDB.Create(ed)
	err, eds, _ := ed.GetInfoListByOpenid(ed.Openid, 1, 20)
	return err, eds
}

// 删除Usereducation
func (ed *Usereducation) DeleteUsereducation() (err error, list interface{}) {
	qmsql.DEFAULTDB.Delete(ed)
	err, eds, _ := ed.GetInfoListByOpenid(ed.Openid, 1, 20)
	return err, eds
}

// 更新Usereducation
func (ed *Usereducation) UpdateUsereducation() (err error, list interface{}) {
	qmsql.DEFAULTDB.Save(ed)
	err, eds, _ := ed.GetInfoListByOpenid(ed.Openid, 1, 20)
	return err, eds
}

// 根据ID查看单条Usereducation
func (ed *Usereducation) FindById() (err error, reed Usereducation) {
	err = qmsql.DEFAULTDB.Where("id = ?", ed.ID).First(&reed).Error
	return err, reed
}

// 分页获取Usereducation
func (ed *Usereducation) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(ed, info)
	if err != nil {
		return
	} else {
		var reUsereducationList []Usereducation
		err = db.Find(&reUsereducationList).Error
		return err, reUsereducationList, total
	}
}

// 分页获取Usereducation
func (ed *Usereducation) GetInfoListByOpenid(openId string, page, pageSize int) (err error, list interface{}, total int) {
	offset := pageSize * (page - 1)
	var reUserEduList []Usereducation
	db := qmsql.DEFAULTDB.Model(ed).Where("openid = ?", openId).Count(&total)
	err = db.Order("id desc").Offset(offset).Limit(pageSize).Find(&reUserEduList).Error
	if err != nil {
		return
	}
	return nil, reUserEduList, total
}
