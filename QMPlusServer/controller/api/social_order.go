package api

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"gin-vue-admin/config"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/socialInsurance"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type WXPayResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
	Nonce_str   string `xml:"nonce_str"`
	Prepay_id   string `xml:"prepay_id"`
}

var ssAppid string = config.GinVueAdminconfig.WeiXinPayInfo.SocialSecurityApp.Appid
var ssAppSecret string = config.GinVueAdminconfig.WeiXinPayInfo.SocialSecurityApp.AppSecret
var mchkey string = config.GinVueAdminconfig.WeiXinPayInfo.MchKey
var mchid string = config.GinVueAdminconfig.WeiXinPayInfo.MchId

// @Tags SocialOrder
// @Summary 创建SocialOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialOrder true "创建SocialOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /so/createSocialOrder [post]
func CreateSocialOrder(c *gin.Context) {
	var so socialInsurance.SocialOrder
	_ = c.ShouldBindJSON(&so)
	err := so.CreateSocialOrder()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SocialOrder
// @Summary 删除SocialOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialOrder true "删除SocialOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /so/deleteSocialOrder [post]
func DeleteSocialOrder(c *gin.Context) {
	var so socialInsurance.SocialOrder
	_ = c.ShouldBindJSON(&so)
	err := so.DeleteSocialOrder()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}

// @Tags SocialOrder
// @Summary 更新SocialOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialOrder true "更新SocialOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /so/updateSocialOrder [post]
func UpdateSocialOrder(c *gin.Context) {
	var so socialInsurance.SocialOrder
	_ = c.ShouldBindJSON(&so)
	err, reet := so.UpdateSocialOrder()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{
			"reet": reet,
		})
	}
}

// @Tags SocialOrder
// @Summary 用id查询SocialOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialOrder true "用id查询SocialOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /so/findSocialOrder [post]
func FindSocialOrder(c *gin.Context) {
	var so socialInsurance.SocialOrder
	_ = c.ShouldBindJSON(&so)
	err, reet := so.FindById()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("查询失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "查询成功", gin.H{
			"reet": reet,
		})
	}
}

// @Tags SocialOrder
// @Summary 分页获取SocialOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取SocialOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /so/getSocialOrderList [post]
func GetSocialOrderList(c *gin.Context) {
	type searchParams struct {
		socialInsurance.SocialOrder
		modelInterface.PageInfo
	}
	var sp searchParams
	_ = c.ShouldBindJSON(&sp)
	err, list, total := sp.SocialOrder.GetInfoList(sp.PageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"socialOrderList": list,
			"total":           total,
			"page":            sp.PageInfo.Page,
			"pageSize":        sp.PageInfo.PageSize,
		})
	}
}

