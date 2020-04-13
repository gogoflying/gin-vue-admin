package userJobs

import (
	"gin-vue-admin/init/qmsql"
	//	"gin-vue-admin/model/modelInterface"
	"testing"

	"github.com/jinzhu/gorm"
)

var DEFAULTDB *gorm.DB

//初始化数据库并产生数据库全局变量
func Init() {
	if db, err := gorm.Open("mysql", "root"+":"+""+"@("+"127.0.0.1"+")/"+"qmplus"+"?"+"charset=utf8\u0026parseTime=True\u0026loc=Local"); err != nil {
		panic(err)
	} else {
		qmsql.DEFAULTDB = db
		qmsql.DEFAULTDB.AutoMigrate(Joblist{})
	}
	return
}

func TestGetInfoListSearch(t *testing.T) {
	Init()

	ts := new(Joblist)

	// err, list, total := ts.GetInfoListSearch("前台", 0, 0, 0, modelInterface.PageInfo{Page: 1, PageSize: 10})

	// t.Log(err, list, total)

	err, list := ts.GetInfoListSearchSimilar([]uint{6, 7}, "前台", 1, 1, 1, 1, 1, 1, 10)
	t.Log(err, list)
}
