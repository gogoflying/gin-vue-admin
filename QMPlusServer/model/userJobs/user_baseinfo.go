// 自动生成模板UserBaseInfo
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type UserBaseInfo struct {
	gorm.Model
	UserName       string `json:"userName" gorm:"column:user_name"`
	AvatarUrl      string `json:"avatarUrl" gorm:"column:avatarUrl"`
	Genderindex    int    `json:"genderindex" gorm:"column:genderindex"`
	EdulevelIndex  int    `json:"edulevelindex" gorm:"column:edulevelindex"`
	WorksYearindex int    `json:"worksYearindex" gorm:"column:worksYearindex"`
	Cityindex      int    `json:"cityindex" gorm:"column:cityindex"`
	Mobile         string `json:"contact" gorm:"column:mobile"`
	Email          string `json:"email" gorm:"column:email"`
	Birthday       string `json:"birthday" gorm:"column:birthday"`
	Position       string `json:"position" gorm:"column:position"`
	Myself         string `json:"myself" gorm:"column:myself"`
	Openid         string `json:"openid" gorm:"column:openid"`
}

// 创建UserBaseInfo
func (bf *UserBaseInfo) CreateUserBaseInfo() (err error) {
	var rebf UserBaseInfo
	qmsql.DEFAULTDB.Where("openid = ?", bf.Openid).First(&rebf)
	if rebf.ID > 0 {
		bf.ID = rebf.ID
		err = qmsql.DEFAULTDB.Save(bf).Error
		return err
	}
	err = qmsql.DEFAULTDB.Create(bf).Error
	return err
}

// 删除UserBaseInfo
func (bf *UserBaseInfo) DeleteUserBaseInfo() (err error) {
	err = qmsql.DEFAULTDB.Delete(bf).Error
	return err
}

// 更新UserBaseInfo
func (bf *UserBaseInfo) UpdateUserBaseInfo() (err error, rebf UserBaseInfo) {
	err = qmsql.DEFAULTDB.Save(bf).Error
	return err, *bf
}

// 根据ID查看单条UserBaseInfo
func (bf *UserBaseInfo) FindById() (err error, rebf UserBaseInfo) {
	err = qmsql.DEFAULTDB.Where("id = ?", bf.ID).First(&rebf).Error
	return err, rebf
}

// 根据openid查看单条UserBaseInfo
func (bf *UserBaseInfo) FindByOpenid() (err error, rebf UserBaseInfo) {
	err = qmsql.DEFAULTDB.Model(bf).Where("openid = ?", bf.Openid).First(&rebf).Error
	return err, rebf
}

// 分页获取UserBaseInfo
func (bf *UserBaseInfo) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(bf, info)
	if err != nil {
		return
	} else {
		var reUserBaseInfoList []UserBaseInfo
		err = db.Find(&reUserBaseInfoList).Error
		return err, reUserBaseInfoList, total
	}
}
