// 自动生成模板Joblist
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userCity"
	"time"

	"github.com/jinzhu/gorm"
)

type Joblist struct {
	gorm.Model
	Jobname      string  `json:"p_name" gorm:"column:job_name;comment:'工作名称'"`
	JobsalaId    int     `json:"p_salary_id" gorm:"column:job_salary_id;comment:'薪资限度id，关联job_salaries表'"`
	Jobsala      string  `json:"p_salary" gorm:"column:job_salary;comment:'薪资限度，关联job_salaries表名称'"`
	JobLatitude  float32 `json:"p_latitude" gorm:"column:job_latitude;comment:'工作地点纬度'"`
	JobLongitude float32 `json:"p_longitude" gorm:"column:job_longitude;comment:'工作地点经度'"`
	JobAddress   string  `json:"p_address" gorm:"column:job_address;comment:'工作地点'"`
	JobCity      string  `json:"p_city" gorm:"column:job_city;comment:'工作城市'"`
	JobCityId    int     `json:"p_city_id" gorm:"column:job_city_id;comment:'工作城市id，关联citynames表id'"`
	JobYears     string  `json:"p_edujy" gorm:"column:job_years;comment:'工作年限'"`
	JobYearsId   int     `json:"p_edujy_id" gorm:"column:job_years_id;comment:'工作年限id，关联job_work_expire表id'"`
	JobEdu       string  `json:"p_education" gorm:"column:job_edu;comment:'工作最低学历'"`
	JobEduId     int     `json:"p_education_id" gorm:"column:job_edu_id;comment:'学历id，-1其他0不限1初中2高中3中专4大专5本科6硕士7博士'"`
	JobType      string  `json:"p_type" gorm:"column:job_type;comment:'工作类型'"`
	JobTypeId    int     `json:"p_type_id" gorm:"column:job_type_id;comment:'工作类型id，0 全职1 兼职2 实习3 志愿者'"`
	CompanyName  string  `json:"enterprise_name" gorm:"column:company_name;comment:'公司名称'"`
	CompanyId    int     `json:"enterprise_id" gorm:"column:company_id;comment:'公司id，关联enterprise_infos表'"`
	CompanyImg   string  `json:"enterprise_logo" gorm:"column:company_img;comment:'公司logo url'"`
	JobDes       string  `json:"p_desc" gorm:"column:job_des;comment:'工作描述'"`
	Views        int     `json:"p_views" gorm:"column:p_views;comment:'浏览数量'"`
	Status       int     `json:"p_status" gorm:"column:p_status;comment:'职位状态，0普通，1急招，2热门，3下线'"`
	Count        int     `json:"p_count" gorm:"column:p_count;comment:'招聘人数'"`
	OutDate      int64   `json:"p_outdate" gorm:"column:p_outdate;comment:'有效日期，如2020-08-08为失效日期的unix时间戳'"`
	Top          int     `json:"p_top" gorm:"column:p_top;comment:'0:普通1:置顶'"`
	Filter       int     `json:"p_filter" gorm:"column:p_filter;comment:'过滤状态，0不启用过滤，1启用过滤'"`
	SendEmail    string  `json:"p_send_email" gorm:"column:p_send_email;comment:'发送邮件,多个用分号划分'"`
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

//更新浏览数量加一
func (jl *Joblist) UpdateViews() {
	err, rejl := jl.FindById()
	if err != nil {
		return
	}
	qmsql.DEFAULTDB.Model(jl).Where("id = ?", jl.ID).Update("p_views", rejl.Views+1)
	return
}

// 分页获取Joblist
func (jl *Joblist) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(jl, info)
	if err != nil {
		return
	} else {
		var reJoblistList []Joblist
		model := qmsql.DEFAULTDB.Model(info)
		if jl.CompanyId != 0 {
			model = model.Where("company_id = ?", jl.CompanyId)
			db = db.Where("company_id = ?", jl.CompanyId)
		}
		if jl.CompanyName != "" {
			model = model.Where("company_name = ?", jl.CompanyName)
			db = db.Where("company_name = ?", jl.CompanyName)
		}
		if jl.Jobname != "" {
			model = model.Where("job_name = ?", jl.Jobname)
			db = db.Where("job_name = ?", jl.Jobname)
		}
		if jl.JobCity != "" {
			model = model.Where("job_city = ?", jl.JobCity)
			db = db.Where("job_city = ?", jl.JobCity)
		}
		err = model.Find(&reJoblistList).Count(&total).Error
		if err != nil {
			return err, reJoblistList, total
		} else {
			err = db.Find(&reJoblistList).Error
		}
		return err, reJoblistList, total
	}
}

