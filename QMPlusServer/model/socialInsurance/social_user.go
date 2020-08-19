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
	Mobile string `json:"mobile" gorm:"column:moblie;comment:'用户手机号'"`
	IdCard string `json:"id_card" gorm:"column:id_card;comment:'身份证号'"`
	Email  string `json:"email" gorm:"column:email;comment:'用户邮箱'"`
	Status int    `json:"status" gorm:"column:status;comment:'用户状态，1表示正常，其他非正常，如禁用之类的'"`
}

// 创建SocialUser
func (su *SocialUser) CreateSocialUser() (err error) {
	err = qmsql.DEFAULTDB.Create(su).Error
	return err
}

func (su *SocialUser) UpdateUsersMobile() (err error) {
	err = qmsql.DEFAULTDB.Model(su).Where("openid = ?", su.Openid).Update("mobile", su.Mobile).Error
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

func (su *SocialUser) UpdateSocialUserByOpenid() (err error) {
	return qmsql.DEFAULTDB.Model(su).Where("openid = ?", su.Openid).Updates(map[string]interface{}{"name": su.Name, "mobile": su.Mobile, "email": su.Email, "id_card": su.IdCard}).Error
}

// 根据ID查看单条SocialUser
func (su *SocialUser) FindById() (err error, reet SocialUser) {
	err = qmsql.DEFAULTDB.Where("id = ?", su.ID).First(&reet).Error
	return err, reet
}

func (su *SocialUser) FindByOpenid() (err error, reet SocialUser) {
	err = qmsql.DEFAULTDB.Where("openid = ?", su.Openid).First(&reet).Error
	return err, reet
}

// 分页获取SocialUser
func (su *SocialUser) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(su, info)
	if err != nil {
		return
	} else {
		var reSocialUserList []SocialUser
		model := qmsql.DEFAULTDB.Model(info)
		if su.Name != "" {
			model = model.Where("name LIKE ?", "%"+su.Name+"%")
			db = db.Where("name LIKE ?", "%"+su.Name+"%")
		}
		err = model.Find(&reSocialUserList).Count(&total).Error
		if err != nil {
			return err, reSocialUserList, total
		} else {
			err = db.Find(&reSocialUserList).Error
		}
		return err, reSocialUserList, total
	}
}
