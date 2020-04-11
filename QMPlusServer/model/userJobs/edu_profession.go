// 自动生成模板EduProfession
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
)

type EduProfession struct {
      gorm.Model 
      Name  string `json:"name" gorm:"column:name"`
}

// 创建EduProfession
func (et *EduProfession)CreateEduProfession()(err error){
        err = qmsql.DEFAULTDB.Create(et).Error
        return err
}

// 删除EduProfession
func (et *EduProfession)DeleteEduProfession()(err error){
        err = qmsql.DEFAULTDB.Delete(et).Error
        return err
}

// 更新EduProfession
func (et *EduProfession)UpdateEduProfession()(err error, reet EduProfession){
        err = qmsql.DEFAULTDB.Save(et).Error
        return err, *et
}

// 根据ID查看单条EduProfession
func (et *EduProfession)FindById()(err error,reet EduProfession){
    err = qmsql.DEFAULTDB.Where("id = ?",et.ID).First(&reet).Error
    return err,reet
}

// 分页获取EduProfession
func (et *EduProfession)GetInfoList(info modelInterface.PageInfo)(err error, list interface{}, total int){
    	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
    	err, db, total := servers.PagingServer(et, info)
    	if err != nil {
    		return
    	} else {
    		var reEduProfessionList []EduProfession
    		err = db.Find(&reEduProfessionList).Error
    		return err, reEduProfessionList, total
    	}
}