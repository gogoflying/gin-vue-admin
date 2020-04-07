// 自动生成模板UserWork
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type UserWork struct {
	gorm.Model
	Join        string `json:"join" gorm:"column:join"`
	Leave       string `json:"leave" gorm:"column:leave"`
	Companyname string `json:"companyname" gorm:"column:companyname"`
	Department  string `json:"department" gorm:"column:department"`
	WorkContent string `json:"workContent" gorm:"column:workContent"`
	Openid      string `json:"openid" gorm:"column:openid"`
}

// 创建UserWork
func (wk *UserWork) CreateUserWork() (err error, rewk interface{}) {
	err = qmsql.DEFAULTDB.Save(wk).Error
	if err != nil {
		return err, nil
	}
	err, list, _ := wk.GetInfoListOpenId(wk.Openid, 1, 20)
	return err, list
}

// 删除UserWork
func (wk *UserWork) DeleteUserWork() (err error) {
	err = qmsql.DEFAULTDB.Delete(wk).Error
	return err
}

// 更新UserWork
func (wk *UserWork) UpdateUserWork() (err error, rewk UserWork) {
	err = qmsql.DEFAULTDB.Save(wk).Error
	return err, *wk
}

// 根据ID查看单条UserWork
func (wk *UserWork) FindById() (err error, rewk UserWork) {
	err = qmsql.DEFAULTDB.Where("id = ?", wk.ID).First(&rewk).Error
	return err, rewk
}

// 分页获取UserWork
func (wk *UserWork) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(wk, info)
	if err != nil {
		return
	} else {
		var reUserWorkList []UserWork
		err = db.Find(&reUserWorkList).Error
		return err, reUserWorkList, total
	}
}

// 分页获取UserWork
func (wk *UserWork) GetInfoListOpenId(openId string, page, pageSize int) (err error, list interface{}, total int) {
	offset := pageSize * (page - 1)
	var reUserWorkList []UserWork
	db := qmsql.DEFAULTDB.Model(wk).Where("openid = ?", openId).Count(&total)
	err = db.Order("id desc").Offset(offset).Limit(pageSize).Find(&reUserWorkList).Error
	if err != nil {
		return
	}
	return nil, reUserWorkList, total
}
