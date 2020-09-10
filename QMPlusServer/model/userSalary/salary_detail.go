// 自动生成模板Salarys
package userSalary

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

/*基本工资、岗位工资、薪资合计、业绩提成、奖金基数、浮动系数、月度奖金、本月工作日天数、加班费、通讯补贴、餐食补贴、交通补贴、其他补贴、补贴合计、其他假期、年假天数、迟到扣款、病假天数、病假扣款、事假天数、事假扣款、扣款合计、应发调整、本月应发工资、代扣五险、代扣住房公积金、代扣个人所得税、实发工资*/
type Salarys struct {
	gorm.Model
	Enterprise   string `json:"enterprise" gorm:"column:enterprise;comment:'入职企业'"`
	EnterpriseId int    `json:"enterprise_id" gorm:"column:enterprise_id;comment:'入职企业id'"`
	Year         int    `json:"year" gorm:"column:year"`
	Month        int    `json:"month" gorm:"column:month"`
	Openid       string `json:"openid" gorm:"column:openid"`
	Name         string `json:"name" gorm:"column:name"`
	Base         string `json:"base" gorm:"column:base"`
	Gangwei      string `json:"gangwei" gorm:"column:gangwei"`
	Xzhj         string `json:"xzhj" gorm:"column:xzhj"`
	Yjtc         string `json:"yjtc" gorm:"column:yjtc"`
	Jjjs         string `json:"jjjs" gorm:"column:jjjs"`
	Fdxs         string `json:"fdxs" gorm:"column:fdxs"`
	Ydjj         string `json:"ydjj" gorm:"column:ydjj"`
	Gzts         string `json:"gzts" gorm:"column:gzts"`
	Jbf          string `json:"jbf" gorm:"column:jbf"`
	Txbt         string `json:"txbt" gorm:"column:txbt"`
	Csbt         string `json:"csbt" gorm:"column:csbt"`
	Jtbt         string `json:"jtbt" gorm:"column:jtbt"`
	Qtbt         string `json:"qtbt" gorm:"column:qtbt"`
	Bthj         string `json:"bthj" gorm:"column:bthj"`
	Qtjq         string `json:"qtjq" gorm:"column:qtjq"`
	Njts         string `json:"njts" gorm:"column:njts"`
	Cdkk         string `json:"cdkk" gorm:"column:cdkk"`
	Bjts         string `json:"bjts" gorm:"column:bjts"`
	Bjkk         string `json:"bjkk" gorm:"column:bjkk"`
	Sjts         string `json:"sjts" gorm:"column:sjts"`
	Sjkk         string `json:"sjkk" gorm:"column:sjkk"`
	Kkhj         string `json:"kkhj" gorm:"column:kkhj"`
	Yftz         string `json:"yftz" gorm:"column:yftz"`
	Byyf         string `json:"byyf" gorm:"column:byyf"`
	Dkwx         string `json:"dkwx" gorm:"column:dkwx"`
	Gjj          string `json:"gjj" gorm:"column:gjj"`
	Dkgs         string `json:"dkgs" gorm:"column:dkgs"`
	Sfgz         string `json:"sfgz" gorm:"column:sfgz"`

	Ext1  string `json:"ext1" gorm:"column:ext1"`
	Ext2  string `json:"ext2" gorm:"column:ext2"`
	Ext3  string `json:"ext3" gorm:"column:ext3"`
	Ext4  string `json:"ext4" gorm:"column:ext4"`
	Ext5  string `json:"ext5" gorm:"column:ext5"`
	Ext6  string `json:"ext6" gorm:"column:ext6"`
	Ext7  string `json:"ext7" gorm:"column:ext7"`
	Ext8  string `json:"ext8" gorm:"column:ext8"`
	Ext9  string `json:"ext9" gorm:"column:ext9"`
	Ext10 string `json:"ext10" gorm:"column:ext10"`
}

// 创建Salarys
func (un *Salarys) CreateSalarys() (err error) {
	err = qmsql.DEFAULTDB.Create(un).Error
	return err
}

func (un *Salarys) GetDbTx() *gorm.DB {
	return qmsql.DEFAULTDB
}

// 删除Salarys
func (un *Salarys) DeleteSalarys() (err error) {
	err = qmsql.DEFAULTDB.Delete(un).Error
	return err
}
func (un *Salarys) DeleteSalarysEx() (err error) {
	err = qmsql.DEFAULTDB.Unscoped().Delete(un).Error
	return err
}

// 批量删除JobSalary
func (un *Salarys) BatchDeleteSalarys(ids []int) (err error) {
	err = qmsql.DEFAULTDB.Unscoped().Where("id IN (?)", ids).Delete(Salarys{}).Error
	return err
}

// 更新Salarys
func (un *Salarys) UpdateSalarys() (err error, reun Salarys) {
	//err = qmsql.DEFAULTDB.Save(un).Error
	//return err, *un

	db := qmsql.DEFAULTDB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var tmp Salarys
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&tmp, un.ID).Error; err != nil {
		tx.Rollback()
		return err, *un
	}
	if err := tx.Save(un).Error; err != nil {
		tx.Rollback()
		return err, *un
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err, *un
	}
	return err, *un
}

// 根据ID查看单条Salarys
func (un *Salarys) FindById() (err error, reun Salarys) {
	err = qmsql.DEFAULTDB.Where("id = ?", un.ID).First(&reun).Error
	return err, reun
}

func (un *Salarys) FindByIdAndOpenid() (err error, reun Salarys) {
	err = qmsql.DEFAULTDB.Where("id = ? and openid = ?", un.ID, un.Openid).First(&reun).Error
	return err, reun
}

// 分页获取Salarys
func (un *Salarys) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(un, info)
	if err != nil {
		return
	} else {
		var reSalarysList []Salarys
		model := qmsql.DEFAULTDB.Model(info)
		if un.EnterpriseId != 0 {
			model = model.Where("enterprise_id = ?", un.EnterpriseId)
			db = db.Where("enterprise_id = ?", un.EnterpriseId)
		}
		if un.Year != 0 {
			model = model.Where("year = ?", un.Year)
			db = db.Where("year = ?", un.Year)
		}
		if un.Month != 0 {
			model = model.Where("month LIKE ?", un.Month)
			db = db.Where("month LIKE ?", un.Month)
		}
		if un.Name != "" {
			model = model.Where("name LIKE ?", "%"+un.Name+"%")
			db = db.Where("name LIKE ?", "%"+un.Name+"%")
		}
		if un.Gangwei != "" {
			model = model.Where("gangwei LIKE ?", "%"+un.Gangwei+"%")
			db = db.Where("gangwei = ?", "%"+un.Gangwei+"%")
		}

		err = model.Find(&reSalarysList).Count(&total).Error
		if err != nil {
			return err, reSalarysList, total
		} else {
			err = db.Find(&reSalarysList).Error
		}
		return err, reSalarysList, total
	}
}

func (un *Salarys) GetInfoListSearch(openid string, yearL, monthL, yearH, monthH, page, pageSize int) (err error, list interface{}) {
	offset := pageSize * (page - 1)
	var reSalarysList []Salarys
	db := qmsql.DEFAULTDB.Model(un).Select("id,year,month").Where("openid = ?", openid)
	if yearL > 0 && yearH > 0 && monthL > 0 && monthH > 0 {
		db = db.Where("year >= ? and year <= ? and month >= ? and month <= ?", yearL, yearH, monthL, monthH)
	}
	err = db.Order("id desc").Offset(offset).Limit(pageSize).Find(&reSalarysList).Error
	return err, reSalarysList
}
