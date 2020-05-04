package middleware

import (
	"fmt"
	//"github.com/gin-gonic/gin"
	"gin-vue-admin/config"
	"gin-vue-admin/controller/servers"
	"io"
	"sync"
	"time"
	"io/ioutil"
	"net/http"
	"bytes"
	"mime/multipart"
	"encoding/json"
)

type ContentInfo struct {
	Content   	 string `json:"content"`
}

type ReturnErrorInfo struct {
	ErrCode   	 int `json:"errcode"`
	ErrMsg   	 string `json:"errmsg"`
}

type WX_Access struct {
	access_token string
}

var g_tocken WX_Access

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

	var access_token string
	if g_tocken.access_token == "" {
		access_token,err := GetAccessTocken()
		if err != nil{
			return false,err
		}
		g_tocken.access_token = access_token 
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

func WechatDetectImg(bts []byte) (bool,error) {
	var bufReader bytes.Buffer
	mpWriter := multipart.NewWriter(&bufReader)
	//文件名无所谓
	fileName := "detect"
	writer, err := mpWriter.CreateFormFile("media", fileName)
	if err != nil {
		fmt.Println(err.Error())
		return false,err
	}
	reader := bytes.NewReader(bts)
	io.Copy(writer, reader)
	//关闭了才能把内容写入
	mpWriter.Close()
 
	if g_tocken.access_token == ""{
		g_tocken.access_token,err = GetAccessTocken()
		if err != nil{
			return false,err
		}
	}
	
	client := http.DefaultClient
	destURL := "https://api.weixin.qq.com/wxa/img_sec_check?access_token=" + g_tocken.access_token
	req, _ := http.NewRequest("POST", destURL, &bufReader)
	//从mpWriter中获取content-Type
	req.Header.Set("Content-Type","application/octet-stream")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	vs := make(map[string]interface{})
	result, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false,err
	}
	err = json.Unmarshal(result, &vs)
	if err != nil {
		return false, err
	}
	//errcode 存在，且为0，返回通过
	if _, ok := vs["errcode"]; ok && vs["errcode"].(float64) == 0.0 {
		return true,nil
	}
 
	return false, err
}


func NewPatchIncrCluster() *WX_Access {
	return &WX_Access{
		access_token: "",
	}
}

func (wx_a *WX_Access) StartRun(){
	var mtx sync.Mutex
	for{
		token ,err:= GetAccessTocken()
		if err != nil{
			time.Sleep(time.Second)
			continue
		} else{
			mtx.Lock()
			g_tocken.access_token = token
			mtx.Unlock()

			time.Sleep(time.Hour)
		}
	}
}