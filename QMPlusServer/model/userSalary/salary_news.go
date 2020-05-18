// 自动生成模板UserNews
package userSalary

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"

	"github.com/jinzhu/gorm"
)

type UserNews struct {
	gorm.Model
	Title    string `json:"title" gorm:"column:title;comment:'新闻标题'"`
	SubTitle string `json:"sub_title" gorm:"column:sub_title;comment:'新闻副标题'"`
	Content  string `json:"content" gorm:"column:content;comment:'新闻内容'"`
	Order    int    `json:"order" gorm:"column:orders;comment:'排序，越大越靠前'"`
	Type     int    `json:"news_type" gorm:"column:type;comment:'新闻类型，1社保，2个税'"`
	Img      string `json:"img" gorm:"column:img;comment:'新闻图片'"`
	Status   int    `json:"status" gorm:"column:status;comment:'状态，1草稿，2发布';default:1"`
	Count    int    `json:"count" gorm:"column:count;comment:'阅读数'`
}

// 创建UserNews
func (un *UserNews) CreateUserNews() (err error) {
	err = qmsql.DEFAULTDB.Create(un).Error
	return err
}

// 删除UserNews
func (un *UserNews) DeleteUserNews() (err error) {
	err = qmsql.DEFAULTDB.Delete(un).Error
	return err
}

// 更新UserNews
func (un *UserNews) UpdateUserNews() (err error, reun UserNews) {
	err = qmsql.DEFAULTDB.Save(un).Error
	return err, *un
}

func (un *UserNews) UpdateUserNewsCount() (err error) {
	err = qmsql.DEFAULTDB.Model(un).UpdateColumn("count", un.Count).Error
	return err
}

// 根据ID查看单条UserNews
func (un *UserNews) FindById() (err error, reun UserNews) {
	err = qmsql.DEFAULTDB.Where("id = ?", un.ID).First(&reun).Error
	return err, reun
}

// 分页获取UserNews
func (un *UserNews) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(un, info)
	if err != nil {
		return
	} else {
		var reUserNewsList []UserNews
		model := qmsql.DEFAULTDB.Model(info)
		if un.Title != "" {
			model = model.Where("title LIKE ?", "%"+un.Title+"%")
			db = db.Where("title LIKE ?", "%"+un.Title+"%")
		}
		err = model.Find(&reUserNewsList).Count(&total).Error
		if err != nil {
			return err, reUserNewsList, total
		} else {
			err = db.Find(&reUserNewsList).Error
		}
		return err, reUserNewsList, total
	}
}

// 分页获取UserNews根据类型，
func (un *UserNews) GetInfoListByTyp(typ, page, pageSize int) (err error, list interface{}) {

	offset := pageSize * (page - 1)

	var reUserNewsList []UserNews

	qmsql.DEFAULTDB.Model(un).Select("id,title,sub_title,img,count").Where("status = 2").Where("type = ?", typ).Limit(pageSize).Offset(offset).Order("orders desc").Find(&reUserNewsList)

	return nil, reUserNewsList

}
