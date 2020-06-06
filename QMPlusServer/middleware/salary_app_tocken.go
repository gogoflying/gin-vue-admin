package middleware

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	//"github.com/gin-gonic/gin"
	"bytes"
	"encoding/json"
	"gin-vue-admin/config"

	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/sysModel"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"sync"
	"time"
)

type ContentInfo_Sa struct {
	Content string `json:"content"`
}

type ReturnErrorInfo_Sa struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type WX_Access_Sa struct {
	access_token_sa string
}

var g_tocken_sa WX_Access_Sa
var mtx_sa sync.Mutex

var salaryAppId string = config.GinVueAdminconfig.WeiXin.SalaryApp.Appid
var salaryAppSecretId string = config.GinVueAdminconfig.WeiXin.SalaryApp.AppSecret

const salaryAppName = "salary-app"

//敏感词过滤
func GetAccessTocken_Sa(appId, secretId string) (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, secretId)
	body, err := servers.GetEx(url)
	if err != nil {
		return "", fmt.Errorf("param empty")
	}
	var mapResult map[string]interface{}
	err = json.Unmarshal([]byte(body), &mapResult)
	if err != nil {
		return "", err
	}
	if mapResult["access_token"] == nil {
		//fmt.Printf("salary result is empty %s--%s\n", salaryAppId, salaryAppSecretId)
		return "", fmt.Errorf("result is empty")
	}

	return mapResult["access_token"].(string), nil
}

func CheckWordsIllegal_Sa(words string) (bool, error) {
	if len(words) == 0 {
		return true, nil
	}
	if len(words) > 500*1024 {
		return false, fmt.Errorf("鉴权文本字节<500k")
	}

	//var access_token_re string
	if g_tocken_sa.access_token_sa == "" {
		fmt.Printf("access tocken illegal")
		return false, fmt.Errorf("access tocken illegal")
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%s", g_tocken_sa.access_token_sa)
	con := ContentInfo_Sa{
		Content: words,
	}

	body, err := servers.PostEx(url, con)
	if err != nil {
		return false, err
	}
	var returnErr ReturnErrorInfo_Sa
	json.Unmarshal([]byte(body), &returnErr)
	if returnErr.ErrCode != 0 {
		RefreshTockenFromDB_Sa()
		return false, fmt.Errorf("%v", "禁止上传敏感信息")
	}

	return true, nil
}

func WechatDetectImg_Sa(file *multipart.FileHeader) (bool, error) {
	var bufReader bytes.Buffer
	b, err := file.Open()
	if err != nil {
		log.Fatal("copying contents:", err)
	}
	_, err = io.Copy(&bufReader, b) //复制文件中的内容到b中
	if err != nil {
		log.Fatal("copying contents:", err)
	}

	if g_tocken_sa.access_token_sa == "" {
		fmt.Printf("access tocken illegal")
		return false, fmt.Errorf("access tocken illegal")
	}

	client := http.DefaultClient
	destURL := "https://api.weixin.qq.com/wxa/img_sec_check?access_token=" + g_tocken_sa.access_token_sa
	req, _ := http.NewRequest("POST", destURL, &bufReader)
	//从mpWriter中获取content-Type
	req.Header.Set("Content-Type", "application/octet-stream")
	resp, err := client.Do(req)
	if err != nil {
		return false, err
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
	} else {
		RefreshTockenFromDB_Sa()
	}

	return false, err
}

func NewFilterTocken_Sa() *WX_Access_Sa {
	return &WX_Access_Sa{
		access_token_sa: "",
	}
}

func (wx_a *WX_Access_Sa) StartRun() {
	var tocken string
	//start Run write new tocken
	RefreshTockenFromDB_Sa()
	for {
		hfirst := sysModel.GetWxFliterTockenHndle(salaryAppName, "")
		tocken = hfirst.GetNewFilterTocken(salaryAppName)
		if tocken == "" {
			ntoken, err := GetAccessTocken_Sa(salaryAppId, salaryAppSecretId)
			if err != nil {
				time.Sleep(time.Second)
				continue
			} else {
				mtx_sa.Lock()
				g_tocken_sa.access_token_sa = ntoken
				mtx_sa.Unlock()
				h := sysModel.GetWxFliterTockenHndle(salaryAppName, tocken)
				h.UpdateFilterTocken(salaryAppName, ntoken)

				time.Sleep(time.Hour)
			}
		} else {
			err := checkTockenExpired_Sa(tocken)
			if err != nil {
				fmt.Printf("tocken :%s had expired,need refresh:%v", tocken, err)
				for {
					tocken, err = GetAccessTocken_Sa(salaryAppId, salaryAppSecretId)
					if err != nil {
						time.Sleep(time.Second)
						continue
					} else {
						break
					}
				}
				h := sysModel.GetWxFliterTockenHndle(salaryAppName, tocken)
				h.UpdateFilterTocken(salaryAppName, tocken)
				fmt.Printf("new tocken :%s", tocken)
			}
			mtx_sa.Lock()
			g_tocken_sa.access_token_sa = tocken
			mtx_sa.Unlock()

			time.Sleep(time.Hour)
		}

	}
}

func WordFilterHandler_Sa() gin.HandlerFunc {
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
			_, err = CheckWordsIllegal_Sa(bodyCopy.String())
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

func checkTockenExpired_Sa(tocken string) error {
	mtx_sa.Lock()
	g_tocken_sa.access_token_sa = tocken
	mtx_sa.Unlock()
	_, err := CheckWordsIllegal_Sa("test content")
	return err
}
func RefreshTockenFromDB_Sa() {
	mtx_sa.Lock()
	for {
		ntoken, err := GetAccessTocken_Sa(salaryAppId, salaryAppSecretId)
		if err != nil {
			time.Sleep(time.Second * 5)
			continue
		} else {
			g_tocken_sa.access_token_sa = ntoken
			break
		}
	}
	mtx_sa.Unlock()
	h := sysModel.GetWxFliterTockenHndle(salaryAppName, g_tocken_sa.access_token_sa)
	h.UpdateFilterTocken(salaryAppName, g_tocken_sa.access_token_sa)
	fmt.Printf("RefreshTockenFromDB_Sa Refresh tocken success")
}
