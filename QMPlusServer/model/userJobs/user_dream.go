// 自动生成模板UserDream
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userCity"

	"github.com/jinzhu/gorm"
)

type JobWorkType struct {
	gorm.Model
	Name string `json:"name" gorm:"column:name";comment:''`
}

type UserDream struct {
	gorm.Model
	Dreamposi     string            `json:"dreamposi" gorm:"column:dreamposi;comment:'期望职位名称'"`
	WorkTypeindex int               `json:"workTypeindex" gorm:"column:workTypeindex;comment:'期望工作类型id，0 全职1 兼职2 实习3 志愿者'"`
	Cityindex     int               `json:"cityindex" gorm:"column:cityindex;comment:'期望工作城市id'"`
	Salaryindex   int               `json:"salaryindex" gorm:"column:salaryindex;comment:'期望薪资范围id，关联job_salary表'"`
	DutyTimeindex int               `json:"dutyTimeindex" gorm:"column:dutyTimeindex;comment:'到岗时间id，关联job_duty_time表'"`
	Openid        string            `json:"openid" gorm:"column:openid;comment:'用户唯一标识'"`
	Jobdutytime   JobDutyTime       `json:"job_duty_time" gorm:"ForeignKey:DutyTimeindex;AssociationForeignKey:ID"`
	Jobsalary     JobSalary         `json:"job_salary" gorm:"ForeignKey:Salaryindex;AssociationForeignKey:ID"`
	JobworkType   JobWorkType       `json:"job_work_type" gorm:"ForeignKey:WorkTypeindex;AssociationForeignKey:ID"`
	Cityname      userCity.Cityname `json:"city_name" gorm:"ForeignKey:Cityindex;AssociationForeignKey:ID"`
}

// 创建UserDream
func (dm *UserDream) CreateUserDream() (err error) {
	_, ndm := dm.FindByOpenid()
	if ndm.ID > 0 {
		dm.ID = ndm.ID
		err = qmsql.DEFAULTDB.Model(dm).Update(dm).Error
		return
	}
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
	err = qmsql.DEFAULTDB.Model(dm).Preload("Jobsalary").Preload("Jobdutytime").Preload("JobworkType").Preload("Cityname").Where("openid = ?", dm.Openid).First(&redm).Error
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
