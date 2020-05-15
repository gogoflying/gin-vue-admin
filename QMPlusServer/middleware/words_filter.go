package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"

	//"github.com/gin-gonic/gin"
	"bytes"
	"encoding/json"
	"gin-vue-admin/config"
	"gin-vue-admin/controller/servers"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"sync"
	"time"
)

type ContentInfo struct {
	Content string `json:"content"`
}

type ReturnErrorInfo struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type WX_Access struct {
	access_token string
}

var g_tocken WX_Access

var appid string = config.GinVueAdminconfig.WeiXin.ResumeApp.Appid
var appSecret string = config.GinVueAdminconfig.WeiXin.ResumeApp.AppSecret

//敏感词过滤
func GetAccessTocken() (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appid, appSecret)
	body, err := servers.GetEx(url)
	if err != nil {
		return "", fmt.Errorf("param empty")
	}
	var mapResult map[string]interface{}
	err = json.Unmarshal([]byte(body), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}

	return mapResult["access_token"].(string), nil
}

func CheckWordsIllegal(words string) (bool, error) {
	if len(words) == 0 {
		return true, nil
	}
	if len(words) > 500*1024 {
		return false, fmt.Errorf("鉴权文本字节<500k")
	}

	//var access_token string
	fmt.Printf("StartRun 001 a:%v\n", g_tocken.access_token)
	if g_tocken.access_token == "" {
		fmt.Printf("StartRun 002 b:%v\n", g_tocken.access_token)
		access_token, err := GetAccessTocken()
		if err != nil {
			return false, err
		}
		g_tocken.access_token = access_token
	}
	fmt.Printf("StartRun 003 c:%v\n", g_tocken.access_token)
	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%s", g_tocken.access_token)
	con := ContentInfo{
		Content: words,
	}

	body, err := servers.PostEx(url, con)
	if err != nil {
		return false, err
	}
	var returnErr ReturnErrorInfo
	json.Unmarshal([]byte(body), &returnErr)
	fmt.Println(returnErr)
	if returnErr.ErrCode != 0 {
		return false, fmt.Errorf("%v", "禁止上传敏感信息")
	}

	return true, nil
}

func WechatDetectImg(bts []byte) (bool, error) {
	var bufReader bytes.Buffer
	mpWriter := multipart.NewWriter(&bufReader)
	//文件名无所谓
	fileName := "detect"
	writer, err := mpWriter.CreateFormFile("media", fileName)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	reader := bytes.NewReader(bts)
	io.Copy(writer, reader)
	//关闭了才能把内容写入
	mpWriter.Close()

	fmt.Printf("StartRun before 001\n")
	if g_tocken.access_token == "" {
		fmt.Printf("StartRun before 002\n")
		g_tocken.access_token, err = GetAccessTocken()
		if err != nil {
			return false, err
		}
	}
	fmt.Printf("StartRun before 003\n")

	client := http.DefaultClient
	destURL := "https://api.weixin.qq.com/wxa/img_sec_check?access_token=" + g_tocken.access_token
	req, _ := http.NewRequest("POST", destURL, &bufReader)
	//从mpWriter中获取content-Type
	req.Header.Set("Content-Type", "application/octet-stream")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	vs := make(map[string]interface{})
	result, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(result, &vs)
	if err != nil {
		return false, err
	}
	//errcode 存在，且为0，返回通过
	if _, ok := vs["errcode"]; ok && vs["errcode"].(float64) == 0.0 {
		return true, nil
	}

	return false, err
}

func NewPatchIncrCluster() *WX_Access {
	return &WX_Access{
		access_token: "",
	}
}

func (wx_a *WX_Access) StartRun() {
	var mtx sync.Mutex
	for {
		fmt.Printf("StartRun before g_id:%v\n", g_tocken.access_token)
		token, err := GetAccessTocken()
		if err != nil {
			time.Sleep(time.Second)
			continue
		} else {
			mtx.Lock()
			g_tocken.access_token = token
			mtx.Unlock()
			fmt.Printf("StartRun after g_id:%v\n", g_tocken.access_token)
			time.Sleep(time.Hour)
		}
	}
}

func WordFilterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Body == nil {
			c.Next()
		} else {
			//获取body内容
			bodyCopy := new(bytes.Buffer)
			_, err := io.Copy(bodyCopy, c.Request.Body)
			if err == nil {
				bodyData := bodyCopy.Bytes()
				c.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))
			}
			//
			_, err = CheckWordsIllegal(bodyCopy.String())
			if err == nil {
				c.Next()
			} else {
				servers.ReportFormat(c, false, fmt.Sprintf("文本鉴权失败：%v", err), gin.H{})
				c.Abort()
				return
			}
		}
	}
}
