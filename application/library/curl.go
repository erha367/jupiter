package library

import (
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//Get方法，后续优化
func CurlGet(host, method, params string) ([]byte, error) {
	//错误恢复
	defer func() {
		if err := recover(); err != nil {
			Logger.Error(`Curl err : `, zap.Any(`msg`, err))
		}
	}()
	client := &http.Client{Timeout: time.Second}
	var url string
	if params == `` {
		url = host + method
	} else {
		url = host + method + `?` + params
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	Logger.Debug("GET请求：", zap.String("请求结果", string(body)), zap.String("host", host), zap.String("path", method), zap.String("params", params))
	return body, err
}

func CurlPost(host, path, params string) ([]byte, error) {
	//错误恢复
	defer func() {
		if err := recover(); err != nil {
			Logger.Error(`Curl err : `, zap.Any(`msg`, err))
		}
	}()
	client := &http.Client{Timeout: time.Second}
	var resp *http.Response
	//使用熔断器，不用降级
	req, err := http.NewRequest("POST", host+path, strings.NewReader(params))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	Logger.Debug("POST请求：", zap.String("请求结果", string(body)), zap.String("host", host), zap.String("path", path), zap.String("params", params))
	return body, nil
}

func CurlPostjson(host, path, params string) ([]byte, error) {
	//错误恢复
	defer func() {
		if err := recover(); err != nil {
			Logger.Error(`Curl err : `, zap.Any(`msg`, err))
		}
	}()
	client := &http.Client{Timeout: time.Second}
	url := host + path
	payload := strings.NewReader(params)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	Logger.Debug("POST请求：", zap.String("请求结果", string(body)), zap.String("host", host), zap.String("path", path), zap.String("params", params))
	return body, nil
}
