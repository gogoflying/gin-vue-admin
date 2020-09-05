package socialInsurance

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"math"
	"time"

	"github.com/jinzhu/gorm"
)

type OrderReqInfo struct {
	Openid   string  `json:"openid"`
	OrderNo  string  `json:"orderno"`
	TotalFee float64 `json:"totalfee"`
}

type RefundReqInfo struct {
	Openid    string  `json:"openid"`
	OrderNo   string  `json:"orderno"`
	TotalFee  float64 `json:"totalfee"`
	RefundFee float64 `json:"refundfee"`
}

type ReqAddOrder struct {
	Openid      string `json:"openid"`
	Cityindex   int    `json:"cityindex"`
	InsuredType int    `json:"insured_type"`
	Name        string `json:"name"`
	IdCard      string `json:"id_card"`
	StartTime   string `json:"start_time"`
	Duration    int    `json:"duration"`
	IsIns       int    `json:"is_ins"`
	IsGjj       int    `json:"is_gjj"`
}

type SocialOrder struct {
	gorm.Model
	OrderId             string          `json:"order_id" gorm:"column:order_id;comment:'订单编号'"`
	Name                string          `json:"name" gorm:"column:name;comment:'参保人姓名'"`
	IdCard              string          `json:"id_card" gorm:"column:id_card;comment:'参保人身份证号'"`
	Openid              string          `json:"openid" gorm:"column:openid;comment:'用户openid'"`
	Status              int             `json:"status" gorm:"column:status;comment:'订单状态 [1,2,3,4,5],[待支付,待激活,已完成,已失效,已放弃]'"`
	StartTime           string          `json:"start_time" gorm:"column:start_time;comment:'开始时间'"`
	EndTime             string          `json:"end_time" gorm:"column:end_time;comment:'结束时间'"`
	Duration            int             `json:"duration" gorm:"column:duration;comment:'缴费时长：单位（月）'"`
	PayAmount           float64         `json:"pay_amount" gorm:"column:pay_amount;comment:'支付金额=单月总价*时长'"`
	IsRefund            int             `json:"is_refund" gorm:"column:is_refund;comment:'[1,2,3][未退单,已退单,退单中]'"`
	IsIns               int             `json:"is_ins" gorm:"column:is_ins;comment:'是否缴纳社保：0、否1、是'"`
	InsBase             float64         `json:"ins_base" gorm:"column:ins_base;comment:'社保缴费基数'"`
	InsuredType         *int            `json:"insured_type" gorm:"column:insured_type;comment:'参保类型:1、本地城镇2、本地农村3、外地城镇4、外地农村5、其他'"`
	EndowmentIns        float64         `json:"endowment_ins" gorm:"column:endowment_ins;comment:'单月养老保险'"`
	MedicalIns          float64         `json:"medical_ins" gorm:"column:medical_ins;comment:'单月医疗保险'"`
	UnemploymentIns     float64         `json:"unemployment_ins" gorm:"column:unemployment_ins;comment:'单月失业保险'"`
	EmploymentInjuryIns float64         `json:"employment_injury_ins" gorm:"column:employment_injury_ins;comment:'单月工伤保险'"`
	MaternityIns        float64         `json:"maternity_ins" gorm:"column:maternity_ins;comment:'单月生育保险'"`
	InsuranceFee        float64         `json:"insurance_fee" gorm:"column:insurance_fee;comment:'单月保险合计'"`
	IsGjj               int             `json:"is_gjj" gorm:"column:is_gjj;comment:'是否缴纳公积金：0、否1、是'"`
	GjjBase             float64         `json:"gjj_base" gorm:"column:gjj_base;comment:'公积金缴费基数'"`
	GjjFee              float64         `json:"gjj_fee" gorm:"column:gjj_fee;comment:'单月公积金'"`
	ServiceFee          float64         `json:"serve_fee" gorm:"column:service_fee;comment:'单月服务费'"`
	TotalFee            float64         `json:"total_fee" gorm:"column:total_fee;comment:'单月总价=单月保险合计+单月公积金+单月服务费'"`
	SocialId            int             `json:"social_id" gorm:"column:social_id;comment:'社保id，关联socialinsurances'"`
	SiInfo              SocialInsurance `json:"si_info" gorm:"ForeignKey:SocialId;AssociationForeignKey:ID;save_associations:false"`
}

