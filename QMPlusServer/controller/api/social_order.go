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
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"sort"
	"unsafe"

	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type WXPayResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
	Nonce_str   string `xml:"nonce_str"`
	Prepay_id   string `xml:"prepay_id"`
}

type WXPayNotifyReq struct {
	Appid          string  `xml:"appid"`
	Bank_type      string  `xml:"bank_type"`
	Cash_fee       float64 `xml:"cash_fee"`
	Fee_type       string  `xml:"fee_type"`
	Is_subscribe   string  `xml:"is_subscribe"`
	Mch_id         string  `xml:"mch_id"`
	Nonce_str      string  `xml:"nonce_str"`
	Openid         string  `xml:"openid"`
	Out_trade_no   string  `xml:"out_trade_no"`
	Result_code    string  `xml:"result_code"`
	Return_code    string  `xml:"return_code"`
	Sign           string  `xml:"sign"`
	Time_end       string  `xml:"time_end"`
	Total_fee      float64 `xml:"total_fee"`
	Trade_type     string  `xml:"trade_type"`
	Transaction_id string  `xml:"transaction_id"`
}

type WXPayNotifyResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
}

type MyMap map[string]interface{}
type xmlMapEntry struct {
	XMLName xml.Name
	Value   interface{} `xml:",chardata"`
}

var ssAppid string = config.GinVueAdminconfig.WeiXin.SocialApp.Appid
var ssAppSecret string = config.GinVueAdminconfig.WeiXin.SocialApp.AppSecret
var mchkey string = config.GinVueAdminconfig.WeiXin.SocialApp.MchKey
var mchid string = config.GinVueAdminconfig.WeiXin.SocialApp.MchId

func AddSocialOrder(c *gin.Context) {
	var req socialInsurance.ReqAddOrder
	_ = c.ShouldBindJSON(&req)
	var so socialInsurance.SocialOrder
	err, reet := so.AddSocialOrder(req)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建订单失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{
			"reet": reet,
		})
	}
}

type ReqUpdate struct {
	Openid  string `json:"openid"`
	OrderId string `json:"order_id"`
	Status  int    `json:"status"`
}

func UpdateOrderStatus(c *gin.Context) {
	var req ReqUpdate
	_ = c.ShouldBindJSON(&req)
	var so socialInsurance.SocialOrder
	err := so.UpdateSocialOrderStatus(req.Openid, req.OrderId, req.Status)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("订单更新失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "更新成功", gin.H{})
	}
}

type ReqOrder struct {
	Openid   string `json:"openid"`
	Status   int    `json:"status"`
	Page     int    `json:"page"`
	PageSize int    `json:"pagesize"`
}

func GetOrders(c *gin.Context) {
	var req ReqOrder
	_ = c.ShouldBindJSON(&req)
	var so socialInsurance.SocialOrder
	err, reet := so.GetOrderList(req.Openid, req.Status, req.Page, req.PageSize)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取订单失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取成功", gin.H{
			"reet": reet,
		})
	}
}

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

