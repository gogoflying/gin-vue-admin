// 自动生成模板JobWorkExpire
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/sysModel"
	"github.com/jinzhu/gorm"
)

type UserAuth struct {
	gorm.Model
	EnterPriseName string `json:"enterprise_name" gorm:"column:enterprise_name;comment:'企业名称'"`
	EnterPriseId   uint   `json:"enterprise_id" gorm:"column:enterprise_id;comment:'企业id'"`
	Username       string `json:"userName" gorm:"column:username;comment:'企业用户名'"`
	UserId         uint   `json:"user_id" gorm:"column:user_id;comment:'企业用户id'"`
	Status         int    `json:"status" gorm:"column:status;comment:'0、审核中1、审核通过'"`
	AuthorityName  string `json:"authorityName" gorm:"-"`
}

// 创建UserAuth
func (ua *UserAuth) CreateUserAuth() (err error) {
	err = qmsql.DEFAULTDB.Create(ua).Error
	return err
}

// 删除JobWorkExpire
func (ua *UserAuth) DeleteUserAuth() (err error) {
	err = qmsql.DEFAULTDB.Delete(ua).Error
	return err
}

// 更新UserAuth
func (ua *UserAuth) UpdateUserAuth() (err error, reuserauth UserAuth) {
	err = qmsql.DEFAULTDB.Save(ua).Error
	return err, *ua
}

// 根据ID查看单条UserAuth
func (ua *UserAuth) FindById() (err error, reuserauth UserAuth) {
	err = qmsql.DEFAULTDB.Where("id = ?", ua.ID).First(&reuserauth).Error
	return err, reuserauth
}

// 根据user_id查看单条UserAuth
func (ua *UserAuth) FindByUserId() (err error) {
	err = qmsql.DEFAULTDB.Where("user_id = ?", ua.UserId).First(ua).Error
	return err
}

// 分页获取JobWorkExpire
func (ua *UserAuth) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(ua, info)
	if err != nil {
		return
	} else {
		reUserauthList := make([]*UserAuth, 0)
		if ua.EnterPriseName != "" {
			db = db.Where("enterprise_name LIKE ?", "%"+ua.EnterPriseName+"%")
		}

		if ua.Username != "" {
			db = db.Where("userName LIKE ?", "%"+ua.Username+"%")
		}

		err = db.Find(&reUserauthList).Count(&total).Error
		if err != nil {
			return err, reUserauthList, total
		} else {
			err = db.Find(&reUserauthList).Error
		}
		if err == nil {
			for _, ru := range reUserauthList {
				err, rers := new(sysModel.SysUser).FindSysUserById(ru.UserId)
				if err != nil {
					return err, reUserauthList, total
				}
				ru.AuthorityName = rers.Authority.AuthorityName
			}
		}

		return err, reUserauthList, total
	}
}
