// 自动生成模板SbConf
package userSalary

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
)

type SbConf struct {
      gorm.Model 
      City  string `json:"city" gorm:"column:city"`
      Price  float64 `json:"price" gorm:"column:price"`
      Attend  float64 `json:"attend" gorm:"column:attend"`
}

// 创建SbConf
func (un *SbConf)CreateSbConf()(err error){
        err = qmsql.DEFAULTDB.Create(un).Error
        return err
}

// 删除SbConf
func (un *SbConf)DeleteSbConf()(err error){
        err = qmsql.DEFAULTDB.Delete(un).Error
        return err
}

// 更新SbConf
func (un *SbConf)UpdateSbConf()(err error, reun SbConf){
        err = qmsql.DEFAULTDB.Save(un).Error
        return err, *un
}

// 根据ID查看单条SbConf
func (un *SbConf)FindById()(err error,reun SbConf){
    err = qmsql.DEFAULTDB.Where("id = ?",un.ID).First(&reun).Error
    return err,reun
}

// 分页获取SbConf
func (un *SbConf)GetInfoList(info modelInterface.PageInfo)(err error, list interface{}, total int){
    	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
    	err, db, total := servers.PagingServer(un, info)
    	if err != nil {
    		return
    	} else {
    		var reSbConfList []SbConf
    		err = db.Find(&reSbConfList).Error
    		return err, reSbConfList, total
    	}
}