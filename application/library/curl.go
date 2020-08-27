package library

import (
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strings"
)

//Get方法，后续优化
func CurlGet(host, method, params string) ([]byte, error) {
	url := host + method + `?` + params
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}

func CurlPost(host, path, params string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", host+path, strings.NewReader(params))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	Logger.Info("POST请求：",
		zap.String("请求结果", string(body)))
	return body, nil
}
