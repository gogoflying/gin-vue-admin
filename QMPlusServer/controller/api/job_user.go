package api

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gin-vue-admin/config"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/userJobs"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func GetWorkTypes(c *gin.Context) {
	var users userJobs.Users
	err, list := users.GetJobWorkTypes()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"list": list,
		})
	}
}

func GetEduLevels(c *gin.Context) {
	var users userJobs.Users
	err, list := users.GetEduLevels()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"list": list,
		})
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

// @Tags Users
// @Summary 用id查询Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.Users true "用id查询Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /users/FindJobUserById [post]
func FindJobUserById(c *gin.Context) {
	var users userJobs.Users
	_ = c.ShouldBindJSON(&users)
	err, reusers := users.FindByJobUserId()
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

func GetUserOptions(c *gin.Context) {
	err, wt, el, jwe, js, jdt, citys := new(userJobs.Users).GetAllInfoOption()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"wt":       wt,
			"el":       el,
			"jwe":      jwe,
			"js":       js,
			"jdt":      jdt,
			"cityinfo": citys,
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
	if len(loginInfo.Code) == 0 {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败"), gin.H{})
		return
	}

	//==============product==================

	// openid, session_key, err := sendWxAuthAPI(loginInfo.Code)
	// if err != nil {
	// 	servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	// 	return
	// }

	// u := userJobs.Users{
	// 	Openid: openid,
	// 	Status: 1,
	// }
	// err, reuser := u.GetByOpenId()
	// if err != nil {
	// 	_ = c.ShouldBindJSON(&u)
	// 	err = u.CreateUsers()
	// 	if err != nil {
	// 		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	// 		return
	// 	}
	// }
	// var isMobile bool = false
	// if len(reuser.Mobile) > 0 {
	// 	isMobile = true
	// }
	// tokenNext_wx(c, openid, session_key, isMobile)
	// fmt.Printf("get openId:%v\n", openid)
	//========================================

	//==============develop===================
	u := userJobs.Users{
		Openid: loginInfo.Code,
		Status: 1,
	}
	err, _ := u.GetByOpenId()
	if err != nil {
		_ = c.ShouldBindJSON(&u)
		err = u.CreateUsers()
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
			return
		}
	}
	var isMobile bool = true

	tokenNext_wx(c, u.Openid, "", isMobile)
	//========================================

}

func sendWxAuthAPI(code string) (string, string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appid, appSecret, code)
	//fmt.Printf("get url:%v\n", url)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return "", "", err
	}
	var wxMap map[string]string
	err = json.NewDecoder(resp.Body).Decode(&wxMap)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	fmt.Printf("session key:%v\n", wxMap["session_key"])

	return wxMap["openid"], wxMap["session_key"], nil
}

//登录以后签发jwt
func tokenNext_wx(c *gin.Context, openid, session_key string, isMobile bool) {
	j := &middleware.JWT_wx{
		[]byte(config.GinVueAdminconfig.JWT.SigningKey), // 唯一签名
	}
	//err := user.FindByUserId()

	clams := middleware.CustomClaims_wx{
		Open_Id:     openid,
		Session_key: session_key,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),     // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24), // 过期时间 24个小时
			Issuer:    "qmPlus",                            //签名的发行者
		},
	}
	token, err := j.CreateToken_wx(clams)
	if err != nil {
		servers.ReportFormat(c, false, "获取token失败", gin.H{})
	} else {
		servers.ReportFormat(c, true, "登录成功", gin.H{"openid": openid, "isPhone": isMobile, "session_key": session_key, "token": token, "expiresAt": clams.StandardClaims.ExpiresAt * 1000})
	}
}

func updateUserMobile(openId, mobileNum string) error {

	if len(openId) == 0 || len(mobileNum) != 11 {
		//servers.ReportFormat(c, false, fmt.Sprintf("创建失败"), gin.H{})
		return fmt.Errorf("param empty ")
	}

	u := userJobs.Users{
		Openid: openId,
		Mobile: mobileNum,
	}

	err, _ := u.UpdateUsersMobile()
	if err != nil {
		//servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
		return err
	}
	return nil
}

func DecodeMobile(c *gin.Context) {
	var enMobile userJobs.DecodeMobile
	_ = c.ShouldBindJSON(&enMobile)

	if len(enMobile.SessionKey) == 0 || len(enMobile.Openid) == 0 {
		servers.ReportFormat(c, false, fmt.Sprintf("SessionKey empty err"), gin.H{})
		return
	}
	fmt.Printf("responst :%v\n", enMobile)
	src, err := Dncrypt(enMobile.EncrypData, enMobile.SessionKey, enMobile.IvData)
	fmt.Println(err)
	var s = map[string]interface{}{}
	json.Unmarshal([]byte(src), &s)
	fmt.Printf("photo num is: %+v\n", s["phoneNumber"])
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("Dncrypt ：%v", err), gin.H{})
		return
	}
	err = updateUserMobile(enMobile.Openid, s["phoneNumber"].(string))
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("updateUserMobile err :%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "更新成功", gin.H{
		"data": "ok",
	})
}

func Dncrypt(rawData, key, iv string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	key_b, err_1 := base64.StdEncoding.DecodeString(key)
	iv_b, _ := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}
	if err_1 != nil {
		return "", err_1
	}
	dnData, err := AesCBCDncrypt(data, key_b, iv_b)
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}

// 解密
func AesCBCDncrypt(encryptData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	if len(encryptData) < blockSize {
		panic("ciphertext too short")
	}
	if len(encryptData)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptData, encryptData)
	// 解填充
	encryptData = PKCS7UnPadding(encryptData)
	return encryptData, nil
}

//去除填充
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
