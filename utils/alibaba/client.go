package alibaba

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	domain20180129 "github.com/alibabacloud-go/domain-20180129/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

// client 操作阿里云域名的客户端
var domainClient *domain20180129.Client

// InitClient 初始化阿里云域名请求的客户端
func InitClient(accessKeyId string, accessKeySecret string) error {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: &accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: &accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("domain.aliyuncs.com")
	result, err := domain20180129.NewClient(config)
	if err != nil {
		return err
	}
	domainClient = result
	return nil
}

// GetDomainClient 获取初始化好的阿里云域名客户端
func GetDomainClient() *domain20180129.Client {
	return domainClient
}
