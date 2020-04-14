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
	UserName       string `json:"userName" gorm:"column:user_name;comment:'用户姓名'"`
	AvatarUrl      string `json:"avatarUrl" gorm:"column:avatarUrl;comment:'用户头像url，暂未使用'"`
	Genderindex    int    `json:"genderindex" gorm:"column:genderindex;comment:'性别，0男1女'"`
	EdulevelIndex  int    `json:"edulevelindex" gorm:"column:edulevelindex;comment:'最高学历id'"`
	WorksYearindex int    `json:"worksYearindex" gorm:"column:worksYearindex;comment:'工作年限id，关联job_work_expire'"`
	Cityindex      int    `json:"cityindex" gorm:"column:cityindex;comment:'城市id，关联citynames'"`
	Mobile         string `json:"contact" gorm:"column:mobile;comment:'手机号'"`
	Email          string `json:"email" gorm:"column:email;comment:'邮箱'"`
	Birthday       string `json:"birthday" gorm:"column:birthday;comment:'出生年月日'"`
	Position       string `json:"position" gorm:"column:position;comment:'未使用'"`
	Myself         string `json:"myself" gorm:"column:myself;comment:'未使用'"`
	Openid         string `json:"openid" gorm:"column:openid;comment:'唯一id'"`
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
