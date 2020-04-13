// 自动生成模板JobCollect
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type JobCollect struct {
	gorm.Model
	Openid string `json:"openid" gorm:"column:openid"`
	JobId  int    `json:"job_id" gorm:"column:job_id"`
}

// 创建JobCollect
func (jc *JobCollect) CreateJobCollect() (err error) {
	err = qmsql.DEFAULTDB.Create(jc).Error
	return err
}

// 删除JobCollect
func (jc *JobCollect) DeleteJobCollect() (err error) {
	err = qmsql.DEFAULTDB.Unscoped().Where("openid = ? and job_id = ?", jc.Openid, jc.JobId).Delete(jc).Error
	return err
}

// 更新JobCollect
func (jc *JobCollect) UpdateJobCollect() (err error, rejc JobCollect) {
	err = qmsql.DEFAULTDB.Save(jc).Error
	return err, *jc
}

// 根据ID查看单条JobCollect
func (jc *JobCollect) FindById() (err error, rejc JobCollect) {
	err = qmsql.DEFAULTDB.Where("id = ?", jc.ID).First(&rejc).Error
	return err, rejc
}

// 分页获取JobCollect
func (jc *JobCollect) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(jc, info)
	if err != nil {
		return
	} else {
		var reJobCollectList []JobCollect
		err = db.Find(&reJobCollectList).Error
		return err, reJobCollectList, total
	}
}

// 分页获取JobCollect
func (jc *JobCollect) GetInfoListByOpenid(openid string, page, pageSize int) (err error, list interface{}) {
	offset := pageSize * (page - 1)
	var jcs []JobCollect
	db := qmsql.DEFAULTDB.Model(jc)
	err = db.Where("openid = ?", openid).Limit(pageSize).Offset(offset).Find(&jcs).Error
	if err != nil || len(jcs) == 0 {
		return nil, nil
	}
	var jobIds []int
	for _, j := range jcs {
		jobIds = append(jobIds, j.JobId)
	}
	var jbs []Joblist
	db.Model(new(Joblist)).Where("id IN (?)", jobIds).Find(&jbs)
	return nil, jbs
}
