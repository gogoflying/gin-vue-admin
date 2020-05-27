// 自动生成模板EnterpriseInfo
package userJobs

import (
	"errors"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userCity"

	"github.com/jinzhu/gorm"
)

type Scale struct {
	Low  int `json:"low"`
	High int `json:"high"`
}

type EnterpriseInfo struct {
	gorm.Model
	EnterPriseName       string    `json:"enterprise_name" gorm:"column:enterprise_name;comment:'企业名称'"`
	EnterpriseLogo       string    `json:"enterprise_logo" gorm:"column:enterprise_logo;comment:'企业logo url'"`
	EnterpriseImg        string    `json:"enterprise_img" gorm:"column:enterprise_img;comment:'企业主图 url'"`
	EnterpriseQfc        string    `json:"enterprise_qfc" gorm:"column:enterprise_qfc;comment:'企业资质 url'"`
	EnterpriseAddress    string    `json:"enterprise_address" gorm:"column:enterprise_address;comment:'企业所在城市名称'"`
	EnterpriseScale      int       `json:"enterprise_scale" gorm:"column:enterprise_scale;comment:'企业人员规模，如200，小程序会显示为200+'"`
	EnterpriseType       int       `json:"enterprise_type" gorm:"column:enterprise_type;comment:'企业类型id，关联enterprise_type'"`
	EnterpriseTypeInfo   string    `json:"enterprise_type_info" gorm:"column:enterprise_type_info;comment:'企业类型'"`
	EnterpriseHot        int       `json:"enterprise_hot" gorm:"column:enterprise_hot;comment:'企业热度'"`
	IndustryType         int       `json:"industry_type" gorm:"column:industry_type;comment:'行业类型id，关联enterprise_industs'"`
	EnterpriseIndustInfo string    `json:"enterprise_indust_info" gorm:"column:enterprise_indust_info;comment:'行业类型'"`
	EnterpriseDesc       string    `json:"enterprise_desc" gorm:"column:enterprise_desc;comment:'企业描述信息'"`
	EnterpriseCityId     int       `json:"city_id" gorm:"column:city_id;comment:'城市id，关联citynames'"`
	JobCount             int       `json:"job_count" gorm:"column:job_count;comment:'该企业发布的职位数量'"`
	Status               int       `json:"status" gorm:"column:status;comment:'0、未审核1、审核通过'"`
	Errors               string    `json:"errors" gorm:"column:errors;comment:'信息错误反馈内容'"`
	Results              []Joblist `json:"result" gorm:"-"`
}

func (info *EnterpriseInfo) GetEnterpeiseSearchDetail(industrys, enterpriseTypes []int, enterpriseScales []Scale) (infos []EnterpriseInfo, err error) {

	db := qmsql.DEFAULTDB.Model(info).Where("status = 1")

	if len(industrys) > 0 {
		db = db.Where("industry_type IN (?)", industrys)
	}
	if len(enterpriseTypes) > 0 {
		db = db.Where("enterprise_type IN (?)", enterpriseTypes)
	}
	if len(enterpriseScales) > 0 {
		db = db.Where("enterprise_scale >= ? and enterprise_scale <= ?", enterpriseScales[0].Low, enterpriseScales[0].High)
	}
	if len(enterpriseScales) > 1 {
		for i := 1; i < len(enterpriseScales); i++ {
			db = db.Or("enterprise_scale >= ? and enterprise_scale <= ?", enterpriseScales[i].Low, enterpriseScales[i].High)
		}
	}
	err = db.Find(&infos).Error
	return
}

// 创建EnterpriseInfo
func (info *EnterpriseInfo) CreateEnterpriseInfo() (err error) {
	var ei EnterpriseInfo
	//判断用户名是否注册
	notResigt := qmsql.DEFAULTDB.Where("enterprise_name = ?", info.EnterPriseName).First(&ei).RecordNotFound()
	//notResigt为false表明读取到了 不能注册
	if !notResigt {
		return errors.New("企业已注册")
	} else {
		err = qmsql.DEFAULTDB.Create(info).Error
	}
	return err
}

// 创建EnterpriseInfo
func (info *EnterpriseInfo) SelectAndCreateEnterpriseInfo() (err error, ep *EnterpriseInfo) {
	var ei EnterpriseInfo
	//判断用户名是否注册
	notResigt := qmsql.DEFAULTDB.Where("enterprise_name = ?", info.EnterPriseName).First(&ei).RecordNotFound()
	//notResigt为false表明读取到了 不能注册
	if !notResigt {
		return nil, &ei
	} else {
		err = qmsql.DEFAULTDB.Create(info).Error
	}
	return err, info
}

// 删除EnterpriseInfo
func (info *EnterpriseInfo) DeleteEnterpriseInfo() (err error) {
	err = qmsql.DEFAULTDB.Delete(info).Error
	return err
}

// 更新EnterpriseInfo
func (info *EnterpriseInfo) UpdateEnterpriseInfo() (err error, reinfo EnterpriseInfo) {
	//err = qmsql.DEFAULTDB.Save(info).Error
	//return err, *info

	db := qmsql.DEFAULTDB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var tmp EnterpriseInfo
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&tmp, info.ID).Error; err != nil {
		tx.Rollback()
		return err, *info
	}
	if err := tx.Save(info).Error; err != nil {
		tx.Rollback()
		return err, *info
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err, *info
	}
	return err, *info
}

