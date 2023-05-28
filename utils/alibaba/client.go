package alibaba

import (
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	domain20180129 "github.com/alibabacloud-go/domain-20180129/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

// client 操作阿里云域名的客户端
var (
	// domainClient 域名操作的客户端
	domainClient *domain20180129.Client
	// dnsClient 域名DNS解析操作的客户端
	dnsClient *alidns20150109.Client
)

// InitClient 初始化阿里云域名请求的客户端
func InitClient(accessKeyId *string, accessKeySecret *string) error {
	// 初始化域名配置
	domainConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	domainConfig.Endpoint = tea.String("domain.aliyuncs.com")
	domainResult, err := domain20180129.NewClient(domainConfig)
	if err != nil {
		return err
	}
	domainClient = domainResult

	// 初始化DNS解析配置
	dnsConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	dnsConfig.Endpoint = tea.String("alidns.cn-hangzhou.aliyuncs.com")
	dnsResult, err := alidns20150109.NewClient(dnsConfig)
	if err != nil {
		return err
	}
	dnsClient = dnsResult
	return nil
}

// GetDomainClient 获取初始化好的阿里云域名客户端
func GetDomainClient() *domain20180129.Client {
	return domainClient
}

// GetDNSClient 获取初始化好的阿里云DNS解析客户端
func GetDNSClient() *alidns20150109.Client {
	return dnsClient
}
