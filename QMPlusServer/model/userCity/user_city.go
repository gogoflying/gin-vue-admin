// 自动生成模板Cityname
package userCity

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type Cityname struct {
	gorm.Model
	Name string `json:"name" gorm:"column:name"`
	Type int    `json:"type" gorm:"column:type"`
}

// 创建Cityname
func (cityname *Cityname) CreateCityname() (err error) {
	err = qmsql.DEFAULTDB.Create(cityname).Error
	return err
}

// 删除Cityname
func (cityname *Cityname) DeleteCityname() (err error) {
	err = qmsql.DEFAULTDB.Delete(cityname).Error
	return err
}

// 更新Cityname
func (cityname *Cityname) UpdateCityname() (err error, recityname Cityname) {
	err = qmsql.DEFAULTDB.Save(cityname).Error
	return err, *cityname
}

// 根据ID查看单条Cityname
func (cityname *Cityname) FindById() (err error, recityname Cityname) {
	err = qmsql.DEFAULTDB.Where("id = ?", cityname.ID).First(&recityname).Error
	return err, recityname
}

// 分页获取Cityname
func (cityname *Cityname) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServerAsc(cityname, info)
	if err != nil {
		return
	} else {
		var reCitynameList []Cityname
		err = db.Find(&reCitynameList).Error
		return err, reCitynameList, total
	}
}
