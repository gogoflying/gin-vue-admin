package socialInsurance

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
)

type SocialUser struct {
	gorm.Model
	Name   string `json:"name" gorm:"column:name;comment:'姓名'"`
	Openid string `json:"openid" gorm:"column:openid;comment:'用户openid'"`
	IdCard string `json:"id_card" gorm:"column:id_card;comment:'身份证号'"`
}

// 创建SocialUser
func (su *SocialUser) CreateSocialUser() (err error) {
	err = qmsql.DEFAULTDB.Create(su).Error
	return err
}

// 删除SocialUser
func (su *SocialUser) DeleteSocialUser() (err error) {
	err = qmsql.DEFAULTDB.Delete(su).Error
	return err
}

// 更新SocialUser
func (su *SocialUser) UpdateSocialUser() (err error, reet SocialUser) {
	err = qmsql.DEFAULTDB.Save(su).Error
	return err, *su
}

// 根据ID查看单条SocialUser
func (su *SocialUser) FindById() (err error, reet SocialUser) {
	err = qmsql.DEFAULTDB.Where("id = ?", su.ID).First(&reet).Error
	return err, reet
}

// 分页获取SocialUser
func (su *SocialUser) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServerAsc(su, info)
	if err != nil {
		return
	} else {
		var reSocialUserList []SocialUser
		err = db.Find(&reSocialUserList).Error
		return err, reSocialUserList, total
	}
}
