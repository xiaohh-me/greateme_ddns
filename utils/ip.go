package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// WanIpResponse 请求获取公网IP地址返回的json结果
type WanIpResponse struct {
	Ip string `json:"ip"`
}

// GetWanIpAddress 获取公网IP地址
func GetWanIpAddress(dnsType *string) (*string, error) {
	var requestURL string
	switch *dnsType {
	case "ipv4":
		// 请求获取ipv4地址
		requestURL = "https://ipv4-a.jsonip.com"
	case "ipv6":
		// 请求获取ipv6地址
		requestURL = "https://ipv6-a.jsonip.com"
	}
	// 访问并获取公网IP地址
	resp, err := http.Get(requestURL)
	// 方法执行完后关闭访问的body对象
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("获取IP地址的时候发生错误:%v", err)
		}
	}(resp.Body)
	// 如果发生错误则直接返回错误信息
	if err != nil {
		return nil, err
	}
	// 读取body当中的所有信息
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// 将返回的内容转换为json
	var ipResponse = &WanIpResponse{}
	err = json.Unmarshal(result, ipResponse)
	if err != nil {
		return nil, err
	}
	return &ipResponse.Ip, nil
}
