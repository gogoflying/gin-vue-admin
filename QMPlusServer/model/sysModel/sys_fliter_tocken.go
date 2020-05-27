// 自动生成模板WxFliterTocken
package sysModel

import (
	"fmt"
	"gin-vue-admin/init/qmsql"

	"github.com/jinzhu/gorm"
)

type WxFliterTocken struct {
	gorm.Model
	AppName string `json:"appname" gorm:"column:appName"`
	Tocken  string `json:"tocken" gorm:"column:tocken"`
}

// 创建SalaryContract
func (wft *WxFliterTocken) CreateWxFliterTocken() (err error) {
	err = qmsql.DEFAULTDB.Create(wft).Error
	return err
}

// 删除SalaryContract
func (wft *WxFliterTocken) DeleteWxFliterTocken() (err error) {
	err = qmsql.DEFAULTDB.Delete(wft).Error
	return err
}

func (wft *WxFliterTocken) UpdateWxFliterTocken() (err error, reun WxFliterTocken) {
	err = qmsql.DEFAULTDB.Model(wft).Where("appname = ? ", wft.AppName).Updates(wft).Error
	return err, *wft
}

func (wft *WxFliterTocken) GetWxFliterTocken(id uint) (err error, rewft WxFliterTocken) {
	err = qmsql.DEFAULTDB.Where("appname = ?", wft.AppName).First(&rewft).Error
	return err, rewft
}

func (wft *WxFliterTocken) FindById() (err error, rewft WxFliterTocken) {
	err = qmsql.DEFAULTDB.Where("appname = ?", wft.AppName).First(&rewft).Error
	return err, rewft
}

func (wft *WxFliterTocken) UpdateFilterTocken(appname string, tocken string) {
	if appname == "" || tocken == "" {
		fmt.Printf("UpdateFilterTocken param empty")
		return
	}
	nTocken := WxFliterTocken{
		AppName: appname,
	}
	err, cur := nTocken.FindById()
	if cur.AppName == "" {
		nTocken.Tocken = tocken
		err = nTocken.CreateWxFliterTocken()
		if err != nil {
			fmt.Printf("UpdateFilterTocken insert err :%v", err)
			return
		}
	} else {
		cur.Tocken = tocken
		err, _ := cur.UpdateWxFliterTocken()
		if err != nil {
			fmt.Printf("UpdateFilterTocken update err :%v", err)
			return
		}
	}
}

func (wft *WxFliterTocken) GetNewFilterTocken(appname string) string {
	if appname == "" {
		fmt.Printf("UpdateFilterTocken param empty")
		return ""
	}
	nTocken := WxFliterTocken{
		AppName: appname,
	}
	err, cur := nTocken.FindById()
	if err != nil {
		fmt.Printf("UpdateFilterTocken FindById err :%v", err)
		return ""
	}
	return cur.Tocken
}

//持久化到数据库  引入自定义规则
func GetWxFliterTockenHndle(appname string, tocken string) *WxFliterTocken {
	return &WxFliterTocken{
		AppName: appname,
		Tocken:  tocken,
	}
}
