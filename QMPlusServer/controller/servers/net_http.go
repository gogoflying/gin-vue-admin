package servers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

