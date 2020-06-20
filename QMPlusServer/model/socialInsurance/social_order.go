package socialInsurance

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
)

type SocialOrder struct {
	gorm.Model
	OrderId             string  `json:"order_id" gorm:"column:order_id;comment:'订单编号'"`
	InsuredType         int     `json:"insured_type" gorm:"column:insured_type;comment:'参保类型:1、城镇社保2、农村社保3、其他'"`
	Openid              string  `json:"openid" gorm:"column:openid;comment:'用户openid'"`
	Status              int     `json:"status" gorm:"column:status;comment:'订单状态 [1,2,3],[待支付,待激活,已完成]'"`
	Duration            int     `json:"duration" gorm:"column:duration;comment:'缴费时长：单位（月）'"`
	PayAmount           float64 `json:"pay_amount" gorm:"column:pay_amount;comment:'支付金额=单月总价*时长'"`
	IsRefund            int     `json:"is_refund" gorm:"column:is_refund;comment:'[1,2,3][未退单,已退单,退单中]'"`
	EndowmentIns        float64 `json:"endowment_ins" gorm:"column:endowment_ins;comment:'单月养老保险'"`
	MedicalIns          float64 `json:"medical_ins" gorm:"column:medical_ins;comment:'单月医疗保险'"`
	UnemploymentIns     float64 `json:"unemployment_ins" gorm:"column:unemployment_ins;comment:'单月失业保险'"`
	EmploymentInjuryIns float64 `json:"employment_injury_ins" gorm:"column:employment_injury_ins;comment:'单月工伤保险'"`
	MaternityIns        float64 `json:"maternity_ins" gorm:"column:maternity_ins;comment:'单月生育保险'"`
	InsuranceFee        float64 `json:"insurance_fee" gorm:"column:insurance_fee;comment:'单月保险合计'"`
	ServiceFee          float64 `json:"serve_fee" gorm:"column:service_fee;comment:'单月服务费'"`
	TotalFee            float64 `json:"total_fee" gorm:"column:total_fee;comment:'单月总价'"`
}

// 创建SocialOrder
func (so *SocialOrder) CreateSocialOrder() (err error) {
	err = qmsql.DEFAULTDB.Create(so).Error
	return err
}

// 删除SocialOrder
func (so *SocialOrder) DeleteSocialOrder() (err error) {
	err = qmsql.DEFAULTDB.Delete(so).Error
	return err
}

// 更新SocialOrder
func (so *SocialOrder) UpdateSocialOrder() (err error, reet SocialOrder) {
	err = qmsql.DEFAULTDB.Save(so).Error
	return err, *so
}

// 根据ID查看单条SocialOrder
func (so *SocialOrder) FindById() (err error, reet SocialOrder) {
	err = qmsql.DEFAULTDB.Where("id = ?", so.ID).First(&reet).Error
	return err, reet
}

// 分页获取SocialOrder
func (so *SocialOrder) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServerAsc(so, info)
	if err != nil {
		return
	} else {
		var reSocialOrderList []SocialOrder
		err = db.Find(&reSocialOrderList).Error
		return err, reSocialOrderList, total
	}
}
