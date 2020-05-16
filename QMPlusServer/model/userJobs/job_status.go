// 自动生成模板ResumeStatus
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type ResumeStatus struct {
	gorm.Model
	Openid       string    `json:"openid" gorm:"column:open_id;comment:'用户识别码'"`
	Jobid        int       `json:"resume_id" gorm:"column:job_id;comment:'职位id'"`
	ResumeStatus int       `json:"resume_status" gorm:"column:resume_status;comment:'简历状态，0未读，1已读，2有意向，3约面试，4不合适'"`
	Results      []Joblist `json:"result" gorm:"-"`
	Jobname      string    `json:"p_name" gorm:"column:job_name;comment:'职位名称'"`
	CompanyName  string    `json:"enterprise_name" gorm:"column:company_name;comment:'公司名称'"`
	CompanyId    int       `json:"enterprise_id" gorm:"column:company_id;comment:'公司id，关联enterprise_infos表'"`
	JobsalaLow   int       `json:"p_salary_low" gorm:"-"`
	JobCityId    int       `json:"p_city_id" gorm:"-"`
	JobYearsId   int       `json:"p_edujy_id" gorm:"-"`
	JobEduId     int       `json:"p_education_id" gorm:"-"`
	JobTypeId    int       `json:"p_type_id" gorm:"-"`
	JobInfo      Joblist   `json:"job_info" gorm:"ForeignKey:Jobid;AssociationForeignKey:ID"`
	UserInfo     Users     `json:"user_info" gorm:"ForeignKey:Openid;AssociationForeignKey:Openid"`
}

// 创建ResumeStatus
func (rs *ResumeStatus) CreateResumeStatus() (list interface{}, err error) {
	err = qmsql.DEFAULTDB.Create(rs).Error
	if err != nil {
		return
	}
	var rss []ResumeStatus
	qmsql.DEFAULTDB.Model(rs).Where("open_id = ?", rs.Openid).Find(&rss)
	var ids []int
	for _, r := range rss {
		ids = append(ids, r.Jobid)
	}
	_, list = (new(Joblist)).GetInfoListSearchSimilar(ids, rs.Jobname, rs.JobYearsId, rs.JobEduId, rs.JobTypeId, rs.JobCityId, rs.JobsalaLow, 1, 5)
	return
}

// 删除ResumeStatus
func (rs *ResumeStatus) DeleteResumeStatus() (err error) {
	err = qmsql.DEFAULTDB.Delete(rs).Error
	return err
}

// 更新ResumeStatus
func (rs *ResumeStatus) UpdateResumeStatus() (err error, rers ResumeStatus) {
	err = qmsql.DEFAULTDB.Save(rs).Error
	return err, *rs
}

// 根据ID查看单条ResumeStatus
func (rs *ResumeStatus) FindById() (err error, rers ResumeStatus) {
	err = qmsql.DEFAULTDB.Where("id = ?", rs.ID).First(&rers).Error
	return err, rers
}

// 分页获取ResumeStatus
func (rs *ResumeStatus) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(rs, info)
	if err != nil {
		return
	} else {
		var reResumeStatusList []ResumeStatus
		model := qmsql.DEFAULTDB.Model(info)
		if rs.CompanyId != 0 {
			model = model.Where("company_id = ?", rs.CompanyId)
			db = db.Where("company_id = ?", rs.CompanyId)
		}
		if rs.CompanyName != "" {
			model = model.Where("company_name = ?", rs.CompanyName)
			db = db.Where("company_name = ?", rs.CompanyName)
		}
		if rs.Jobname != "" {
			model = model.Where("job_name = ?", rs.Jobname)
			db = db.Where("job_name = ?", rs.Jobname)
		}
		err = model.Preload("JobInfo").Preload("UserInfo").Find(&reResumeStatusList).Count(&total).Error
		if err != nil {
			return err, reResumeStatusList, total
		} else {
			err = db.Preload("JobInfo").Preload("UserInfo").Find(&reResumeStatusList).Error
		}
		return err, reResumeStatusList, total
	}
}

// 分页获取ResumeStatus
func (rs *ResumeStatus) GetResumeStatus(openId string, status, page, limit int) (err error, rers ResumeStatus, total int) {
	offset := limit * (page - 1)
	err = qmsql.DEFAULTDB.Model(rs).Count(&total).Error
	if err != nil {
		return
	}
	var rss []ResumeStatus
	db := qmsql.DEFAULTDB.Model(rs).Where("open_id = ?", openId).Where("resume_status = ?", status)
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&rss).Error
	if err != nil {
		return
	}
	for _, v := range rss {
		jb := new(Joblist)
		err = qmsql.DEFAULTDB.Model(new(Joblist)).Where("id = ?", v.Jobid).Find(jb).Error
		if err != nil {
			continue
		}
		rs.Results = append(rs.Results, *jb)
	}
	return err, *rs, total
}
