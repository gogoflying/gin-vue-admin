// 自动生成模板SbUsers
package userSalary

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type SbUsers struct {
	gorm.Model
	Openid string `json:"openid" gorm:"column:openid"`
	Name   string `json:"name" gorm:"column:name;comment:'用户名'"`
	Card   string `json:"card" gorm:"column:card;comment:'身份证号'"`
	Mobile string `json:"mobile" gorm:"column:mobile;comment:'手机号'"`
}

// 创建SbUsers
func (un *SbUsers) CreateSbUsers() (err error) {
	err = qmsql.DEFAULTDB.Create(un).Error
	return err
}

// 删除SbUsers
func (un *SbUsers) DeleteSbUsers() (err error) {
	err = qmsql.DEFAULTDB.Delete(un).Error
	return err
}

// 更新SbUsers
func (un *SbUsers) UpdateSbUsers() (err error, reun SbUsers) {
	err = qmsql.DEFAULTDB.Save(un).Error
	return err, *un
}

// 根据ID查看单条SbUsers
func (un *SbUsers) FindById() (err error, reun SbUsers) {
	err = qmsql.DEFAULTDB.Where("id = ?", un.ID).First(&reun).Error
	return err, reun
}

func (un *SbUsers) FindByOpenid() (err error, reun SbUsers) {
	err = qmsql.DEFAULTDB.Where("openid = ?", un.Openid).First(&reun).Error
	return err, reun
}

// 分页获取SbUsers
func (un *SbUsers) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(un, info)
	if err != nil {
		return
	} else {
		var reSbUsersList []SbUsers
		err = db.Find(&reSbUsersList).Error
		return err, reSbUsersList, total
	}
}
