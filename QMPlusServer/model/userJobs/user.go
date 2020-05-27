// 自动生成模板Users
package userJobs

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userCity"

	"github.com/jinzhu/gorm"
)

type DecodeMobile struct {
	EncrypData string `json:"encrypdata"`
	IvData     string `json:"ivdata"`
	SessionKey string `json:"sessionkey"`
	Openid     string `json:"openid"`
}

type UserLoginInfo struct {
	Code   string `json:"code"`
	Openid string `json:"openid"`
	Mobile string `json:"mobile"`
}

type Users struct {
	gorm.Model
	Openid   string `json:"openid" gorm:"column:openid;comment:'用户标识码'"`
	Mobile   string `json:"mobile" gorm:"column:mobile;comment:'手机号'"`
	IsResume int    `json:"is_resume" gorm:"column:is_resume;comment:'是否有简历，1表示有，0没有'"`
	Status   int    `json:"status" gorm:"column:status;comment:'用户状态，1表示正常，其他非正常，如禁用之类的'"`
}

// 创建Users
func (users *Users) CreateUsers() (err error) {
	err = qmsql.DEFAULTDB.Create(users).Error
	return err
}

// 删除Users
func (users *Users) DeleteUsers() (err error) {
	err = qmsql.DEFAULTDB.Delete(users).Error
	return err
}

// 更新Users
func (users *Users) UpdateUsers() (err error, reusers Users) {
	err = qmsql.DEFAULTDB.Model(users).Where("openid = ? and status = 1", users.Openid).Update("is_resume", users.IsResume).Error
	return err, *users
}
func (users *Users) UpdateUsersMobile() (err error, reusers Users) {
	err = qmsql.DEFAULTDB.Model(users).Where("openid = ?", users.Openid).Update("mobile", users.Mobile).Error
	return err, *users
}

// 根据ID查看单条Users
func (users *Users) FindById() (err error, reusers Users) {
	err = qmsql.DEFAULTDB.Where("openid = ?", users.Openid).First(&reusers).Error
	return err, reusers

}

// 根据ID查看单条Users
func (users *Users) FindByJobUserId() (err error, reusers Users) {
	err = qmsql.DEFAULTDB.Where("id = ?", users.ID).First(&reusers).Error
	return err, reusers
}

// 根据ID查看单条Users
func (users *Users) GetByOpenId() (err error, reusers Users) {
	err = qmsql.DEFAULTDB.Model(users).Where("openid = ? and status = 1", users.Openid).First(&reusers).Error
	return
}

// 根据ID查看单条Users
func (users *Users) FindByOpenId() (err error, reusers Users, edus, wks, base, dream interface{}) {
	err = qmsql.DEFAULTDB.Model(users).Where("openid = ? and status = 1", users.Openid).First(&reusers).Error
	if err != nil {
		return
	}
	if reusers.IsResume == 1 {
		_, edus, _ = new(UserEducation).GetInfoListByOpenid(users.Openid, 1, 20)
		_, wks, _ = new(UserWork).GetInfoListOpenId(users.Openid, 1, 20)
		_, base = (&UserBaseInfo{Openid: users.Openid}).FindByOpenid()
		_, dream = (&UserDream{Openid: users.Openid}).FindByOpenid()
	}
	return

}

func (users *Users) FindResumeInfoByOpenId() (err error, reusers Users, edus, wks, base, dream interface{}) {
	err = qmsql.DEFAULTDB.Model(users).Where("openid = ?", users.Openid).First(&reusers).Error
	if err != nil {
		return
	}
	_, edus, _ = new(UserEducation).GetInfoListByOpenid(users.Openid, 1, 20)
	_, wks, _ = new(UserWork).GetInfoListOpenId(users.Openid, 1, 20)
	err1, baseTemp := (&UserBaseInfo{Openid: users.Openid}).FindByOpenid()
	if err1 == nil {
		base = baseTemp
	}
	err2, dreamTemp := (&UserDream{Openid: users.Openid}).FindByOpenid()
	if err2 == nil {
		dream = dreamTemp
	}
	return
}

func (users *Users) GetEduLevels() (err error, list interface{}) {
	var els []EduLevel
	err = qmsql.DEFAULTDB.Select("id,name").Find(&els).Error
	return err, els
}

func (users *Users) GetJobWorkTypes() (err error, list interface{}) {
	var wts []JobWorkType
	err = qmsql.DEFAULTDB.Select("id,name").Find(&wts).Error
	return err, wts
}

func (users *Users) GetAllInfoOption() (err error, list1, list2, list3, list4, list5, list6 interface{}) {
	var wt []JobWorkType
	var el []EduLevel
	var jwe []JobWorkExpire
	var js []JobSalary
	var jdt []JobDutyTime
	var citys []userCity.Cityname

	err = qmsql.DEFAULTDB.Select("id,name").Find(&wt).Error
	if err != nil {
		return err, wt, el, jwe, js, jdt, citys
	}

	err = qmsql.DEFAULTDB.Select("id,name").Find(&el).Error
	if err != nil {
		return err, wt, el, jwe, js, jdt, citys
	}

	err = qmsql.DEFAULTDB.Select("id,name").Find(&jwe).Error
	if err != nil {
		return err, wt, el, jwe, js, jdt, citys
	}

	err = qmsql.DEFAULTDB.Select("id,name").Find(&js).Error
	if err != nil {
		return err, wt, el, jwe, js, jdt, citys
	}

	err = qmsql.DEFAULTDB.Select("id,name").Find(&jdt).Error
	if err != nil {
		return err, wt, el, jwe, js, jdt, citys
	}

	err = qmsql.DEFAULTDB.Select("id,name").Find(&citys).Error
	if err != nil {
		return err, wt, el, jwe, js, jdt, citys
	}

	return err, wt, el, jwe, js, jdt, citys
}

