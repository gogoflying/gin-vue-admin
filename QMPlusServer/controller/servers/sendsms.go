package servers

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gin-vue-admin/config"
	"io/ioutil"
	"net/http"
	"regexp"
)

var user string = config.GinVueAdminconfig.Sendsms.User
var password string = config.GinVueAdminconfig.Sendsms.Passwd
var smsKind string = config.GinVueAdminconfig.Sendsms.SmsKind

type SmsMsgType int

const (
	SMSMSGTYPE_VERIFY SmsMsgType = iota
	SMSMSGTYPE_ENTER_CONTRACT
	SMSMSGTYPE_LEAVEL_CONTRACT
	SMSMSGTYPE_INTERVIEW

	SMSMSGTYPE_MAX
)

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func SendSmsMsg(sendToTel, msgContext string) error {
	if !VerifyMobileFormat(sendToTel) || len(msgContext) == 0 {
		fmt.Printf("illdge mobile num please check\n")
		return nil
	}
	msg := "LoginName=" + user
	msg += "&Password="
	msg += md5V(password)[8:24]
	msg += "&SmsKind=" + smsKind
	msg += "&SendSim="
	msg += sendToTel
	msg += "&ExpSmsId="
	msg += "&MsgContext="
	msg += msgContext
	msg += "【亦庄云学堂】"
	fmt.Printf("http msg :%v", msg)

	url := "http://sdk.zyer.cn/SmsService/SmsService.asmx/SendEx"
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(msg))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("http err :%v", err)
		return err
	}
	defer resp.Body.Close()
	statuscode := resp.StatusCode
	hea := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(statuscode)
	fmt.Println(hea)

	return nil
}

func SendSms(sendToTel, userAccount string, msgType SmsMsgType) error {
	if !VerifyMobileFormat(sendToTel) || len(user) == 0 || msgType < SMSMSGTYPE_VERIFY || msgType > SMSMSGTYPE_MAX {
		fmt.Printf("illdge mobile num please check\n")
		return fmt.Errorf("param input err")
	}
	switch msgType {
	case SMSMSGTYPE_VERIFY:
		msgContext := fmt.Sprintf("请输入验证码 %s 完成手机验证（5分钟内有效）。如非本人操作请忽略。回TD退订 ", userAccount)
		err := SendSmsMsg(sendToTel, msgContext)
		return err
	case SMSMSGTYPE_ENTER_CONTRACT:
		msgContext := fmt.Sprintf("亲爱的用户:%s 您好，入职电子合同已生成，请及时阅读确认签字。如非本人操作请忽略。回TD退订 ", userAccount)
		err := SendSmsMsg(sendToTel, msgContext)
		return err
	case SMSMSGTYPE_LEAVEL_CONTRACT:
		msgContext := fmt.Sprintf("亲爱的用户:%s 您好，离职手续已顺利办理完成，祝您工作顺利，下次合作。如非本人操作请忽略。回TD退订 ", userAccount)
		err := SendSmsMsg(sendToTel, msgContext)
		return err
	case SMSMSGTYPE_INTERVIEW:
		msgContext := fmt.Sprintf("亲爱的用户:%s 您好，约您在%s 来面试。如非本人操作请忽略。回TD退订 ", userAccount,)
		err := SendSmsMsg(sendToTel, msgContext)
		return err
	}
	return nil
}

func SendSmsInterview(sendToTel, userAccount ,empName,interviewTime string) error {
	if !VerifyMobileFormat(sendToTel) || len(user) == 0 {
		fmt.Printf("illdge mobile num please check\n")
		return fmt.Errorf("param input err")
	}
	
	msgContext := fmt.Sprintf("亲爱的用户:%s 您好，%s 约您在%s 来公司面试。如非本人操作请忽略。回TD退订 ", userAccount,empName,interviewTime)
	err := SendSmsMsg(sendToTel, msgContext)
	return err
}