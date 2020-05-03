// 自动生成模板SbOrders
package userSalary

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
)

type SbOrders struct {
      gorm.Model 
      City  string `json:"city" gorm:"column:city"`
      Name  string `json:"name" gorm:"column:name"`
      Uid  int `json:"uid" gorm:"column:uid"`
      Begin  string `json:"begin" gorm:"column:begin"`
      Money  string `json:"money" gorm:"column:money"`
      Count  int `json:"count" gorm:"column:count"`
      Status  int `json:"status" gorm:"column:status"`
}

// 创建SbOrders
func (un *SbOrders)CreateSbOrders()(err error){
        err = qmsql.DEFAULTDB.Create(un).Error
        return err
}

// 删除SbOrders
func (un *SbOrders)DeleteSbOrders()(err error){
        err = qmsql.DEFAULTDB.Delete(un).Error
        return err
}

// 更新SbOrders
func (un *SbOrders)UpdateSbOrders()(err error, reun SbOrders){
        err = qmsql.DEFAULTDB.Save(un).Error
        return err, *un
}

// 根据ID查看单条SbOrders
func (un *SbOrders)FindById()(err error,reun SbOrders){
    err = qmsql.DEFAULTDB.Where("id = ?",un.ID).First(&reun).Error
    return err,reun
}

// 分页获取SbOrders
func (un *SbOrders)GetInfoList(info modelInterface.PageInfo)(err error, list interface{}, total int){
    	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
    	err, db, total := servers.PagingServer(un, info)
    	if err != nil {
    		return
    	} else {
    		var reSbOrdersList []SbOrders
    		err = db.Find(&reSbOrdersList).Error
    		return err, reSbOrdersList, total
    	}
}