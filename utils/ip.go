package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

// WanIpResponse 请求获取公网IP地址返回的json结果
type WanIpResponse struct {
	Ip string `json:"ip"`
}

func GetIpAddress(dnsType, ipType, interfaceName *string) (*string, error) {
	if strings.Compare("wan", *ipType) == 0 {
		return getWanIpAddress(dnsType)
	} else {
		if interfaceName == nil || strings.Compare("", *interfaceName) == 0 {
			log.Fatalln("网卡名字为空")
		}
		// 获取所有网卡列表
		interfaces, err := net.Interfaces()
		if err != nil {
			return nil, err
		}
		// 遍历网卡列表，获取名字相同的网卡
		// 目标网卡名字
		var targetInterface *net.Interface = nil
		for _, currentInterface := range interfaces {
			if strings.Compare(*interfaceName, currentInterface.Name) == 0 {
				targetInterface = &currentInterface
				break
			}
		}
		// 如果目标网卡为空，则代表没有找到这个网卡
		if targetInterface == nil {
			log.Fatalf("系统中未找到名字为“%v”的网卡\n", *interfaceName)
		}
		// 获取网卡地址列表
		addressList, err := targetInterface.Addrs()
		if err != nil {
			return nil, err
		}
		// 遍历IP地址列表
		var targetIp *string = nil
		switch *dnsType {
		case "ipv4":
			// 获取网卡的ipv4地址
			for _, address := range addressList {
				ip, _, err := net.ParseCIDR(address.String())
				if err != nil {
					continue
				}
				if ip.To4() != nil && !ip.IsLinkLocalUnicast() {
					targetIp = tea.String(ip.String())
				}
			}
		case "ipv6":
			// 获取网卡的ipv6地址
			for _, address := range addressList {
				ip, _, err := net.ParseCIDR(address.String())
				if err != nil {
					continue
				}
				if ip.To4() == nil && !ip.IsLinkLocalUnicast() {
					targetIp = tea.String(ip.String())
				}
			}
		}
		// 判断获取的IP地址是否为空
		if targetIp == nil {
			return nil, errors.New(fmt.Sprintf("在网卡%s上没有找到正确的%s地址", *interfaceName, *dnsType))
		}
		return targetIp, nil
	}
}

// getWanIpAddress 获取公网IP地址
func getWanIpAddress(dnsType *string) (*string, error) {
	var requestURL string
	requestURL = "https://" + *dnsType + ".jsonip.com"
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