func (jl *Joblist) GetInfoListSearchDetail(keyword string, cityId, order int, industrys, enterpriseTypes, salarys, expires, edus, jobTyps []int, enterpriseScales []Scale, info modelInterface.PageInfo) (err error, list interface{}, total int) {

	reJoblistList := make([]Joblist, 0)

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	enterprises, err := new(EnterpriseInfo).GetEnterpeiseSearchDetail(industrys, enterpriseTypes, enterpriseScales)
	if err != nil {
		return nil, reJoblistList, 0
	}

	var enterpriseIds []uint
	if len(enterprises) > 0 {
		for _, enterprise := range enterprises {
			enterpriseIds = append(enterpriseIds, enterprise.ID)
		}
	} else {
		return nil, reJoblistList, 0
	}

	outData := time.Now().Unix()

	db := qmsql.DEFAULTDB.Model(jl).Where("p_status != 3 and p_outdate >= ?", outData)

	if cityId > 0 {
		db = db.Where("job_city_id = ?", cityId)
	}
	if len(keyword) > 0 {
		db = db.Where("job_name LIKE ? or company_name LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if len(salarys) > 0 {
		db = db.Where("job_salary_id IN (?)", salarys)
	}
	if len(expires) > 0 {
		db = db.Where("job_years_id IN (?)", expires)
	}
	if len(edus) > 0 {
		db = db.Where("job_edu_id IN (?)", edus)
	}
	if len(jobTyps) > 0 {
		db = db.Where("job_type_id IN (?)", jobTyps)
	}
	if len(enterpriseIds) > 0 {
		db = db.Where("company_id IN (?)", enterpriseIds)
	}

	db = db.Order("p_top desc")

	if order == 0 {
		db = db.Order("id desc")
	} else {
		db = db.Order("p_views desc")
	}

	db = db.Count(&total)

	err = db.Limit(limit).Offset(offset).Find(&reJoblistList).Error
	return err, reJoblistList, total
}

// 分页获取Joblist
func (jl *Joblist) GetInfoListSearch(keyword string, cityId int, low int, hight int, info modelInterface.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err = qmsql.DEFAULTDB.Model(jl).Count(&total).Error
	if err != nil {
		return
	}
	var reJoblistList []Joblist
	outData := time.Now().Unix()
	db := qmsql.DEFAULTDB.Model(jl)
	if cityId > 0 {
		db = db.Where("job_city_id = ?", cityId)
	}
	if len(keyword) > 0 {
		db = db.Where("job_name LIKE ? or company_name LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if low > 0 {
		db = db.Where("job_salary_low >= ?", low)
	}
	if hight > 0 {
		db = db.Where("job_salary_high >= ?", hight)
	}
	db = db.Where("p_status != 3 and p_outdate >= ?", outData)
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&reJoblistList).Error
	return err, reJoblistList, total
}

// 分页获取Joblist
func (jl *Joblist) GetInfoListSearchSimilar(ids []int, name string, eduJyId, eduId, jobtypeId, cityId, low, page, pageSize int) (err error, list interface{}) {
	limit := pageSize
	offset := pageSize * (page - 1)
	outData := time.Now().Unix()
	var reJoblistList []Joblist

	db := qmsql.DEFAULTDB.Model(jl)

	db = db.Where("job_city_id = ? AND job_edu_id = ? AND job_years_id = ? ", cityId, eduId, eduJyId)

	db = db.Where("job_type_id = ? AND job_salary_low >= ?", jobtypeId, low)

	if len(name) > 0 {
		db = db.Where("job_name LIKE ?", "%"+name+"%")
	}
	if len(ids) > 0 {
		db = db.Where("id NOT IN (?)", ids)
	}
	db = db.Where("p_status != 3 and p_outdate >= ?", outData)
	err = db.Limit(limit).Offset(offset).Order("p_top desc").Order("id desc").Find(&reJoblistList).Error
	return err, reJoblistList
}

func (jl *Joblist) GetAllInfoOption() (err error, list1, list2, list3, list4, list5 interface{}) {
	var jbe []JobWorkExpire
	var js []JobSalary
	var jwt []JobWorkType
	var el []EduLevel
	var citys []userCity.Cityname
	err = qmsql.DEFAULTDB.Select("id,name").Find(&jbe).Error
	if err != nil {
		return err, jbe, js, jwt, el, citys
	}
	err = qmsql.DEFAULTDB.Select("id,name").Find(&js).Error
	if err != nil {
		return err, jbe, js, jwt, el, citys
	}
	err = qmsql.DEFAULTDB.Select("id,name").Find(&citys).Error
	if err != nil {
		return err, jbe, js, jwt, el, citys
	}
	err = qmsql.DEFAULTDB.Select("id,name").Find(&jwt).Error
	if err != nil {
		return err, jbe, js, jwt, el, citys
	}
	err = qmsql.DEFAULTDB.Select("id,name").Find(&el).Error
	if err != nil {
		return err, jbe, js, jwt, el, citys
	}
	return err, jbe, js, jwt, el, citys
}
