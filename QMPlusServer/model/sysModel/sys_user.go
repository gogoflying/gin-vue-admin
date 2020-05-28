package sysModel

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/tools"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	gorm.Model
	UUID        uuid.UUID    `json:"uuid"`
	Username    string       `json:"userName"`
	Password    string       `json:"-"`
	NickName    string       `json:"nickName" gorm:"default:'BDAPlusUser'"`
	HeaderImg   string       `json:"headerImg" gorm:"default:'https://bda-edu-hr.oss-cn-beijing.aliyuncs.com/002.png'"`
	Authority   SysAuthority `json:"authority" gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId"`
	AuthorityId string       `json:"authorityId" gorm:"default:9528"`
	Email       string       `json:"email" gorm:"email"`
}

//type Propertie struct {
//	gorm.Model
//}

//注册接口model方法
func (u *SysUser) Register() (err error, userInter *SysUser) {
	var user SysUser
	//判断用户名是否注册
	notResigt := qmsql.DEFAULTDB.Where("username = ?", u.Username).First(&user).RecordNotFound()
	//notResigt为false表明读取到了 不能注册
	if !notResigt {
		return errors.New("用户名已注册"), userInter
	} else {
		// 否则 附加uuid 密码md5简单加密 注册
		u.Password = tools.MD5V([]byte(u.Password))
		u.UUID = uuid.NewV4()
		err = qmsql.DEFAULTDB.Create(u).Error
	}
	return err, u
}

// 根据ID查看单条User
func (u *SysUser) FindById() (err error, user SysUser) {
	err = qmsql.DEFAULTDB.Preload("Authority").Where("id = ?", u.ID).First(&user).Error
	return err, user
}

// 根据ID查看单条User
func (u *SysUser) FindSysUserById(id uint) (err error, user SysUser) {
	qmsql.DEFAULTDB.LogMode(true)
	err = qmsql.DEFAULTDB.Preload("Authority").Where("id = ?", id).First(&user).Error
	return err, user
}

// 更新User
func (u *SysUser) ResetPassword() (err error, ru SysUser) {
	var user SysUser
	password := tools.MD5V([]byte("123456"))
	err = qmsql.DEFAULTDB.Where("username = ?", u.Username).First(&user).Update("password", password).Error
	return err, *u
}

// 删除User
func (u *SysUser) DeleteUser() (err error) {
	err = qmsql.DEFAULTDB.Delete(u).Error
	return err
}

//修改用户密码
func (u *SysUser) ChangePassword(newPassword string) (err error, userInter *SysUser) {
	var user SysUser
	//后期修改jwt+password模式
	u.Password = tools.MD5V([]byte(u.Password))
	err = qmsql.DEFAULTDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", tools.MD5V([]byte(newPassword))).Error
	return err, u
}

//用户更新接口
func (u *SysUser) SetUserAuthority(uuid uuid.UUID, AuthorityId string) (err error) {
	err = qmsql.DEFAULTDB.Where("uuid = ?", uuid).First(&SysUser{}).Update("authority_id", AuthorityId).Error
	return err
}

func (u *SysUser) GetByUsername() (err error, userInter *SysUser) {
	var user SysUser
	err = qmsql.DEFAULTDB.Model(u).Where("username = ? and authority_id = ?", u.Username, u.AuthorityId).First(&user).Error
	return err, &user
}

// 更新User
func (u *SysUser) ResetPasswordForget() (err error, ru SysUser) {
	var user SysUser
	password := tools.MD5V([]byte("12345678"))
	err = qmsql.DEFAULTDB.Where("username = ? and authority_id = ?", u.Username, u.AuthorityId).First(&user).Update("password", password).Error
	return err, *u
}

//用户登录
func (u *SysUser) Login() (err error, userInter *SysUser) {
	var user SysUser
	u.Password = tools.MD5V([]byte(u.Password))
	err = qmsql.DEFAULTDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	if err != nil {
		return err, &user
	}
	err = qmsql.DEFAULTDB.Where("authority_id = ?", user.AuthorityId).First(&user.Authority).Error
	return err, &user
}

// 用户头像上传更新地址
func (u *SysUser) UploadHeaderImg(uuid uuid.UUID, filePath string) (err error, userInter *SysUser) {
	var user SysUser
	err = qmsql.DEFAULTDB.Where("uuid = ?", uuid).First(&user).Update("header_img", filePath).First(&user).Error
	return err, &user
}

// 分页获取数据  需要分页实现这个接口即可
func (u *SysUser) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(u, info)
	if err != nil {
		return
	} else {
		var userList []SysUser
		err = db.Preload("Authority").Find(&userList).Error
		return err, userList, total
	}
}
