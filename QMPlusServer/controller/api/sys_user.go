package api

import (
	"fmt"
	"gin-vue-admin/config"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/sysModel"
	"gin-vue-admin/model/userJobs"
	"math/rand"
	"mime/multipart"
	"time"

	"github.com/dchest/captcha"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
)

var (
	USER_HEADER_IMG_PATH string = "http://qmplusimg.henrongyi.top"
	//USER_HEADER_BUCKET   string = "qm-plus-img"
	USER_HEADER_BUCKET string = "bda-edu-hr"
)

type RegisterAndLoginStuct struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

type RegestStuct struct {
	Username       string `json:"userName"`
	Password       string `json:"passWord"`
	EnterPriseName string `json:"enterprise_name"`
	NickName       string `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg      string `json:"headerImg" gorm:"default:'https://bda-edu-hr.oss-cn-beijing.aliyuncs.com/002.png'"`
	AuthorityId    string `json:"authorityId" gorm:"default:9528"`
	Email          string `json:"email" gorm:"email"`
}

type ForgetPass struct {
	Username  string `json:"userName"`
	Code      string `json:"code" `
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
	RoleId    string `json:"roleId"`
}

// @Tags Base
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body sysModel.SysUser true "用户注册接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /base/register [post]
func Register(c *gin.Context) {
	var R RegestStuct
	_ = c.ShouldBindJSON(&R)
	if R.AuthorityId == "" {
		R.AuthorityId = "9528"
	}
	user := &sysModel.SysUser{Username: R.Username, NickName: R.NickName, Password: R.Password, HeaderImg: R.HeaderImg, AuthorityId: R.AuthorityId, Email: R.Email}
	err, user := user.Register()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("%v", err), gin.H{
			"user": user,
		})
	} else {
		if R.EnterPriseName != "" {
			var info userJobs.EnterpriseInfo
			info.EnterPriseName = R.EnterPriseName
			err, ep := info.SelectAndCreateEnterpriseInfo()
			if err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("%v", err), gin.H{
					"user": user,
				})
			} else {
				auth := userJobs.UserAuth{
					EnterPriseName: ep.EnterPriseName,
					EnterPriseId:   ep.ID,
					Username:       user.Username,
					UserId:         user.ID,
				}
				err = auth.CreateUserAuth()
				if err != nil {
					servers.ReportFormat(c, false, fmt.Sprintf("%v", err), gin.H{
						"user": user,
					})
				}
			}
		}
		servers.ReportFormat(c, true, "创建成功", gin.H{
			"user": user,
		})
	}
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body api.RegisterAndLoginStuct true "用户登录接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var L RegisterAndLoginStuct
	_ = c.ShouldBindJSON(&L)
	if captcha.VerifyString(L.CaptchaId, L.Captcha) {
		U := &sysModel.SysUser{Username: L.Username, Password: L.Password}
		if err, user := U.Login(); err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("用户名密码错误或%v", err), gin.H{})
		} else {
			tokenNext(c, *user)
		}
	} else {
		servers.ReportFormat(c, false, "验证码错误", gin.H{})
	}

}

func VerifyEmailCode(c *gin.Context) {
	var req ForgetPass
	_ = c.ShouldBindJSON(&req)
	code := servers.GetMap(req.Username)
	if code == nil {
		servers.ReportFormat(c, false, "邮箱验证码失效或者尚未发送", gin.H{})
		return
	}
	if code.(string) != req.Code {
		servers.ReportFormat(c, false, "邮箱验证码错误", gin.H{})
		return
	}
	U := &sysModel.SysUser{Username: req.Username, AuthorityId: req.RoleId}
	if err, _ := U.ResetPasswordForget(); err != nil {
		servers.ReportFormat(c, false, "重置密码失败", gin.H{})
		return
	}
	servers.ReportFormat(c, true, "密码重置为：12345678, 请登录后及时修改！", gin.H{})
}

func SendEmailForget(c *gin.Context) {
	var req ForgetPass
	_ = c.ShouldBindJSON(&req)
	if captcha.VerifyString(req.CaptchaId, req.Captcha) {
		U := &sysModel.SysUser{Username: req.Username, AuthorityId: req.RoleId}
		err, user := U.GetByUsername()
		if err != nil {
			servers.ReportFormat(c, false, "用户不存在", gin.H{})
			return
		}
		var s string
		for i := 0; i < 6; i++ {
			rand.Seed(time.Now().UnixNano())
			s = s + fmt.Sprintf("%v", rand.Intn(10))
		}
		if err := servers.SendEmail(user.Email, "密码找回", "你正在使用邮箱找回密码功能，本次邮箱验证码为： "+s); err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("发送邮箱验证码失败", err), gin.H{})
			return
		}
		servers.AddMap(req.Username, s)
		servers.ReportFormat(c, true, "发送成功，请登录邮箱获取验证码", gin.H{})
	} else {
		servers.ReportFormat(c, false, "验证码错误", gin.H{})
	}
}

//登录以后签发jwt
func tokenNext(c *gin.Context, user sysModel.SysUser) {
	j := &middleware.JWT{
		[]byte(config.GinVueAdminconfig.JWT.SigningKey), // 唯一签名
	}
	enPriseId := uint(0)
	ua := &userJobs.UserAuth{
		UserId: user.ID,
	}
	err := ua.FindByUserId()
	if err == nil {
		enPriseId = ua.EnterPriseId
	}
	clams := middleware.CustomClaims{
		UUID:         user.UUID,
		ID:           user.ID,
		EnterPriseId: enPriseId,
		NickName:     user.NickName,
		AuthorityId:  user.AuthorityId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),     // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24), // 过期时间 24个小时
			Issuer:    "qmPlus",                            //签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		servers.ReportFormat(c, false, "获取token失败", gin.H{})
	} else {
		if config.GinVueAdminconfig.System.UseMultipoint {
			var loginJwt sysModel.JwtBlacklist
			loginJwt.Jwt = token
			err, jwtStr := loginJwt.GetRedisJWT(user.Username)
			if err == redis.Nil {
				err2 := loginJwt.SetRedisJWT(user.Username)
				if err2 != nil {
					servers.ReportFormat(c, false, "设置登录状态失败", gin.H{})
				} else {
					servers.ReportFormat(c, true, "登录成功", gin.H{"user": user, "enPriseId": enPriseId, "token": token, "expiresAt": clams.StandardClaims.ExpiresAt * 1000})
				}
			} else if err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("%v", err), gin.H{})
			} else {
				var blackjWT sysModel.JwtBlacklist
				blackjWT.Jwt = jwtStr
				err3 := blackjWT.JsonInBlacklist()
				if err3 != nil {
					servers.ReportFormat(c, false, "jwt作废失败", gin.H{})
				} else {
					err2 := loginJwt.SetRedisJWT(user.Username)
					if err2 != nil {
						servers.ReportFormat(c, false, "设置登录状态失败", gin.H{})
					} else {
						servers.ReportFormat(c, true, "登录成功", gin.H{"user": user, "enPriseId": enPriseId, "token": token, "expiresAt": clams.StandardClaims.ExpiresAt * 1000})
					}
				}
			}
		} else {
			servers.ReportFormat(c, true, "登录成功", gin.H{"user": user, "enPriseId": enPriseId, "token": token, "expiresAt": clams.StandardClaims.ExpiresAt * 1000})
		}
	}
}

type ChangePasswordStutrc struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// @Tags SysUser
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body api.ChangePasswordStutrc true "用户修改密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [post]
func ChangePassword(c *gin.Context) {
	var params ChangePasswordStutrc
	_ = c.ShouldBindJSON(&params)
	U := &sysModel.SysUser{Username: params.Username, Password: params.Password}
	if err, _ := U.ChangePassword(params.NewPassword); err != nil {
		servers.ReportFormat(c, false, "修改失败，请检查用户名密码", gin.H{})
	} else {
		servers.ReportFormat(c, true, "修改成功", gin.H{})
	}
}

type UserHeaderImg struct {
	HeaderImg multipart.File `json:"headerImg"`
}

// @Tags SysUser
// @Summary 用户上传头像
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param headerImg formData file true "用户上传头像"
// @Param username formData string true "用户上传头像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /user/uploadHeaderImg [post]
func UploadHeaderImg(c *gin.Context) {
	claims, _ := c.Get("claims")
	//获取头像文件
	// 这里我们通过断言获取 claims内的所有内容
	waitUse := claims.(*middleware.CustomClaims)
	uuid := waitUse.UUID
	_, header, err := c.Request.FormFile("headerImg")
	//便于找到用户 以后从jwt中取
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("上传文件失败，%v", err), gin.H{})
	} else {
		//文件上传后拿到文件路径
		err, filePath, _ := servers.Upload(header, USER_HEADER_BUCKET, USER_HEADER_IMG_PATH)
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("接收返回值失败，%v", err), gin.H{})
		} else {
			//修改数据库后得到修改后的user并且返回供前端使用
			err, user := new(sysModel.SysUser).UploadHeaderImg(uuid, filePath)
			if err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("修改数据库链接失败，%v", err), gin.H{})
			} else {
				servers.ReportFormat(c, true, "上传成功", gin.H{"user": user})
			}
		}
	}
}

func ResetPassword(c *gin.Context) {
	var ur sysModel.SysUser
	_ = c.ShouldBindJSON(&ur)
	err, _ := ur.ResetPassword()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("重置密码失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "重置密码成功", gin.H{})
	}
}

func DeleteUser(c *gin.Context) {
	var ur sysModel.SysUser
	_ = c.ShouldBindJSON(&ur)
	err := ur.DeleteUser()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

func FindUserById(c *gin.Context) {
	var rs sysModel.SysUser
	_ = c.ShouldBindJSON(&rs)
	err, rers := rs.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"user": rers,
		})
	}
}

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func GetUserList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(sysModel.SysUser).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"userList": list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})
	}
}

type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`
	AuthorityId string    `json:"authorityId"`
}

// @Tags SysUser
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.SetUserAuth true "设置用户权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func SetUserAuthority(c *gin.Context) {
	var sua SetUserAuth
	_ = c.ShouldBindJSON(&sua)
	err := new(sysModel.SysUser).SetUserAuthority(sua.UUID, sua.AuthorityId)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("修改失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "修改成功", gin.H{})
	}
}