func WxPay(openId, RemoteAddr, orderNo string, total_fee float64) (error, map[string]interface{}) {

	info := make(map[string]interface{}, 0)
	var resMap = make(map[string]interface{}, 0)
	//ip := RemoteAddr
	//orderNo := "wx" + ToStr(time.Now().Unix()) + Krand()
	fmt.Printf("req param %s--%s--%d:\n", openId, RemoteAddr, total_fee)
	fmt.Printf("param %s--%s--%s--%s:\n", ssAppid, ssAppSecret, mchkey, mchid)
	//随机数
	//nonceStr := time.Now().Format("20060102150405") + CreateRand()
	var reqMap = make(map[string]interface{}, 0)
	reqMap["appid"] = ssAppid                                          //微信小程序appid
	reqMap["body"] = "亦庄教育-费用说明" + orderNo                             //商品描述
	reqMap["mch_id"] = mchid                                           //商户号
	reqMap["nonce_str"] = randStr(32, "alphanum")                      //随机数
	reqMap["notify_url"] = "http://47.93.185.194:8888/si/notifyResult" //通知地址
	reqMap["openid"] = openId                                          //商户唯一标识 openid
	reqMap["out_trade_no"] = orderNo                                   //订单号
	reqMap["spbill_create_ip"] = getIP(RemoteAddr)                     //用户端ip   //订单生成的机器 IP
	reqMap["total_fee"] = Float2String(total_fee * 10)                 //订单总金额，单位为分
	reqMap["trade_type"] = "JSAPI"                                     //trade_type=JSAPI时（即公众号支付），此参数必传，此参数为微信用户在商户对应appid下的唯一标识
	reqMap["sign"] = WxPayCalcSign(reqMap, mchkey)

	//reqStr := Map2Xml(reqMap)
	reqStr, _ := xml.Marshal(MyMap(reqMap))
	fmt.Println("请求xml:", reqStr)

	// 调用支付统一下单API
	bytes_req := []byte(reqStr)
	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/pay/unifiedorder", bytes.NewReader(bytes_req))
	if err != nil {
		return err, nil
		// handle error
	}
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	//req.Header.Set("Content-Type", "text/xml;charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf(" client.Do err:%v\n", err)
		return err, nil
	}
	defer resp.Body.Close()

	body2, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("解析响应内容失败", err)
		return err, nil
	}
	fmt.Println("响应数据", string(body2))

	var resp1 WXPayResp
	err = xml.Unmarshal(body2, &resp1)
	if err != nil {
		return err, nil
	}
	if strings.ToUpper(resp1.Return_code) == "FAIL" {
		return fmt.Errorf("wx response err ,info:%s", string(body2)), nil
	}

	// 返回预付单信息
	if strings.ToUpper(resp1.Return_code) == "SUCCESS" {
		fmt.Println("预支付申请成功")
		// 再次签名

		resMap["appId"] = ssAppid
		//resMap["nonceStr"] = resp1.Nonce_str               //商品描述
		resMap["nonceStr"] = RandomString()
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
	return nil, resMap
}

// 微信金额浮点转字符串
func Float2String(moneyFee float64) string {
	aDecimal := decimal.NewFromFloat(moneyFee)
	bDecimal := decimal.NewFromFloat(100)
	return aDecimal.Mul(bDecimal).Truncate(0).String()
}

func WxRefund(openId, RemoteAddr, outTradeNo string, total_fee, refund_fee float64) error {

	//info := make(map[string]interface{}, 0)
	//var resMap = make(map[string]interface{}, 0)
	//ip := RemoteAddr
	outRefundNo := "wx" + ToStr(time.Now().Unix()) + Krand()
	fmt.Printf("req param %s--%s--%d:\n", openId, RemoteAddr, total_fee)
	fmt.Printf("param %s--%s--%s--%s:\n", ssAppid, ssAppSecret, mchkey, mchid)
	//随机数
	//nonceStr := time.Now().Format("20060102150405") + CreateRand()
	var reqMap = make(map[string]interface{}, 0)
	reqMap["appid"] = ssAppid                                                //微信小程序appid
	reqMap["body"] = "亦庄教育-费用说明" + outTradeNo                                //商品描述
	reqMap["mch_id"] = mchid                                                 //商户号
	reqMap["nonce_str"] = randStr(32, "alphanum")                            //随机数
	reqMap["notify_url"] = "http://47.93.185.194:8888/si/notifyRefundResult" //通知地址
	reqMap["openid"] = openId                                                //商户唯一标识 openid
	reqMap["out_trade_no"] = outTradeNo                                      //订单号
	reqMap["out_refund_no"] = outRefundNo                                    //退款单号
	reqMap["spbill_create_ip"] = getIP(RemoteAddr)                           //用户端ip   //订单生成的机器 IP
	reqMap["total_fee"] = total_fee * 100                                    //订单总金额，单位为分
	reqMap["refund_fee"] = refund_fee * 100
	reqMap["op_user_id"] = mchid
	reqMap["trade_type"] = "JSAPI" //trade_type=JSAPI时（即公众号支付），此参数必传，此参数为微信用户在商户对应appid下的唯一标识
	reqMap["sign"] = WxPayCalcSign(reqMap, mchkey)

	//reqStr := Map2Xml(reqMap)
	reqStr, _ := xml.Marshal(MyMap(reqMap))
	fmt.Println("请求xml:", reqStr)

	// 调用支付统一下单API
	bytes_req := []byte(reqStr)
	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/secapi/pay/refund", bytes.NewReader(bytes_req))
	if err != nil {
		return err
		// handle error
	}
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	//req.Header.Set("Content-Type", "text/xml;charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf(" client.Do err:%v\n", err)
		return err
	}
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
		return err
	}
	if strings.ToUpper(resp1.Return_code) == "FAIL" {
		return fmt.Errorf("wx response err ,info:%s", string(body2))
	}

	// 返回预付单信息
	if strings.ToUpper(resp1.Return_code) == "SUCCESS" {
		return nil
	} else {
		return fmt.Errorf("微信请求支付失败")
		//index.Console(info)
	}
}

