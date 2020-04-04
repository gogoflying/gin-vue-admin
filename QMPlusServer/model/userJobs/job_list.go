// 自动生成模板Joblist
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type Joblist struct {
	gorm.Model
	Jobname     string `json:"job_name" gorm:"column:job_name"`
	JobsalaHigh int    `json:"job_salary_high" gorm:"column:job_salary_high"`
	JobsalaLow  int    `json:"job_salary_low" gorm:"column:job_salary_low"`
	JobCity     string `json:"job_city" gorm:"column:job_city"`
	JobYears    string `json:"job_years" gorm:"column:job_years"`
	JobEdu      string `json:"job_edu" gorm:"column:job_edu"`
	JobType     string `json:"job_type" gorm:"column:job_type"`
	CompanyName string `json:"company_name" gorm:"column:company_name"`
	CompanyImg  string `json:"company_img" gorm:"column:company_img"`
	JobCityId   int    `json:"job_city_id" gorm:"column:job_city_id"`
	CompanyId   int    `json:"company_id" gorm:"column:company_id"`
}

// 创建Joblist
func (jl *Joblist) CreateJoblist() (err error) {
	err = qmsql.DEFAULTDB.Create(jl).Error
	return err
}

// 删除Joblist
func (jl *Joblist) DeleteJoblist() (err error) {
	err = qmsql.DEFAULTDB.Delete(jl).Error
	return err
}

// 更新Joblist
func (jl *Joblist) UpdateJoblist() (err error, rejl Joblist) {
	err = qmsql.DEFAULTDB.Save(jl).Error
	return err, *jl
}

// 根据ID查看单条Joblist
func (jl *Joblist) FindById() (err error, rejl Joblist) {
	err = qmsql.DEFAULTDB.Where("id = ?", jl.ID).First(&rejl).Error
	return err, rejl
}

// 分页获取Joblist
func (jl *Joblist) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(jl, info)
	if err != nil {
		return
	} else {
		var reJoblistList []Joblist
		err = db.Find(&reJoblistList).Error
		return err, reJoblistList, total
	}
}

// 分页获取Joblist
func (jl *Joblist) GetInfoListSearch(keyword string, low int, hight int, info modelInterface.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err = qmsql.DEFAULTDB.Model(jl).Count(&total).Error
	if err != nil {
		return
	}
	var reJoblistList []Joblist
	db := qmsql.DEFAULTDB.Model(jl)
	if len(keyword) > 0 {
		db = db.Where("job_name LIKE ? or company_name LIKE ?", keyword, keyword)
	}
	if low > 0 {
		db = db.Where("job_salary_low >= ?", low)
	}
	if hight > 0 {
		db = db.Where("job_salary_high <= ?", hight)
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&reJoblistList).Error
	err = db.Find(&reJoblistList).Error
	return err, reJoblistList, total
}
