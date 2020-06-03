// 自动生成模板ResumeStatus
package userJobs

import (
	"errors"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userSalary"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ResumeStatus struct {
	gorm.Model
	Openid            string       `json:"openid" gorm:"column:open_id;comment:'用户识别码'"`
	Jobid             int          `json:"resume_id" gorm:"column:job_id;comment:'职位id'"`
	ResumeStatus      int          `json:"resume_status" gorm:"column:resume_status;comment:'简历状态，0未读，1已读，2有意向，3约面试，4不合适'"`
	Results           []Joblist    `json:"result" gorm:"-"`
	Jobname           string       `json:"p_name" gorm:"column:job_name;comment:'职位名称'"`
	CompanyName       string       `json:"enterprise_name" gorm:"column:company_name;comment:'公司名称'"`
	CompanyId         int          `json:"enterprise_id" gorm:"column:company_id;comment:'公司id，关联enterprise_infos表'"`
	JobsalaLow        int          `json:"p_salary_low" gorm:"-"`
	JobCityId         int          `json:"p_city_id" gorm:"-"`
	JobYearsId        int          `json:"p_edujy_id" gorm:"-"`
	JobEduId          int          `json:"p_education_id" gorm:"-"`
	JobTypeId         int          `json:"p_type_id" gorm:"-"`
	JobInfo           Joblist      `json:"job_info" gorm:"ForeignKey:Jobid;AssociationForeignKey:ID;save_associations:false"`
	UserInfo          UserBaseInfo `json:"user_info" gorm:"ForeignKey:Openid;AssociationForeignKey:Openid;save_associations:false"`
	Memo              string       `json:"p_memo" gorm:"column:memo;comment:'memo'"`
	Remark            string       `json:"p_remark" gorm:"column:remark;comment:'给投递人的备注'"`
	Mobile            string       `json:"mobile" gorm:"-"`
	UserName          string       `json:"userName" gorm:"-"`
	ResumeStatuSearch *int         `json:"resume_status_search" gorm:"-"`
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

func (rs *ResumeStatus) UpdateMemo() error {
	return qmsql.DEFAULTDB.Model(rs).Update("memo", rs.Memo).Error
}

func (rs *ResumeStatus) UpdateRemark() error {
	return qmsql.DEFAULTDB.Model(rs).Update("remark", rs.Remark).Error
}

// 更新ResumeStatus
func (rs *ResumeStatus) UpdateResumeStatus() (err error, rers ResumeStatus) {
	//err = qmsql.DEFAULTDB.Save(rs).Error
	//return err, *rs

	db := qmsql.DEFAULTDB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var tmp ResumeStatus
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&tmp, rs.ID).Error; err != nil {
		tx.Rollback()
		return err, *rs
	}
	if err := tx.Save(rs).Error; err != nil {
		tx.Rollback()
		return err, *rs
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err, *rs
	}
	return err, *rs
}

func (rs *ResumeStatus) UpdateResumeByOpenidAndJobid() (err error) {
	return qmsql.DEFAULTDB.Model(rs).Where("open_id = ? and job_id = ?", rs.Openid, rs.Jobid).Update("resume_status", rs.ResumeStatus).Error
}

func (rs *ResumeStatus) EnterJob() (err error) {
	if rs.ResumeStatus == 6 { //入职
		bf := &UserBaseInfo{
			Openid: rs.Openid,
		}
		err, rebf := bf.FindByOpenid()
		if err != nil {
			return err
		}
		if len(rebf.IdCard) == 0 {
			return errors.New("请先去 '简历编辑' 补充身份证号")
		}
		if err = rs.UpdateResumeByOpenidAndJobid(); err != nil {
			return err
		}
		us := &userSalary.SalaryUsers{
			Name:         rebf.UserName,
			Mobile:       rebf.Mobile,
			Card:         rebf.IdCard,
			Enterprise:   rs.CompanyName,
			EnterpriseId: rs.CompanyId,
			JobName:      rs.Jobname,
			Email:        rebf.Email,
		}
		return us.CreateSalaryUsers()

	} else if rs.ResumeStatus == 7 {
		if err = rs.UpdateResumeByOpenidAndJobid(); err != nil {
			return err
		}
	}
	return nil
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
		var newStatusList []ResumeStatus
		model := qmsql.DEFAULTDB.Model(info)
		if rs.CompanyId != 0 {
			model = model.Where("company_id = ?", rs.CompanyId)
			db = db.Where("company_id = ?", rs.CompanyId)
		}
		if rs.CompanyName != "" {
			model = model.Where("company_name LIKE ?", "%"+rs.CompanyName+"%")
			db = db.Where("company_name LIKE ?", "%"+rs.CompanyName+"%")
		}
		if rs.Jobname != "" {
			model = model.Where("job_name LIKE ?", "%"+rs.Jobname+"%")
			db = db.Where("job_name LIKE ?", "%"+rs.Jobname+"%")
		}
		if rs.ResumeStatuSearch != nil {
			model = model.Where("resume_status = ?", rs.ResumeStatuSearch)
			db = db.Where("resume_status = ?", rs.ResumeStatuSearch)
		}

		err = model.Preload("JobInfo").Preload("UserInfo").Find(&reResumeStatusList).Count(&total).Error
		if err != nil {
			return err, reResumeStatusList, total
		}

		err = db.Preload("JobInfo").Preload("UserInfo").Find(&reResumeStatusList).Error

		if rs.Mobile != "" {
			total = 0
			for _, lst := range reResumeStatusList {
				total++
				if strings.Contains(lst.UserInfo.Mobile, rs.Mobile) {
					newStatusList = append(newStatusList, lst)
				}
			}
			return err, newStatusList, total
		}

		if rs.UserName != "" {
			total = 0
			for _, lst := range reResumeStatusList {
				total++
				if strings.Contains(lst.UserInfo.UserName, rs.UserName) {
					newStatusList = append(newStatusList, lst)
				}
			}
			return err, newStatusList, total
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
		err = qmsql.DEFAULTDB.Model(new(Joblist)).Where("id = ? and p_status != 3 and (p_outdate >= ? or p_outdate = 0 ) and p_count > 0", v.Jobid, time.Now().Unix()).Find(jb).Error
		if err != nil {
			continue
		}
		rs.Results = append(rs.Results, *jb)
	}
	return err, *rs, total
}

func (rs *ResumeStatus) GetCount() (err error, count int) {
	err = qmsql.DEFAULTDB.Model(rs).Where("job_id = ?", rs.Jobid).Count(&count).Error
	return err, count
}
