// 自动生成模板Usereducation
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type UserEducation struct {
	gorm.Model
	Graduation     string   `json:"graduation" gorm:"column:graduation;comment:''"`
	Schoolname     string   `json:"schoolname" gorm:"column:schoolname;comment:''"`
	EdulevelIndex  int      `json:"edulevelindex" gorm:"column:edulevelindex;comment:''"`
	Edulevel       EduLevel `json:"edu_level" gorm:"ForeignKey:EdulevelIndex;AssociationForeignKey:ID"`
	Profession     string   `json:"profession" gorm:"column:profession;comment:''"`
	GraduationTime string   `json:"graduationTime" gorm:"column:graduationTime;comment:''"`
	Openid         string   `json:"openid" gorm:"column:openid"`
}

// 创建Usereducation
func (ed *UserEducation) CreateUsereducation() (err error, list interface{}) {
	qmsql.DEFAULTDB.Save(ed)
	err, eds, _ := ed.GetInfoListByOpenid(ed.Openid, 1, 20)
	return err, eds
}

// 删除Usereducation
func (ed *UserEducation) DeleteUsereducation() (err error, list interface{}) {
	qmsql.DEFAULTDB.Model(ed).Where("openid = ? and id = ?", ed.Openid, ed.ID).Delete(&UserEducation{})
	err, eds, _ := ed.GetInfoListByOpenid(ed.Openid, 1, 20)
	return err, eds
}

// 更新Usereducation
func (ed *UserEducation) UpdateUsereducation() (err error, list interface{}) {
	//qmsql.DEFAULTDB.Save(ed)
	//err, eds, _ := ed.GetInfoListByOpenid(ed.Openid, 1, 20)
	//return err, eds

	db := qmsql.DEFAULTDB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var tmp UserEducation
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&tmp, ed.ID).Error; err != nil {
		tx.Rollback()
		return err, nil
	}
	if err := tx.Save(ed).Error; err != nil {
		tx.Rollback()
		return err, nil
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err, nil
	}
	err, eds, _ := ed.GetInfoListByOpenid(ed.Openid, 1, 20)
	return err, eds
}

// 根据ID查看单条Usereducation
func (ed *UserEducation) FindById() (err error, reed UserEducation) {
	err = qmsql.DEFAULTDB.Where("id = ?", ed.ID).First(&reed).Error
	return err, reed
}

// 分页获取Usereducation
func (ed *UserEducation) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(ed, info)
	if err != nil {
		return
	} else {
		var reUsereducationList []UserEducation
		err = db.Find(&reUsereducationList).Error
		return err, reUsereducationList, total
	}
}

// 分页获取Usereducation
func (ed *UserEducation) GetInfoListByOpenid(openId string, page, pageSize int) (err error, list interface{}, total int) {
	offset := pageSize * (page - 1)
	var reUserEduList []UserEducation
	db := qmsql.DEFAULTDB.Model(ed).Where("openid = ?", openId).Count(&total)
	err = db.Order("id desc").Offset(offset).Limit(pageSize).Preload("Edulevel").Find(&reUserEduList).Error
	if err != nil {
		return
	}
	for _, re := range reUserEduList {
		if re.EdulevelIndex == 0 {
			var el EduLevel
			err = qmsql.DEFAULTDB.Where("id = ?", 0).First(&el).Error
			if err == nil {
				re.Edulevel = el
			}
		}
	}
	return nil, reUserEduList, total
}
