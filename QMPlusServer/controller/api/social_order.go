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
	"runtime"
	"sort"
	"unsafe"

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

func WxPay(openId, RemoteAddr string, total_fee float64) (error, map[string]interface{}) {

	info := make(map[string]interface{}, 0)
	var resMap = make(map[string]interface{}, 0)
	//ip := RemoteAddr
	orderNo := "wx" + ToStr(time.Now().Unix()) + Krand()
	fmt.Printf("req param %s--%s--%d:\n", openId, RemoteAddr, total_fee)
	fmt.Printf("param %s--%s--%s--%s:\n", ssAppid, ssAppSecret, mchkey, mchid)
	//随机数
	//nonceStr := time.Now().Format("20060102150405") + CreateRand()
	var reqMap = make(map[string]interface{}, 0)
	reqMap["appid"] = ssAppid                                      //微信小程序appid
	reqMap["body"] = "亦庄教育-费用说明" + orderNo                         //商品描述
	reqMap["mch_id"] = mchid                                       //商户号
	reqMap["nonce_str"] = randStr(32, "alphanum")                  //随机数
	reqMap["notify_url"] = "http://127.0.0.1:8888/si/notifyResult" //通知地址
	reqMap["openid"] = openId                                      //商户唯一标识 openid
	reqMap["out_trade_no"] = orderNo                               //订单号
	reqMap["spbill_create_ip"] = getIP(RemoteAddr)                 //用户端ip   //订单生成的机器 IP
	reqMap["total_fee"] = total_fee * 100                          //订单总金额，单位为分
	reqMap["attach"] = "abc"
	reqMap["trade_type"] = "JSAPI" //trade_type=JSAPI时（即公众号支付），此参数必传，此参数为微信用户在商户对应appid下的唯一标识
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
	return nil, resMap
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
		return
	}

	fmt.Printf("PaymentReq param %v:\n", req)
	err, mm := WxPay(req.Openid, c.Request.RemoteAddr, req.TotalFee)
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

func ConfirmPayment(c *gin.Context) {
	/*var req socialInsurance.OrderReqInfo
	_ = c.ShouldBindJSON(&req)

	err, mm := WxPay(req.Openid, c.Request.RemoteAddr, req.OrderNo, req.TotalFee)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("订单支付失败，%v", err), gin.H{})
	} else {
		//1. 更新数据库状态
		servers.ReportFormat(c, true, fmt.Sprintf("订单支付成功，%v", err), gin.H{
			"appId":     mm["appId"],
			"nonceStr":  mm["nonceStr"],
			"package":   mm["package"],
			"timeStamp": mm["timeStamp"],
			"paySign":   mm["paySign"],
		})
	}*/
}

func CancelPayment(c *gin.Context) {
	//1. 更新订单状态
	var err error
	servers.ReportFormat(c, true, fmt.Sprintf("取消订单成功，%v", err), gin.H{})
}

func NotifyResult(c *gin.Context) {
	//1. 更新订单状态
	fmt.Printf("notify.result is :%v\n", c)
	var mr WXPayNotifyReq
	err := c.ShouldBindJSON(&mr)
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

	//var resp WXPayNotifyResp
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
		//resp.Return_code = "FAIL"
		//resp.Return_msg = "failed to verify sign, please retry!"
	}

	//结果返回，微信要求如果成功需要返回return_code "SUCCESS"
	/*bytes, _err := xml.Marshal(resp) //string(bytes)
	strResp := strings.Replace(bytes2str(bytes), "WXPayNotifyResp", "xml", -1)
	if _err != nil {
		fmt.Print("xml编码失败，原因：%v", _err)
		return
	}*/
	//rw.(http.ResponseWriter).WriteHeader(http.StatusOK)
	//fmt.Fprint(rw.(http.ResponseWriter), strResp)
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

	WECHAT_API_KEY := "" //微信商户key
	signCalc := wxpayCalcSign(needVerifyM, WECHAT_API_KEY)
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
