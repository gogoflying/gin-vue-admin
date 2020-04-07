// 自动生成模板UserDream
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type UserDream struct {
	gorm.Model
	Dreamposi     string `json:"dreamposi" gorm:"column:dreamposi"`
	WorkTypeindex int    `json:"workTypeindex" gorm:"column:workTypeindex"`
	Cityindex     int    `json:"cityindex" gorm:"column:cityindex"`
	Salaryindex   int    `json:"salaryindex" gorm:"column:salaryindex"`
	Openid        string `json:"openid" gorm:"column:openid"`
}

// 创建UserDream
func (dm *UserDream) CreateUserDream() (err error) {
	err = qmsql.DEFAULTDB.Create(dm).Error
	return err
}

// 删除UserDream
func (dm *UserDream) DeleteUserDream() (err error) {
	err = qmsql.DEFAULTDB.Delete(dm).Error
	return err
}

// 更新UserDream
func (dm *UserDream) UpdateUserDream() (err error, redm UserDream) {
	err = qmsql.DEFAULTDB.Save(dm).Error
	return err, *dm
}

// 根据ID查看单条UserDream
func (dm *UserDream) FindById() (err error, redm UserDream) {
	err = qmsql.DEFAULTDB.Where("id = ?", dm.ID).First(&redm).Error
	return err, redm
}

// 根据openid查看单条UserDream
func (dm *UserDream) FindByOpenid() (err error, redm UserDream) {
	err = qmsql.DEFAULTDB.Where("openid = ?", dm.Openid).First(&redm).Error
	return err, redm
}

// 分页获取UserDream
func (dm *UserDream) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(dm, info)
	if err != nil {
		return
	} else {
		var reUserDreamList []UserDream
		err = db.Find(&reUserDreamList).Error
		return err, reUserDreamList, total
	}
}
