package socialInsurance

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type OrderReqInfo struct {
	RemoteAddr string `json:"remoteaddr"`
	Openid     string `json:"openid"`
	OrderNo    string `json:"orderno"`
	TotalFee   int    `json:"totalfee"`
}

type SocialOrder struct {
	gorm.Model
	OrderId             string          `json:"order_id" gorm:"column:order_id;comment:'订单编号'"`
	InsBase             float64         `json:"ins_base" gorm:"column:ins_base;comment:'缴费基数'"`
	InsuredType         int             `json:"insured_type" gorm:"column:insured_type;comment:'参保类型:1、城镇社保2、农村社保3、其他'"`
	SocialId            int             `json:"social_id" gorm:"column:social_id;comment:'社保id，关联socialinsurances'"`
	SiInfo              SocialInsurance `json:"si_info" gorm:"ForeignKey:SocialId;AssociationForeignKey:ID;save_associations:false"`
	Name                string          `json:"name" gorm:"column:name;comment:'参保人姓名'"`
	IdCard              string          `json:"id_card" gorm:"column:id_card;comment:'参保人身份证号'"`
	Openid              string          `json:"openid" gorm:"column:openid;comment:'用户openid'"`
	Status              int             `json:"status" gorm:"column:status;comment:'订单状态 [1,2,3],[待支付,待激活,已完成]'"`
	Duration            int             `json:"duration" gorm:"column:duration;comment:'缴费时长：单位（月）'"`
	PayAmount           float64         `json:"pay_amount" gorm:"column:pay_amount;comment:'支付金额=单月总价*时长'"`
	IsRefund            int             `json:"is_refund" gorm:"column:is_refund;comment:'[1,2,3][未退单,已退单,退单中]'"`
	EndowmentIns        float64         `json:"endowment_ins" gorm:"column:endowment_ins;comment:'单月养老保险'"`
	MedicalIns          float64         `json:"medical_ins" gorm:"column:medical_ins;comment:'单月医疗保险'"`
	UnemploymentIns     float64         `json:"unemployment_ins" gorm:"column:unemployment_ins;comment:'单月失业保险'"`
	EmploymentInjuryIns float64         `json:"employment_injury_ins" gorm:"column:employment_injury_ins;comment:'单月工伤保险'"`
	MaternityIns        float64         `json:"maternity_ins" gorm:"column:maternity_ins;comment:'单月生育保险'"`
	InsuranceFee        float64         `json:"insurance_fee" gorm:"column:insurance_fee;comment:'单月保险合计'"`
	IsGjj               int             `json:"is_gjj" gorm:"column:is_gjj;comment:'是否缴公积金：0、否1、是'"`
	GjjFee              float64         `json:"gjj_fee" gorm:"column:gjj_fee;comment:'单月公积金'"`
	ServiceFee          float64         `json:"serve_fee" gorm:"column:service_fee;comment:'单月服务费'"`
	TotalFee            float64         `json:"total_fee" gorm:"column:total_fee;comment:'单月总价=单月保险合计+单月公积金+单月服务费'"`
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
	err, db, total := servers.PagingServer(so, info)
	if err != nil {
		return
	} else {
		var reSocialOrderList []SocialOrder
		model := qmsql.DEFAULTDB.Model(info)
		if so.Name != "" {
			model = model.Where("name LIKE ?", "%"+so.Name+"%")
			db = db.Where("name LIKE ?", "%"+so.Name+"%")
		}
		if so.Status != 0 {
			model = model.Where("status = ?", so.Status)
			db = db.Where("status = ?", so.Status)
		}
		if so.IdCard != "" {
			model = model.Where("id_card = ?", so.IdCard)
			db = db.Where("id_card = ?", so.IdCard)
		}
		if so.OrderId != "" {
			model = model.Where("order_id = ?", so.OrderId)
			db = db.Where("order_id = ?", so.OrderId)
		}
		err = model.Find(&reSocialOrderList).Count(&total).Error
		if err != nil {
			return err, reSocialOrderList, total
		} else {
			err = db.Preload("SiInfo").Find(&reSocialOrderList).Error
		}
		return err, reSocialOrderList, total
	}
}
