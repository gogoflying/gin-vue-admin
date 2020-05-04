package middleware

import (
	"fmt"
	//"github.com/gin-gonic/gin"
	"gin-vue-admin/config"
	"gin-vue-admin/controller/servers"
	//"net/http"
	"encoding/json"
)

type ContentInfo struct {
	Content   	 string `json:"content"`
}

type ReturnErrorInfo struct {
	ErrCode   	 int `json:"errcode"`
	ErrMsg   	 string `json:"errmsg"`
}

var appid string = config.GinVueAdminconfig.WeiXin.Appid
var appSecret string = config.GinVueAdminconfig.WeiXin.AppSecret
//敏感词过滤
func GetAccessTocken() (string,error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",appid,appSecret)
	body,err := servers.GetEx(url)
	if err != nil{
		return "",fmt.Errorf("param empty")
	}
	var mapResult map[string]interface{}
	err = json.Unmarshal([]byte(body), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}

	return mapResult["access_token"].(string),nil
}

func CheckWordsIllegal(words string) (bool,error) {
	if len(words) == 0 || len(words) > 500*1024{
		return false,fmt.Errorf("param empty")
	}
	access_token,err := GetAccessTocken()
	if err != nil{
		return false,err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%s",access_token)
	con := ContentInfo{
		Content:words,
	}

	body,err := servers.PostEx(url,con)
	if err != nil{
		return false,err
	}
	var returnErr ReturnErrorInfo
	json.Unmarshal([]byte(body), &returnErr)
	fmt.Println(returnErr)
	if returnErr.ErrCode != 0{
		return false,fmt.Errorf("Unmarshal err :%v",body)
	}
		
	return true,nil
}