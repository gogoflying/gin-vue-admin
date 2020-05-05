// 自动生成模板Users
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type UserLoginInfo struct {
	Code   string `json:"code"`
	Openid string `json:"openid"`
	Mobile string `json:"mobile"`
}

type Users struct {
	gorm.Model
	Openid   string `json:"openid" gorm:"column:openid;comment:'用户标识码'"`
	Mobile   string `json:"mobile" gorm:"column:mobile;comment:'手机号'"`
	IsResume int    `json:"is_resume" gorm:"column:is_resume;comment:'是否有简历，1表示有，0没有'"`
	Status   int    `json:"status" gorm:"column:status;comment:'用户状态，1表示正常，其他非正常，如禁用之类的'"`
}

// 创建Users
func (users *Users) CreateUsers() (err error) {
	err = qmsql.DEFAULTDB.Create(users).Error
	return err
}

// 删除Users
func (users *Users) DeleteUsers() (err error) {
	err = qmsql.DEFAULTDB.Delete(users).Error
	return err
}

// 更新Users
func (users *Users) UpdateUsers() (err error, reusers Users) {
	err = qmsql.DEFAULTDB.Model(users).Where("openid = ? and status = 1", users.Openid).Update("is_resume", users.IsResume).Error
	return err, *users
}

// 根据ID查看单条Users
func (users *Users) FindById() (err error, reusers Users) {
	err = qmsql.DEFAULTDB.Where("openid = ?", users.Openid).First(&reusers).Error
	return err, reusers

}

// 根据ID查看单条Users
func (users *Users) FindByOpenId() (err error, reusers Users, edus, wks, base, dream interface{}) {
	err = qmsql.DEFAULTDB.Model(users).Where("openid = ? and status = 1", users.Openid).First(&reusers).Error
	if err != nil {
		return
	}
	if reusers.IsResume == 1 {
		_, edus, _ = new(UserEducation).GetInfoListByOpenid(users.Openid, 1, 20)
		_, wks, _ = new(UserWork).GetInfoListOpenId(users.Openid, 1, 20)
		_, base = (&UserBaseInfo{Openid: users.Openid}).FindByOpenid()
		_, dream = (&UserDream{Openid: users.Openid}).FindByOpenid()
	}
	return

}

// 分页获取Users
func (users *Users) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(users, info)
	if err != nil {
		return
	} else {
		var reUsersList []Users
		err = db.Find(&reUsersList).Error
		return err, reUsersList, total
	}
}
