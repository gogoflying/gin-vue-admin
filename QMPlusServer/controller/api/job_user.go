package api

import (
	"fmt"
	"encoding/json"
	"net/http"
	"time"

	"gin-vue-admin/config"
	"gin-vue-admin/middleware"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

var appid string = config.GinVueAdminconfig.WeiXin.Appid
var appSecret string = config.GinVueAdminconfig.WeiXin.AppSecret

func CreateUsers(c *gin.Context) {
	var users userJobs.Users
	_ = c.ShouldBindJSON(&users)
	err := users.CreateUsers()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

func DeleteUsers(c *gin.Context) {
	var users userJobs.Users
	_ = c.ShouldBindJSON(&users)
	err := users.DeleteUsers()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags Users
// @Summary 更新Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.Users true "更新Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /users/updateUsers [post]
func UpdateUsers(c *gin.Context) {
	var users userJobs.Users
	_ = c.ShouldBindJSON(&users)
	err, reusers := users.UpdateUsers()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reusers": reusers,
		})
	}
}

// @Tags Users
// @Summary 用id查询Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.Users true "用id查询Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /users/findUsers [post]
func FindUsers(c *gin.Context) {
	var users userJobs.Users
	_ = c.ShouldBindJSON(&users)
	err, reusers := users.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reusers": reusers,
		})
	}
}

func FindUsersByOpenid(c *gin.Context) {
	var users userJobs.Users
	_ = c.ShouldBindJSON(&users)
	err, user, edus, works, base, dream := users.FindByOpenId()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"user":  user,
			"edus":  edus,
			"works": works,
			"base":  base,
			"dream": dream,
		})
	}
}


func GetUsersList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(userJobs.Users).GetInfoList(pageInfo)
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

func JobUserLogin(c *gin.Context) {
	var loginInfo userJobs.UserLoginInfo
	_ = c.ShouldBindJSON(&loginInfo)
	if len(loginInfo.Code) == 0{
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败"), gin.H{})
		return
	}
	openid,err:= sendWxAuthAPI(loginInfo.Code)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
		return
	} 
	
	u:= userJobs.Users{
		Openid:openid,
		Status:1,
	}
	_ = c.ShouldBindJSON(&u)
	err= u.CreateUsers()
	if err != nil{
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
		return
	}
	tokenNext_wx(c,u)
	fmt.Printf("get openId:%v\n", openid) 
}

func sendWxAuthAPI(code string) (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",appid,appSecret,code)
	//fmt.Printf("get url:%v\n", url)
    resp, err := http.DefaultClient.Get(url)
    if err != nil {
        return "", err
    }
    var wxMap map[string]string
    err = json.NewDecoder(resp.Body).Decode(&wxMap)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    return wxMap["openid"], nil
}

//登录以后签发jwt
func tokenNext_wx(c *gin.Context, user userJobs.Users) {
	j := &middleware.JWT_wx{
		[]byte(config.GinVueAdminconfig.JWT.SigningKey), // 唯一签名
	}
	//err := user.FindByUserId()
    
	clams := middleware.CustomClaims_wx{
		Open_Id:         user.Openid,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),  // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60), // 过期时间 一个小时
			Issuer:    "qmPlus",                         //签名的发行者
		},
	}
	token, err := j.CreateToken_wx(clams)
	if err != nil {
		servers.ReportFormat(c, false, "获取token失败", gin.H{})
	} else {
		servers.ReportFormat(c, true, "登录成功", gin.H{"openid": user.Openid,"token": token, "expiresAt": clams.StandardClaims.ExpiresAt * 1000})
	}
}


func SendUserMobile(c *gin.Context) {
	var loginInfo userJobs.UserLoginInfo
	_ = c.ShouldBindJSON(&loginInfo)
	if len(loginInfo.Openid) == 0 || len(loginInfo.Mobile) != 11 {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败"), gin.H{})
		return
	}
	
	u:= userJobs.Users{
		Openid:loginInfo.Openid,
		Mobile:loginInfo.Mobile,
	}
	_ = c.ShouldBindJSON(&u)
	err,reusers := u.FindById()
	if err != nil{
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
		return
	}
	if reusers.Openid == ""{
		servers.ReportFormat(c, false, fmt.Sprintf("非法用户：%v", err), gin.H{})
		return 
	}
	err,reusers = u.UpdateUsers()
	if err != nil{
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "更新成功", gin.H{
			"reusers": reusers,
		})
}