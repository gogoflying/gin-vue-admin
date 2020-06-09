package userSalary

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
)

/*基本工资、岗位工资、薪资合计、业绩提成、奖金基数、浮动系数、月度奖金、本月工作日天数、加班费、通讯补贴、餐食补贴、交通补贴、其他补贴、补贴合计、其他假期、年假天数、迟到扣款、病假天数、病假扣款、事假天数、事假扣款、扣款合计、应发调整、本月应发工资、代扣五险、代扣住房公积金、代扣个人所得税、实发工资*/
type SalaryTemplates struct {
	gorm.Model
	Enterprise   string `json:"enterprise" gorm:"column:enterprise;comment:'所属企业'"`
	EnterpriseId int    `json:"enterprise_id" gorm:"column:enterprise_id;comment:'所属企业id'"`
	GwTempName   string `json:"gwtemp_name" gorm:"column:gwtemp_name"`
	Base         string `json:"base" gorm:"column:base"`
	Gangwei      string `json:"gangwei" gorm:"column:gangwei"`
	Jjjs         string `json:"jjjs" gorm:"column:jjjs"`
	Gzts         string `json:"gzts" gorm:"column:gzts"`
	Txbt         string `json:"txbt" gorm:"column:txbt"`
	Csbt         string `json:"csbt" gorm:"column:csbt"`
	Jtbt         string `json:"jtbt" gorm:"column:jtbt"`
	Qtbt         string `json:"qtbt" gorm:"column:qtbt"`
	Bthj         string `json:"bthj" gorm:"column:bthj"`
	Dkwx         string `json:"dkwx" gorm:"column:dkwx"`
	Gjj          string `json:"gjj" gorm:"column:gjj"`
	Dkgs         string `json:"dkgs" gorm:"column:dkgs"`
}

// 创建SalaryTemplates
func (st *SalaryTemplates) CreateSalaryTemplates() (err error) {
	err = qmsql.DEFAULTDB.Create(st).Error
	return err
}

// 删除SalaryTemplates
func (st *SalaryTemplates) DeleteSalaryTemplates() (err error) {
	err = qmsql.DEFAULTDB.Delete(st).Error
	return err
}

// 更新SalaryTemplates
func (st *SalaryTemplates) UpdateSalaryTemplates() (err error, reet SalaryTemplates) {
	err = qmsql.DEFAULTDB.Save(st).Error
	return err, *st
}

// 根据ID查看单条SalaryTemplates
func (st *SalaryTemplates) FindById() (err error, reet SalaryTemplates) {
	err = qmsql.DEFAULTDB.Where("id = ?", st.ID).First(&reet).Error
	return err, reet
}

// 分页获取SalaryTemplates
func (st *SalaryTemplates) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServerAsc(st, info)
	if err != nil {
		return
	} else {
		var reSalaryTemplatesList []SalaryTemplates
		model := qmsql.DEFAULTDB.Model(info)
		if st.EnterpriseId != 0 {
			model = model.Where("enterprise_id = ?", st.EnterpriseId)
			db = db.Where("enterprise_id = ?", st.EnterpriseId)
		}
		if st.GwTempName != "" {
			model = model.Where("gwtemp_name LIKE ?", "%"+st.GwTempName+"%")
			db = db.Where("gwtemp_name LIKE ?", "%"+st.GwTempName+"%")
		}
		err = model.Find(&reSalaryTemplatesList).Count(&total).Error
		if err != nil {
			return err, reSalaryTemplatesList, total
		} else {
			err = db.Find(&reSalaryTemplatesList).Error
		}
		return err, reSalaryTemplatesList, total
	}
}

func (st *SalaryTemplates) GetAllInfoList(query string) (err error, list interface{}) {
	var reSalaryTemplatesList []SalaryTemplates
	db := qmsql.DEFAULTDB
	if st.EnterpriseId != 0 {
		db = db.Where("enterprise_id = ?", st.EnterpriseId)
	}
	if len(query) > 0 {
		db = db.Where("gwtemp_name LIKE ? ", "%"+query+"%")
	}
	err = db.Find(&reSalaryTemplatesList).Error
	return err, reSalaryTemplatesList
}
