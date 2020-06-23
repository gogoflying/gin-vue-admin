package socialInsurance

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userCity"
	"github.com/jinzhu/gorm"
)

type SocialInsurance struct {
	gorm.Model
	Cityindex   int               `json:"cityindex" gorm:"column:cityindex;comment:'城市id，关联citynames'"`
	Cityname    userCity.Cityname `json:"city_name" gorm:"ForeignKey:Cityindex;AssociationForeignKey:ID;save_associations:false"`
	InsuredType int               `json:"insured_type" gorm:"column:insured_type;comment:'参保类型:1、城镇社保2、农村社保3、其他'"`
	//InsInfo               string            `json:"ins_info" gorm:"column:ins_info;comment:'缴费比例:json格式'"`
	YangLaoLowerLimit     float64 `json:"yanglao_lower_limit" gorm:"column:yanglao_lower_limit;comment:'养老缴费基数下限'"`
	YangLaoUpperLimit     float64 `json:"yanglao_upper_limit" gorm:"column:yanglao_upper_limit;comment:'养老缴费基数上限'"`
	YangLaoCompanyRatio   float64 `json:"yanglao_company_ratio" gorm:"column:yanglao_company_ratio;comment:'养老企业缴费比例'"`
	YangLaoPersonRatio    float64 `json:"yanglao_person_ratio" gorm:"column:yanglao_person_ratio;comment:'养老个人缴费比例'"`
	YiLiaoLowerLimit      float64 `json:"yiliao_lower_limit" gorm:"column:yiliao_lower_limit;comment:'医疗缴费基数下限'"`
	YiLiaoUpperLimit      float64 `json:"yiliao_upper_limit" gorm:"column:yiliao_upper_limit;comment:'医疗缴费基数上限'"`
	YiLiaoCompanyRatio    float64 `json:"yiliao_company_ratio" gorm:"column:yiliao_company_ratio;comment:'医疗企业缴费比例'"`
	YiLiaoPersonRatio     float64 `json:"yiliao_person_ratio" gorm:"column:yiliao_person_ratio;comment:'医疗个人缴费比例'"`
	ShengYuLowerLimit     float64 `json:"shengyu_lower_limit" gorm:"column:shengyu_lower_limit;comment:'生育缴费基数下限'"`
	ShengYuUpperLimit     float64 `json:"shengyu_upper_limit" gorm:"column:shengyu_upper_limit;comment:'生育缴费基数上限'"`
	ShengYuCompanyRatio   float64 `json:"shengyu_company_ratio" gorm:"column:shengyu_company_ratio;comment:'生育企业缴费比例'"`
	ShengYuPersonRatio    float64 `json:"shengyu_person_ratio" gorm:"column:shengyu_person_ratio;comment:'生育个人缴费比例'"`
	ShiYeLowerLimit       float64 `json:"shiye_lower_limit" gorm:"column:shiye_lower_limit;comment:'失业缴费基数下限'"`
	ShiYeUpperLimit       float64 `json:"shiye_upper_limit" gorm:"column:shiye_upper_limit;comment:'失业缴费基数上限'"`
	ShiYeCompanyRatio     float64 `json:"shiye_company_ratio" gorm:"column:shiye_company_ratio;comment:'失业企业缴费比例'"`
	ShiYePersonRatio      float64 `json:"shiye_person_ratio" gorm:"column:shiye_person_ratio;comment:'失业个人缴费比例'"`
	GongShangLowerLimit   float64 `json:"gongshang_lower_limit" gorm:"column:gongshang_lower_limit;comment:'工伤缴费基数下限'"`
	GongShangUpperLimit   float64 `json:"gongshang_upper_limit" gorm:"column:gongshang_upper_limit;comment:'工伤缴费基数上限'"`
	GongShangCompanyRatio float64 `json:"gongshang_company_ratio" gorm:"column:gongshang_company_ratio;comment:'工伤企业缴费比例'"`
	GongShangPersonRatio  float64 `json:"gongshang_person_ratio" gorm:"column:gongshang_person_ratio;comment:'工伤个人缴费比例'"`
	GJJLowerLimit         float64 `json:"gjj_lower_limit" gorm:"column:gjj_lower_limit;comment:'公积金缴费基数下限'"`
	GJJUpperLimit         float64 `json:"gjj_upper_limit" gorm:"column:gjj_upper_limit;comment:'公积金缴费基数上限'"`
	GJJCompanyRatio       float64 `json:"gjj_company_ratio" gorm:"column:gjj_company_ratio;comment:'公积金企业缴费比例'"`
	GJJPersonRatio        float64 `json:"gjj_person_ratio" gorm:"column:gjj_person_ratio;comment:'公积金个人缴费比例'"`
}

// 创建SocialInsurance
func (si *SocialInsurance) CreateSocialInsurance() (err error) {
	err = qmsql.DEFAULTDB.Create(si).Error
	return err
}

// 删除SocialInsurance
func (si *SocialInsurance) DeleteSocialInsurance() (err error) {
	err = qmsql.DEFAULTDB.Delete(si).Error
	return err
}

// 更新SocialInsurance
func (si *SocialInsurance) UpdateSocialInsurance() (err error, reet SocialInsurance) {
	err = qmsql.DEFAULTDB.Save(si).Error
	return err, *si
}

// 根据ID查看单条SocialInsurance
func (si *SocialInsurance) FindById() (err error, reet SocialInsurance) {
	err = qmsql.DEFAULTDB.Where("id = ?", si.ID).First(&reet).Error
	return err, reet
}

// 分页获取SocialInsurance
func (si *SocialInsurance) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(si, info)
	if err != nil {
		return
	} else {
		var reSocialInsuranceList []SocialInsurance
		model := qmsql.DEFAULTDB.Model(info)
		if si.Cityindex != 0 {
			model = model.Where("cityindex = ?", si.Cityindex)
			db = db.Where("cityindex = ?", si.Cityindex)
		}
		err = model.Find(&reSocialInsuranceList).Count(&total).Error
		if err != nil {
			return err, reSocialInsuranceList, total
		} else {
			err = db.Find(&reSocialInsuranceList).Error
		}
		return err, reSocialInsuranceList, total
	}
}

func (si *SocialInsurance) GetAllInfoOption() (err error, list interface{}) {
	var citys []userCity.Cityname
	err = qmsql.DEFAULTDB.Select("id,name").Where("type = 0").Find(&citys).Error
	if err != nil {
		return err, citys
	}
	return err, citys
}