//微信支付计算签名的函数
func WxPayCalcSign_old(mReq map[string]interface{}, key string) (sign string) {
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
func WxPayCalcSign(m map[string]interface{}, key string) string {
	var signData []string
	for k, v := range m {
		//签名之前的拼接数据，需要过滤掉sign和key关键字
		if v != "" && k != "sign" && k != "key" {
			signData = append(signData, fmt.Sprintf("%s=%s", k, v))
		}
	}
	//按ascii字母排序
	sort.Strings(signData)
	//将字符串数组按照&拼接起来
	signStr := strings.Join(signData, "&")
	//最后拼接上key，得到签名之前的字符串
	signStr = signStr + "&key=" + key
	log.Println("签名之前的字符串------------", signStr)
	md5Handle := md5.New()
	_, err := md5Handle.Write([]byte(signStr))
	if err != nil {
		return ""
	}
	signByte := md5Handle.Sum(nil)
	if err != nil {
		return ""
	}

	tosign := strings.ToUpper(fmt.Sprintf("%x", signByte))
	log.Println("签名的结果为-------", tosign)
	return tosign
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

func (m MyMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}

	return e.EncodeToken(start.End())
}

func ToStr(idata interface{}) string {
	return fmt.Sprintf("%d", idata)
}

func Krand() string {
	return fmt.Sprintf("%.4d", rand.Int31()%10000)
}
func RandomString() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
func CreateRand() string {
	return fmt.Sprintf("%18v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

func randStr(strSize int, randType string) string {

	var dictionary string

	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "alpha" {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "number" {
		dictionary = "0123456789"
	}

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func PaymentReq(c *gin.Context) {
	var req socialInsurance.OrderReqInfo
	//_ = c.ShouldBindJSON(&req)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Print("解析HTTP Body格式到xml失败，原因!", err)
		servers.ReportFormat(c, false, fmt.Sprintf("解析HTTP Body格式到xml失败，原因，%v", err), gin.H{})
		return
	}

	fmt.Printf("PaymentReq param %v:\n", req)
	if len(req.Openid) <= 0 || len(req.OrderNo) == 0 || req.TotalFee <= 0 {
		fmt.Print("请求参数错误 ", err)
		servers.ReportFormat(c, false, fmt.Sprintf("请求参数错误"), gin.H{})
		return
	}

	err, mm := WxPay(req.Openid, c.Request.RemoteAddr, req.OrderNo, req.TotalFee)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("请求下单失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, fmt.Sprintf("请求下单成功，%v", err), gin.H{
			"appId":     mm["appId"],
			"nonceStr":  mm["nonceStr"],
			"package":   mm["package"],
			"timeStamp": mm["timeStamp"],
			"encryType": "MD5",
			"paySign":   mm["paySign"],
		})
	}
}

func CancelPayment(c *gin.Context) {
	//1. 更新订单状态
	var req socialInsurance.OrderReqInfo
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Print("解析HTTP Body格式到xml失败，原因!", err)
		servers.ReportFormat(c, false, fmt.Sprintf("解析HTTP Body格式到xml失败，原因，%v", err), gin.H{})
		return
	}

	var so socialInsurance.SocialOrder
	so.OrderId = req.OrderNo
	err, _ = so.FindByOrderId()
	if err != nil {
		fmt.Print("系统订单查询价格错误", err)
		servers.ReportFormat(c, false, fmt.Sprintf("取消订单 订单号%s在系统中没有查找到 %v", req.OrderNo, err), gin.H{})
		return
	}

	status := 4 //订单取消，已放弃
	err = so.UpdateSocialOrderStatus(req.Openid, req.OrderNo, status)
	if err != nil {
		fmt.Print("系统UpdateSocialOrderStatus 更新订单失败", err)
		servers.ReportFormat(c, false, fmt.Sprintf("取消订单 订单号%s在系统中没有查找到 %v", req.OrderNo, err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, fmt.Sprintf("取消订单成功，%v", err), gin.H{})
	//status = 4,5 已失效,已放弃
}

//订单退款
func RefundReq(c *gin.Context) {
	var req socialInsurance.RefundReqInfo
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Print("解析HTTP Body格式到xml失败，原因!", err)
		servers.ReportFormat(c, false, fmt.Sprintf("解析HTTP Body格式到xml失败，原因，%v", err), gin.H{})
		return
	}

	fmt.Printf("PaymentReq param %v:\n", req)
	if len(req.Openid) <= 0 || len(req.OrderNo) == 0 || req.TotalFee <= 0 {
		fmt.Print("请求参数错误 ", err)
		servers.ReportFormat(c, false, fmt.Sprintf("请求参数错误"), gin.H{})
		return
	}

	err = WxRefund(req.Openid, c.Request.RemoteAddr, req.OrderNo, req.TotalFee, req.RefundFee)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("申请退款失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, fmt.Sprintf("申请退款成功，%v", err), gin.H{})
	}
}

//查询退款
//提交退款申请后，通过调用该接口查询退款状态。退款有一定延时，用零钱支付的退款20分钟内到账，银行卡支付的退款3个工作日后重新查询退款状态。
//注意：如果单个支付订单部分退款次数超过20次请使用退款单号查询
func RefundQuery(c *gin.Context) {
	/*err := WxRefund(req.Openid, c.Request.RemoteAddr, req.OrderNo, req.TotalFee, req.RefundFee)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("申请退款失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, fmt.Sprintf("申请退款成功，%v", err), gin.H{})
	}*/
}

func NotifyResult_old(c *gin.Context) {
	//1. 更新订单状态
	fmt.Printf("notify.result is :%v\n", c)
	var mr WXPayNotifyReq
	err := c.ShouldBindXML(&mr)
	if err != nil {
		fmt.Print("解析HTTP Body格式到xml失败，原因!", err)
		return
	}

	var reqMap map[string]interface{}
	reqMap = make(map[string]interface{}, 0)

	reqMap["appid"] = mr.Appid
	reqMap["bank_type"] = mr.Bank_type
	reqMap["cash_fee"] = mr.Cash_fee
	reqMap["fee_type"] = mr.Fee_type
	reqMap["is_subscribe"] = mr.Is_subscribe
	reqMap["mch_id"] = mr.Mch_id
	reqMap["nonce_str"] = mr.Nonce_str
	reqMap["openid"] = mr.Openid
	reqMap["out_trade_no"] = mr.Out_trade_no
	reqMap["result_code"] = mr.Result_code
	reqMap["return_code"] = mr.Return_code
	reqMap["time_end"] = mr.Time_end
	reqMap["total_fee"] = mr.Total_fee
	reqMap["trade_type"] = mr.Trade_type
	reqMap["transaction_id"] = mr.Transaction_id
	reqMap["signType"] = "MD5"

	var resp WXPayNotifyResp
	//进行签名校验
	if wxpayVerifySign(reqMap, mr.Sign) {
		//transactionId := reqMap["transaction_id"]
		//orderCode := reqMap["out_trade_no"]
		total_fee := reqMap["total_fee"].(float64) //分->元 除以100

		var so socialInsurance.SocialOrder
		so.OrderId = mr.Out_trade_no
		err, resultOrder := so.FindByOrderId()
		if err != nil {
			fmt.Print("系统订单查询价格错误", err)
			return
		}
		if resultOrder.TotalFee != total_fee {
			fmt.Print("系统订单价格与支付金额不符合，错误", err)
			return
		}

		status := 3 //订单已完成
		err = so.UpdateSocialOrderStatus(mr.Openid, mr.Out_trade_no, status)
		if err != nil {
			fmt.Print("系统UpdateSocialOrderStatus 更新订单失败", err)
			return
		}

		resp.Return_code = "SUCCESS"
		resp.Return_msg = "OK"
	} else {
		resp.Return_code = "FAIL"
		resp.Return_msg = "failed to verify sign, please retry!"
	}

	//结果返回，微信要求如果成功需要返回return_code "SUCCESS"
	fmt.Print("result weixin1:%s", resp)
	bytes, _err := xml.Marshal(resp) //string(bytes)
	strResp := strings.Replace(bytes2str(bytes), "WXPayNotifyResp", "xml", -1)
	if _err != nil {
		fmt.Print("xml编码失败，原因：%v", _err)
		return
	}
	fmt.Print("result weixin2:%s", strResp)
	servers.ReportFormatXML(c, strResp)

	//rw.(http.ResponseWriter).WriteHeader(http.StatusOK)
	//fmt.Fprint(rw.(http.ResponseWriter), strResp)
}

type MiniPayCommonResult struct {
	ReturnCode string `xml:"return_code" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg" json:"return_msg,omitempty"`
}
type MiniPayResultData struct {
	OpenID             string `xml:"openid,omitempty" json:"openid,omitempty"`
	IsSubscribe        string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"`
	TradeType          string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	TradeState         string `xml:"trade_state,omitempty" json:"trade_state,omitempty"`
	BankType           string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	TotalFee           string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee string `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	FeeType            string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	CashFee            string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	CashFeeType        string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty"`

	TransactionID  string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNO     string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach         string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd        string `xml:"time_end,omitempty" json:"time_end,omitempty"`
	TradeStateDesc string `xml:"trade_state_desc" json:"trade_state_desc,omitempty"`
}
type MiniPayReturnSuccessData struct {
	AppID string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchID string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	//DeviceInfo 统一下单默认就有，查询结果在return_code 、result_code、trade_state都为SUCCESS时有返回，这里统一放在这里
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
}
type MiniPayAsyncResult struct {
	//统一下单与查询结果通用部分
	MiniPayCommonResult
	//统一下单与查询结果通用部分
	MiniPayReturnSuccessData
	// 查询结果或者下单返回公用部分
	MiniPayResultData
}

func NotifyResult(c *gin.Context) {
	var returnCode = "FAIL"
	var returnMsg = ""
	var miniPayCommonResult MiniPayCommonResult
	//log.Println(w)
	defer func() {
		log.Println("log.Println(miniPayCommonResult)---before-----", miniPayCommonResult)
		miniPayCommonResult.ReturnCode = returnCode
		miniPayCommonResult.ReturnMsg = returnMsg
		sendResultRsp(c, "SUCCESS", "SUCCESS")
		log.Println("log.Println(miniPayCommonResult)---after-----", miniPayCommonResult)

	}()
	//var reXML MiniPayAsyncResult
	/*log.Println(body)
	log.Println(string(body))
	if string(body) == "" {
		returnCode = "FAIL"
		returnMsg = "Bodyerror"
		return &reXML, &miniPayCommonResult, errors.New(returnCode + ":" + returnMsg)
	}
	err = xml.Unmarshal(body, &reXML)
	if err != nil {
		returnCode = "FAIL"
		returnMsg = "参数错误"
		return &reXML, &miniPayCommonResult, errors.New(returnCode + ":" + returnMsg)
	}*/
	var mr WXPayNotifyReq
	//var resp WXPayNotifyResp

	err := c.ShouldBindXML(&mr)
	if err != nil {
		fmt.Print("解析HTTP Body格式到xml失败，原因!", err)
		return
	}
	var reqMap map[string]interface{}
	reqMap = make(map[string]interface{}, 0)

	reqMap["appid"] = mr.Appid
	reqMap["bank_type"] = mr.Bank_type
	reqMap["cash_fee"] = mr.Cash_fee
	reqMap["fee_type"] = mr.Fee_type
	reqMap["is_subscribe"] = mr.Is_subscribe
	reqMap["mch_id"] = mr.Mch_id
	reqMap["nonce_str"] = mr.Nonce_str
	reqMap["openid"] = mr.Openid
	reqMap["out_trade_no"] = mr.Out_trade_no
	reqMap["result_code"] = mr.Result_code
	reqMap["return_code"] = mr.Return_code
	reqMap["time_end"] = mr.Time_end
	reqMap["total_fee"] = mr.Total_fee
	reqMap["trade_type"] = mr.Trade_type
	reqMap["transaction_id"] = mr.Transaction_id

	if mr.Result_code != "SUCCESS" {
		sendResultRsp(c, "FAIL", "failed to verify sign, please retry!")
		return
	}
	//m := XmlToMap(body)
	var signData []string
	for k, v := range reqMap {
		if k == "sign" {
			continue
		}
		signData = append(signData, fmt.Sprintf("%v=%v", k, v))
	}

	log.Println(signData)

	//log.Println("minipay()----", &Minipay().Key)
	//key := Minipay().Key
	//log.Println("key------", key)
	//mySign, err := MinipaySign(key, m)
	mySign := wxpayCalcSign(reqMap, mchkey)
	if err != nil {
		sendResultRsp(c, "FAIL", "MinipaySign not eq!")
		return
	}
	fmt.Printf("计算出来的sign:%s \n", mySign)
	fmt.Printf("微信异步通知sign:%s \n", reqMap["sign"])

	if mySign != reqMap["sign"] {
		fmt.Printf("签名交易错误")
		sendResultRsp(c, "FAIL", "singal err!")
	}

	sendResultRsp(c, "SUCCESS", "SUCCESS")
}

func sendResultRsp(c *gin.Context, code, msg string) {
	resp := WXPayNotifyResp{
		Return_code: code,
		Return_msg:  msg,
	}

	bytes, _err := xml.Marshal(resp) //string(bytes)
	strResp := strings.Replace(bytes2str(bytes), "WXPayNotifyResp", "xml", -1)
	if _err != nil {
		fmt.Print("xml编码失败，原因：%v", _err)
		return
	}
	fmt.Print("result weixin2:%s", strResp)
	servers.ReportFormatXML(c, strResp)
}

func NotifyRefundResult(c *gin.Context) {
	//1. 更新订单状态
	fmt.Printf("notify. refund result is :%v\n", c)
	var mr WXPayNotifyReq
	err := c.ShouldBindXML(&mr)
	if err != nil {
		fmt.Print("解析HTTP Body格式到xml失败，原因!", err)
		return
	}
	return
}

func WeixinNoticeHandler(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Print("读取http body失败，原因!", err)
		http.Error(rw.(http.ResponseWriter), http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	fmt.Print("微信支付异步通知，HTTP Body:", string(body))

	var mr WXPayNotifyReq
	err = xml.Unmarshal(body, &mr)
	if err != nil {
		fmt.Print("解析HTTP Body格式到xml失败，原因!", err)
		http.Error(rw.(http.ResponseWriter), http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var reqMap map[string]interface{}
	reqMap = make(map[string]interface{}, 0)

	reqMap["appid"] = mr.Appid
	reqMap["bank_type"] = mr.Bank_type
	reqMap["cash_fee"] = mr.Cash_fee
	reqMap["fee_type"] = mr.Fee_type
	reqMap["is_subscribe"] = mr.Is_subscribe
	reqMap["mch_id"] = mr.Mch_id
	reqMap["nonce_str"] = mr.Nonce_str
	reqMap["openid"] = mr.Openid
	reqMap["out_trade_no"] = mr.Out_trade_no
	reqMap["result_code"] = mr.Result_code
	reqMap["return_code"] = mr.Return_code
	reqMap["time_end"] = mr.Time_end
	reqMap["total_fee"] = mr.Total_fee
	reqMap["trade_type"] = mr.Trade_type
	reqMap["transaction_id"] = mr.Transaction_id

	var resp WXPayNotifyResp
	//进行签名校验
	if wxpayVerifySign(reqMap, mr.Sign) {
		//transactionId := reqMap["transaction_id"]
		/*orderCode := reqMap["out_trade_no"]
		total_fee := reqMap["total_fee"].(float64) //分->元 除以100
		rows, err := mysqlDB.Query("SELECT * FROM canyin_order WHERE dno = ?", orderCode)
		if err != nil {
			fmt.Print("微信查询价格错误", err)
			return
		}
		defer rows.Close()
		orders := RowResult(rows)
		if len(orders) > 0 {
			orderInfo := orders[0].(map[string]interface{})
			//orderId := ToStr(orderInfo["id"])
			allcost, _ := strconv.ParseFloat(ToStr(orderInfo["allcost"]), 64)
			fmt.Print("价格比对", "---", allcost, "---", total_fee)
			//商户系统对于支付结果通知的内容一定要做签名验证,并校验返回的订单金额是否与商户侧的订单金额一致，防止数据泄漏导致出现“假通知”，造成资金损失
			if allcost == total_fee {
				fmt.Print("订单验证成功")
				//以下是业务处理
			}
			resp.Return_code = "SUCCESS"
			resp.Return_msg = "OK"
		} else {
			resp.Return_code = "FAIL"
			resp.Return_msg = "无此订单"
		}*/
	} else {
		resp.Return_code = "FAIL"
		resp.Return_msg = "failed to verify sign, please retry!"
	}

	//结果返回，微信要求如果成功需要返回return_code "SUCCESS"
	bytes, _err := xml.Marshal(resp) //string(bytes)
	strResp := strings.Replace(bytes2str(bytes), "WXPayNotifyResp", "xml", -1)
	if _err != nil {
		fmt.Print("xml编码失败，原因：%v", _err)
		http.Error(rw.(http.ResponseWriter), http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	rw.(http.ResponseWriter).WriteHeader(http.StatusOK)
	fmt.Fprint(rw.(http.ResponseWriter), strResp)
}

//微信支付签名验证函数
func wxpayVerifySign(needVerifyM map[string]interface{}, sign string) bool {
	pc, _, line, _ := runtime.Caller(0)
	fc := runtime.FuncForPC(pc)

	//WECHAT_API_KEY := mchkey //微信商户key
	signCalc := wxpayCalcSign(needVerifyM, mchkey)
	fmt.Print(fc.Name(), line, "计算出来的sign: ", signCalc)
	fmt.Print(fc.Name(), line, "微信异步通知sign: ", sign)
	if sign == signCalc {
		fmt.Print(fc.Name(), line, "签名校验通过!")
		return true
	}

	fmt.Print(fc.Name(), line, "签名校验失败!")
	return false
}

//微信支付计算签名的函数
func wxpayCalcSign(mReq map[string]interface{}, key string) (sign string) {
	//方法名 行数
	pc, _, line, _ := runtime.Caller(0)
	fc := runtime.FuncForPC(pc)

	fmt.Print(fc.Name(), line, "微信支付签名计算, API KEY:", key)
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}

	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		fmt.Print("k=%v, v=%v\n", k, mReq[k])
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write(str2bytes(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}
func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func getIP(ip string) string {
	r := strings.Split(ip, ":")
	return r[0]
}

func XmlToMap(xmlData []byte) map[string]interface{} {
	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	m := make(map[string]interface{})
	var token xml.Token
	var err error
	var k string
	for token, err = decoder.Token(); err == nil; token, err = decoder.Token() {
		if v, ok := token.(xml.StartElement); ok {
			k = v.Name.Local
			continue
		}
		if v, ok := token.(xml.CharData); ok {
			data := string(v.Copy())
			if strings.TrimSpace(data) == "" {
				continue
			}
			m[k] = data
		}
	}

	if err != nil && err != io.EOF {
		panic(err)
	}
	return m
}
