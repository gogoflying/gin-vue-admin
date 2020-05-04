package servers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"io"
)

func Post(req interface{}, url string, rsp interface{}) ([]byte, error) {
	reqByt, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	rq, err := http.NewRequest("POST", url, bytes.NewBuffer(reqByt))
	if err != nil {
		return nil, err
	}
	rq.Header.Set("Content-Type", "application/json")
	rq.Header.Set("Connection", "keep-alive")

	body, err := client.Do(rq)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(body.Body)
	if err != nil {
		return nil, err
	}
	if rsp != nil {
		if err := json.Unmarshal(b, rsp); err != nil {
			return nil, err
		}
	}
	return b, nil
}

func Get(url string) (body *http.Response, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Host", "jandan.net")
	req.Header.Set("Connection", "keep-alive")
	return client.Do(req)
}

func PostEx(url string, data interface{}) (string ,error) {
    jsonStr, _ := json.Marshal(data)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Connection", "keep-alive")
    if err != nil {
        return "",err
    }
    defer req.Body.Close()

    client := &http.Client{Timeout: 5 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return "",err
    }
    defer resp.Body.Close()

    result, _ := ioutil.ReadAll(resp.Body)
    content := string(result)
    return content,nil
}

func GetEx(url string) (string,error) {
    client := http.Client{Timeout: 5 * time.Second}
    resp, err := client.Get(url)
    defer resp.Body.Close()
    if err != nil {
        return "",err
    }

    var buffer [1024]byte
    result := bytes.NewBuffer(nil)
    for {
        n, err := resp.Body.Read(buffer[0:])
        result.Write(buffer[0:n])
        if err != nil && err == io.EOF {
            break
        } else if err != nil {
            return "",err
        }
    }

    rsp := result.String()
    return rsp,nil
}