func WxPay(openId, RemoteAddr string, orderNo string, total_fee int) error {
	info := make(map[string]interface{}, 0)

	//fmt.Println("访问ip", index.Request.RemoteAddr)
	//ip := utils.Substr(index.Request.RemoteAddr, 0, strings.Index(index.Request.RemoteAddr, ":"))
	ip := RemoteAddr
	//total_fee, _ := strconv.ParseFloat(index.GetString("total_fee"), 64) //单位 分
	//openId := index.GetString("openId") //"oKYr_0GkE-Izt9N9Wn43sapI9Pqw"
	body := "费用说明"
	//订单号
	//orderNo := index.GetString("orderNo") //"wx"+utils.ToStr(time.Now().Unix()) + string(utils.Krand(4, 0))
	//随机数
	nonceStr := time.Now().Format("20060102150405") + CreateRand()
	var reqMap = make(map[string]interface{}, 0)
	reqMap["appid"] = ssAppid                                      //微信小程序appid
	reqMap["body"] = body                                          //商品描述
	reqMap["mch_id"] = mchid                                       //商户号
	reqMap["nonce_str"] = nonceStr                                 //随机数
	reqMap["notify_url"] = "http://127.0.0.1:8888/si/notifyResult" //通知地址
	reqMap["openid"] = openId                                      //商户唯一标识 openid
	reqMap["out_trade_no"] = orderNo                               //订单号
	reqMap["spbill_create_ip"] = ip                                //用户端ip   //订单生成的机器 IP
	reqMap["total_fee"] = total_fee * 100                          //订单总金额，单位为分
	reqMap["trade_type"] = "JSAPI"                                 //trade_type=JSAPI时（即公众号支付），此参数必传，此参数为微信用户在商户对应appid下的唯一标识
	reqMap["sign"] = WxPayCalcSign(reqMap, mchkey)

	reqStr := Map2Xml(reqMap)
	fmt.Println("请求xml", reqStr)

	client := &http.Client{}

	// 调用支付统一下单API
	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/pay/unifiedorder", strings.NewReader(reqStr))
	if err != nil {
		return err
		// handle error
	}
	req.Header.Set("Content-Type", "text/xml;charset=utf-8")

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body2, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("解析响应内容失败", err)
		return err
	}
	fmt.Println("响应数据", string(body2))

	var resp1 WXPayResp
	err = xml.Unmarshal(body2, &resp1)
	if err != nil {
		panic(err)
		return err
	}
	if strings.ToUpper(resp1.Return_code) == "FAIL" {
		return fmt.Errorf("wx response err ,info:%s", string(body2))
	}

	// 返回预付单信息
	if strings.ToUpper(resp1.Return_code) == "SUCCESS" {
		fmt.Println("预支付申请成功")
		// 再次签名
		var resMap = make(map[string]interface{}, 0)
		resMap["appId"] = ssAppid
		resMap["nonceStr"] = resp1.Nonce_str               //商品描述
		resMap["package"] = "prepay_id=" + resp1.Prepay_id //商户号
		resMap["signType"] = "MD5"                         //签名类型
		resMap["timeStamp"] = ToStr(time.Now().Unix())     //当前时间戳

		resMap["paySign"] = WxPayCalcSign(resMap, mchkey)
		// 返回5个支付参数及sign 用户进行确认支付

		fmt.Println("支付参数", resMap)
		//index.Console(resMap)
	} else {
		info["msg"] = "微信请求支付失败"
		//index.Console(info)
	}
	return nil
}

//微信支付计算签名的函数
func WxPayCalcSign(mReq map[string]interface{}, key string) (sign string) {
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		fmt.Printf("k=%v, v=%v\n", k, mReq[k])
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	fmt.Println("加密前-----", signStrings)
	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings)) //
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))

	fmt.Println("加密后-----", upperSign)
	return upperSign
}

//微信支付计算签名的函数
func Map2Xml(mReq map[string]interface{}) (xml string) {
	sb := bytes.Buffer{}
	sb.WriteString("<xml>")
	for k, v := range mReq {
		sb.WriteString("<" + k + ">" + ToStr(v) + "</" + k + ">")
	}
	sb.WriteString("</xml>")
	return sb.String()
}

func ToStr(idata interface{}) string {
	return fmt.Sprintf("%d", idata)
}

func CreateRand() string {
	return fmt.Sprintf("%18v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

func PaymentReq(c *gin.Context) {
	//var req OrderReqInfo
	//_ = c.ShouldBindJSON(&req)
	//1. 获取orderNo test
	var err error
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("订单支付失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "请求支付成功", gin.H{
			"open_id":   "aaaaa",
			"order_no":  "112233",
			"prepay_id": "aaaabbb",
		})
	}
}

func ConfirmPayment(c *gin.Context) {
	var req socialInsurance.OrderReqInfo
	_ = c.ShouldBindJSON(&req)

	err := WxPay(req.Openid, req.RemoteAddr, req.OrderNo, req.TotalFee)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("订单支付失败，%v", err), gin.H{})
	} else {
		//1. 更新数据库状态
		servers.ReportFormat(c, true, fmt.Sprintf("订单支付成功，%v", err), gin.H{})
	}
}

func CancelPayment(c *gin.Context) {
	//1. 更新订单状态
	var err error
	servers.ReportFormat(c, true, fmt.Sprintf("取消订单成功，%v", err), gin.H{})
}

func NotifyResult(c *gin.Context) {
	//1. 更新订单状态
	fmt.Printf("notify.result is :%v\n", c)
}