func (so *SocialOrder) AddSocialOrder(req ReqAddOrder) (err error, ret SocialOrder) {
	var totalMonth float64
	var tmpSo SocialOrder
	var count int
	qmsql.DEFAULTDB.Model(so).Where("openid = ? and status = 1", req.Openid).Find(&tmpSo).Count(&count)
	if count > 0 {
		err = fmt.Errorf("存在未支付订单，请处理")
		return err, *so
	}

	siInfo := new(SocialInsurance)
	err, reet := siInfo.GetByCityIndexAndType(req.Cityindex, req.InsuredType)
	if err != nil {
		return err, *so
	}
	so.OrderId = fmt.Sprintf("%v", time.Now().UnixNano())
	so.Name = req.Name
	so.IdCard = req.IdCard
	so.Openid = req.Openid
	so.Status = 1
	so.StartTime = req.StartTime
	so.Duration = req.Duration
	so.IsRefund = 1
	so.InsuredType = &req.InsuredType
	so.SocialId = int(reet.ID)

	if req.IsIns == 1 {
		so.IsIns = req.IsIns
		so.InsBase = reet.YangLaoLowerLimit
		so.EndowmentIns = math.Floor(reet.YangLaoLowerLimit*(reet.YangLaoCompanyRatio+reet.YangLaoPersonRatio) + 0.5)
		so.MedicalIns = math.Floor(reet.YiLiaoLowerLimit*(reet.YiLiaoCompanyRatio+reet.YiLiaoPersonRatio) + 0.5)
		so.UnemploymentIns = math.Floor(reet.ShiYeLowerLimit*(reet.ShiYeCompanyRatio+reet.ShiYePersonRatio) + 0.5)
		so.MaternityIns = math.Floor(reet.ShengYuLowerLimit*(reet.ShengYuCompanyRatio+reet.ShengYuPersonRatio) + 0.5)
		so.EmploymentInjuryIns = math.Floor(reet.GongShangLowerLimit*(reet.GongShangCompanyRatio+reet.GongShangPersonRatio) + 0.5)
		so.ServiceFee += reet.InsuredServiceFee
		totalMonth = so.EndowmentIns + so.MedicalIns + so.UnemploymentIns + so.MaternityIns + so.EmploymentInjuryIns
	}

	if req.IsGjj == 1 {
		so.IsGjj = req.IsGjj
		so.GjjBase = reet.GJJLowerLimit
		so.GjjFee = math.Floor(reet.GJJLowerLimit*(reet.GJJCompanyRatio+reet.GJJPersonRatio) + 0.5)
		so.ServiceFee += reet.GJJServiceFee
		totalMonth += so.GjjFee
	}

	so.TotalFee = (totalMonth + so.ServiceFee) * float64(so.Duration)

	err = so.CreateSocialOrder()

	return err, *so

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

func (so *SocialOrder) UpdateSocialOrderStatus(openid, orderid string, status int) (err error) {
	return qmsql.DEFAULTDB.Model(so).Where("openid = ? and order_id = ?", openid, orderid).Update("status", status).Error
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

func (so *SocialOrder) FindByOrderId() (err error, reet SocialOrder) {
	err = qmsql.DEFAULTDB.Where("status = ?", so.Status).First(&reet).Error
	return err, reet
}

func (so *SocialOrder) GetOrderList(openid string, status, page, pageSize int) (err error, list interface{}) {
	limit := pageSize
	offset := pageSize * (page - 1)

	var orderList []SocialOrder
	model := qmsql.DEFAULTDB.Model(so)

	err = model.Where("openid = ? and status = ?", openid, status).Limit(limit).Offset(offset).Order("id desc").Find(&orderList).Error

	return err, orderList
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
