package api

import (
	"encoding/json"
	"fmt"
	"gin-vue-admin/config"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/socialInsurance"
	"gin-vue-admin/model/userJobs"

	"github.com/gin-gonic/gin"
)

var sappid string = config.GinVueAdminconfig.WeiXinPayInfo.SocialSecurityApp.Appid
var sappSecret string = config.GinVueAdminconfig.WeiXinPayInfo.SocialSecurityApp.AppSecret

// @Tags SocialUser
// @Summary 创建SocialUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialUser true "创建SocialUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /su/createSocialUser [post]
func CreateSocialUser(c *gin.Context) {
	var su socialInsurance.SocialUser
	_ = c.ShouldBindJSON(&su)
	err := su.CreateSocialUser()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SocialUser
// @Summary 删除SocialUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialUser true "删除SocialUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /su/deleteSocialUser [post]
func DeleteSocialUser(c *gin.Context) {
	var su socialInsurance.SocialUser
	_ = c.ShouldBindJSON(&su)
	err := su.DeleteSocialUser()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SocialUser
// @Summary 更新SocialUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialUser true "更新SocialUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /su/updateSocialUser [post]
func UpdateSocialUser(c *gin.Context) {
	var su socialInsurance.SocialUser
	_ = c.ShouldBindJSON(&su)
	err, reet := su.UpdateSocialUser()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reet": reet,
		})
	}
}

// @Tags SocialUser
// @Summary 用id查询SocialUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialUser true "用id查询SocialUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /su/findSocialUser [post]
func FindSocialUser(c *gin.Context) {
	var su socialInsurance.SocialUser
	_ = c.ShouldBindJSON(&su)
	err, reet := su.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reet": reet,
		})
	}
}

func UpdateUserByOpenid(c *gin.Context) {
	var su socialInsurance.SocialUser
	_ = c.ShouldBindJSON(&su)
	err := su.UpdateSocialUserByOpenid()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{})
	}
}

func FindSocialUserByOpenid(c *gin.Context) {
	var su socialInsurance.SocialUser
	_ = c.ShouldBindJSON(&su)
	err, reet := su.FindByOpenid()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reet": reet,
		})
	}
}

// @Tags SocialUser
// @Summary 分页获取SocialUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取SocialUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /su/getSocialUserList [post]
func GetSocialUserList(c *gin.Context) {
	type searchParams struct {
		socialInsurance.SocialUser
		modelInterface.PageInfo
	}
	var sp searchParams
	_ = c.ShouldBindJSON(&sp)
	err, list, total := sp.SocialUser.GetInfoList(sp.PageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"socialUserList": list,
			"total":          total,
			"page":           sp.PageInfo.Page,
			"pageSize":       sp.PageInfo.PageSize,
		})
	}
}

// func SocialUserLogin(c *gin.Context) {
// 	var loginInfo userJobs.UserLoginInfo
// 	_ = c.ShouldBindJSON(&loginInfo)
// 	if len(loginInfo.Code) == 0 {
// 		servers.ReportFormat(c, false, fmt.Sprintf("创建失败"), gin.H{})
// 		return
// 	}

// 	//==============develop===================
// 	u := socialInsurance.SocialUser{
// 		Openid: loginInfo.Code,
// 		Status: 1,
// 	}
// 	err, _ := u.FindById()
// 	if err != nil {
// 		_ = c.ShouldBindJSON(&u)
// 		err = u.CreateSocialUser()
// 		if err != nil {
// 			servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
// 			return
// 		}
// 	}
// 	//var isMobile bool = true

// 	//tokenNext_wx(c, u.Openid, "", isMobile)
// 	//========================================

// }

func SocialUserLogin(c *gin.Context) {
	var loginInfo userJobs.UserLoginInfo
	var userName string
	_ = c.ShouldBindJSON(&loginInfo)
	if len(loginInfo.Code) == 0 {
		servers.ReportFormat(c, false, fmt.Sprintf("登录失败"), gin.H{})
		return
	}

	openid, session_key, err := sendWxAuthAPI(sappid, sappSecret, loginInfo.Code)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("登录失败：%v", err), gin.H{})
		return
	}

	u := socialInsurance.SocialUser{
		Openid: openid,
		Status: 1,
	}
	err, reuser := u.FindByOpenid()
	if err != nil {
		err = u.CreateSocialUser()
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
			return
		}
	} else {
		userName = reuser.Name
	}
	var isMobile bool = false
	if len(reuser.Mobile) > 0 {
		isMobile = true
	}
	tokenNext_wx(c, openid, session_key, isMobile, userName)
}

func DecodeSocialMobile(c *gin.Context) {

	var enMobile userJobs.DecodeMobile

	_ = c.ShouldBindJSON(&enMobile)

	if len(enMobile.SessionKey) == 0 || len(enMobile.Openid) == 0 {
		servers.ReportFormat(c, false, fmt.Sprintf("SessionKey empty err"), gin.H{})
		return
	}

	src, err := Dncrypt(enMobile.EncrypData, enMobile.SessionKey, enMobile.IvData)

	var s = map[string]interface{}{}

	json.Unmarshal([]byte(src), &s)

	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("Dncrypt ：%v", err), gin.H{})
		return
	}

	err = updateSocialUserMobile(enMobile.Openid, s["phoneNumber"].(string))
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("updateUserMobile err :%v", err), gin.H{})
		return
	}

	servers.ReportFormat(c, true, "更新成功", gin.H{
		"data": "ok",
	})
}

func updateSocialUserMobile(openId, mobileNum string) error {

	if len(openId) == 0 || len(mobileNum) != 11 {
		return fmt.Errorf("param empty ")
	}

	u := socialInsurance.SocialUser{
		Openid: openId,
		Mobile: mobileNum,
	}

	err := u.UpdateUsersMobile()
	if err != nil {
		return err
	}
	return nil
}