// 根据ID查看单条EnterpriseInfo
func (info *EnterpriseInfo) FindById() (err error, reinfo EnterpriseInfo) {
	err = qmsql.DEFAULTDB.Where("id = ?", info.ID).First(&reinfo).Error

	return err, reinfo
}

func (info *EnterpriseInfo) GeteEpById(id int) (err error, reinfo EnterpriseInfo) {
	err = qmsql.DEFAULTDB.Where("id = ? and status = 1", id).First(&reinfo).Error
	return
}

// 根据ID查看单条EnterpriseInfo
func (info *EnterpriseInfo) FindPositionsByCompId(page, pageSize int) (err error, reinfo EnterpriseInfo) {
	err = qmsql.DEFAULTDB.Where("id = ?", info.ID).First(&reinfo).Error
	if err != nil {
		return
	}
	limit := pageSize
	offset := pageSize * (page - 1)
	var jobList []Joblist
	err = qmsql.DEFAULTDB.Model(new(Joblist)).Where("company_id = ?", info.ID).Limit(limit).Offset(offset).Order("id desc").Find(&jobList).Error
	if err != nil {
		return
	}
	reinfo.Results = append(reinfo.Results, jobList...)
	return err, reinfo
}

// 分页获取EnterpriseInfo
func (info *EnterpriseInfo) GetInfoList(pinfo modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(info, pinfo)
	if err != nil {
		return
	} else {
		reEnterpriseInfoList := make([]*EnterpriseInfo, 0)
		model := qmsql.DEFAULTDB.Model(pinfo)
		if info.EnterPriseName != "" {
			model = model.Where("enterprise_name LIKE ?", "%"+info.EnterPriseName+"%")
			db = db.Where("enterprise_name LIKE ?", "%"+info.EnterPriseName+"%")
		}

		if info.IndustryType > 0 {
			model = model.Where("industry_type = ?", info.IndustryType)
			db = db.Where("industry_type = ?", info.IndustryType)
		}

		if info.EnterpriseType > 0 {
			model = model.Where("enterprise_type = ?", info.EnterpriseType)
			db = db.Where("enterprise_type = ?", info.EnterpriseType)
		}

		if info.EnterpriseCityId > 0 {
			model = model.Where("city_id = ?", info.EnterpriseCityId)
			db = db.Where("city_id = ?", info.EnterpriseCityId)
		}

		err = model.Find(&reEnterpriseInfoList).Count(&total).Error
		if err != nil {
			return err, reEnterpriseInfoList, total
		} else {
			err = db.Find(&reEnterpriseInfoList).Error
			/*for _, reP := range reEnterpriseInfoList {
				var count int
				qmsql.DEFAULTDB.Model(new(Joblist)).Where("company_id = ?", reP.ID).Count(&count)
				reP.JobCount = count
			}*/
		}
		/*err = db.Find(&reEnterpriseInfoList).Error
		if err == nil {
			for _, reP := range reEnterpriseInfoList {
				var count int
				qmsql.DEFAULTDB.Model(new(Joblist)).Where("company_id = ?", reP.ID).Count(&count)
				reP.JobCount = count
			}
		}*/
		return err, reEnterpriseInfoList, total
	}
}

// 分页获取Joblist
func (info *EnterpriseInfo) GetInfoListSearch(keyword string, cityId, page, pageSize int) (err error, list interface{}, total int) {
	limit := pageSize
	offset := pageSize * (page - 1)
	err = qmsql.DEFAULTDB.Model(info).Count(&total).Error
	if err != nil {
		return
	}
	var infoList []EnterpriseInfo
	db := qmsql.DEFAULTDB.Model(info).Where("status = 1")
	if cityId > 0 {
		db = db.Where("city_id = ?", cityId)
	}
	if len(keyword) > 0 {
		db = db.Where("enterprise_name LIKE ? ", "%"+keyword+"%")
	}
	err = db.Limit(limit).Offset(offset).Order("enterprise_hot desc").Order("updated_at desc").Find(&infoList).Error
	return err, infoList, total
}

func (info *EnterpriseInfo) GetAllInfoList(query string) (err error, list interface{}) {
	var reEnterpriseInfoList []EnterpriseInfo
	db := qmsql.DEFAULTDB.Select("id,enterprise_name,enterprise_logo").Where("status = 1")
	if len(query) > 0 {
		db = db.Where("enterprise_name LIKE ? ", "%"+query+"%")
	}
	err = db.Find(&reEnterpriseInfoList).Error
	return err, reEnterpriseInfoList
}

func (info *EnterpriseInfo) GetAllInfoOption() (err error, list1 interface{}, list2 interface{}, list3 interface{}) {
	var enPtype []EnterpriseType
	var enPindust []EnterpriseIndust
	var citys []userCity.Cityname
	err = qmsql.DEFAULTDB.Select("id,name").Find(&enPtype).Error
	if err != nil {
		return err, enPtype, enPindust, citys
	}

	err = qmsql.DEFAULTDB.Select("id,name").Find(&enPindust).Error
	if err != nil {
		return err, enPtype, enPindust, citys
	}
	err = qmsql.DEFAULTDB.Select("id,name").Find(&citys).Error
	if err != nil {
		return err, enPtype, enPindust, citys
	}
	return err, enPtype, enPindust, citys
}