// 分页获取Users
func (users *Users) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(users, info)
	if err != nil {
		return
	} else {
		var reUsersList []Users
		model := qmsql.DEFAULTDB.Model(info)
		if users.Mobile != "" {
			model = model.Where("mobile LIKE ?", "%"+users.Mobile+"%")
			db = db.Where("mobile LIKE ?", "%"+users.Mobile+"%")
		}
		err = model.Find(&reUsersList).Count(&total).Error
		if err != nil {
			return err, reUsersList, total
		} else {
			err = db.Find(&reUsersList).Error
		}

		return err, reUsersList, total
	}
}

type ResumeUserView struct {
	Users
	Mobile         string `json:"contact"`
	UserName       string `json:"userName"`
	HasAvatarUrl   *int   `json:"hasAvatarurl"`
	Genderindex    *int   `json:"genderindex"`
	EdulevelIndex  *int   `json:"edulevelindex"`
	WorksYearindex *int   `json:"worksYearindex"`

	Dreamposi     string `json:"dreamposi" `
	WorkTypeindex *int   `json:"workTypeindex"`
	Cityindex     *int   `json:"cityindex"`
	Salaryindex   *int   `json:"salaryindex"`
	DutyTimeindex *int   `json:"dutyTimeindex"`
}

// 分页获取Users
func (users *Users) GetInfoListNew(ruv ResumeUserView, info modelInterface.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := qmsql.DEFAULTDB
	db.LogMode(true)
	table := db.Select([]string{
		"users.*",
		"user_base_infos.mobile",
		"user_base_infos.user_name",
		"user_base_infos.avatarUrl",
		"user_base_infos.genderindex",
		"user_base_infos.edulevelindex",
		"user_base_infos.worksYearindex",
		"user_dreams.dreamposi",
		"user_dreams.workTypeindex",
		"user_dreams.cityindex",
		"user_dreams.salaryindex",
		"user_dreams.dutyTimeindex",
	}).Table("users")
	table = table.Joins("left join user_base_infos on users.openid=user_base_infos.openid")
	table = table.Joins("left join user_dreams on users.openid=user_dreams.openid")

	var reUsersViewList []ResumeUserView
	if ruv.Mobile != "" {
		table = table.Where("user_base_infos.mobile LIKE ?", "%"+ruv.Mobile+"%")
	}
	if ruv.UserName != "" {
		table = table.Where("user_base_infos.user_name LIKE ?", "%"+ruv.UserName+"%")
	}
	if ruv.HasAvatarUrl != nil {
		if *ruv.HasAvatarUrl == 0 {
			table = table.Where("LENGTH(trim(user_base_infos.avatarUrl))<1")
		} else {
			table = table.Where("LENGTH(trim(user_base_infos.avatarUrl))>0")
		}
	}
	if ruv.Genderindex != nil {
		table = table.Where("user_base_infos.genderindex = ?", *ruv.Genderindex)
	}
	if ruv.EdulevelIndex != nil {
		table = table.Where("user_base_infos.edulevelindex = ?", *ruv.EdulevelIndex)
	}
	if ruv.WorksYearindex != nil {
		table = table.Where("user_base_infos.worksYearindex = ?", *ruv.WorksYearindex)
	}
	if ruv.Dreamposi != "" {
		table = table.Where("user_dreams.dreamposi LIKE ?", "%"+ruv.Dreamposi+"%")
	}
	if ruv.WorkTypeindex != nil {
		table = table.Where("user_dreams.workTypeindex = ?", *ruv.WorkTypeindex)
	}
	if ruv.Cityindex != nil {
		table = table.Where("user_dreams.cityindex = ?", *ruv.Cityindex)
	}
	if ruv.Salaryindex != nil {
		table = table.Where("user_dreams.salaryindex = ?", *ruv.Salaryindex)
	}
	if ruv.DutyTimeindex != nil {
		table = table.Where("user_dreams.dutyTimeindex = ?", *ruv.DutyTimeindex)
	}

	err = table.Find(&reUsersViewList).Count(&total).Error
	if err != nil {
		return err, reUsersViewList, total
	} else {
		table = table.Limit(limit).Offset(offset)
		err = table.Find(&reUsersViewList).Error
	}
	return err, reUsersViewList, total
}

// 分页获取Users投递记录
func (users *Users) GetResumeList(info modelInterface.PageInfo, openid string, dataRange []string, resumeStatus int) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := qmsql.DEFAULTDB.Limit(limit).Offset(offset).Order("id desc")

	if openid == "" {
		return
	} else {
		var resumeList []ResumeStatus
		model := qmsql.DEFAULTDB.Model(info)
		model = model.Where("open_id = ?", openid)
		db = db.Where("open_id = ?", openid)

		if len(dataRange) == 2 {
			model = model.Where("created_at > ? and created_at < ?", dataRange[0], dataRange[1])
			db = db.Where("created_at > ? and created_at < ?", dataRange[0], dataRange[1])
		}

		if resumeStatus != -1 {
			model = model.Where("resume_status = ?", resumeStatus)
			db = db.Where("resume_status = ?", resumeStatus)
		}

		err = model.Preload("JobInfo").Find(&resumeList).Count(&total).Error
		if err != nil {
			return err, resumeList, total
		} else {
			err = db.Preload("JobInfo").Find(&resumeList).Error
		}

		return err, resumeList, total
	}
}
