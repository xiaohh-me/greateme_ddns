package conf

import (
	"errors"
	"fmt"
	"github.com/go-ini/ini"
	"strings"
	"time"
)

// DnsConfig 动态DNS的解析配置实体
type DnsConfig struct {
	// AccessKeyId 阿里云的AccessKeyId
	AccessKeyId *string
	// AccessKeySecret 阿里云的AccessKeySecret
	AccessKeySecret *string
	// DomainEndpoint 域名的Endpoint
	DomainEndpoint *string
	// DnsEndpoint dns的Endpoint
	DnsEndpoint *string
	// DomainList 需要被解析的域名列表
	DomainList *[]string
	// DnsType 解析类型，只能是 ipv4 和 ipv6 （注意全部小写且不能为大写）
	DnsType *string
	// DurationMinute 时隔多久同步一次域名解析，单位为分钟
	DurationMinute *time.Duration
}

// GetConfig 获取配置文件内容
func GetConfig(path *string) (*DnsConfig, error) {
	config, err := ini.Load(*path)
	if err != nil {
		return nil, err
	}
	// 读取阿里云相关的配置
	accessKeyId, accessKeySecret, domainEndpoint, dnsEndpoint, err := getAliyunConfig(config)
	if err != nil {
		return nil, err
	}
	// 读取域名相关配置
	domainList, dnsType, err := getDomainList(config)
	if err != nil {
		return nil, err
	}
	// 判断解析类型配置是否正确
	if strings.Compare("ipv4", *dnsType) != 0 &&
		strings.Compare("ipv6", *dnsType) != 0 {
		// 如果配置的既不是ipv4也不是ipv6，那么返回一个错误
		return nil, errors.New(fmt.Sprintf("IP地址解析类型错误，请填写ipv4或ipv6（且只能填写小写）！您填写的值为：%v", *dnsType))
	}
	// 获取同步时间
	durationMinute, err := getDurationMinute(config)
	if err != nil {
		return nil, err
	}
	// 读取域名配置
	return &DnsConfig{
		// AccessKeyId 阿里云的AccessKeyId
		AccessKeyId: accessKeyId,
		// AccessKeySecret 阿里云的AccessKeySecret
		AccessKeySecret: accessKeySecret,
		// DomainEndpoint 域名的Endpoint
		DomainEndpoint: domainEndpoint,
		// DnsEndpoint dns的Endpoint
		DnsEndpoint: dnsEndpoint,
		// DomainList 需要被解析的域名列表
		DomainList: domainList,
		// DnsType 解析类型，只能是 ipv4 和 ipv6 （注意全部小写且不能为大写）
		DnsType: dnsType,
		// DurationMinute 时隔多久同步一次域名解析，单位为分钟
		DurationMinute: durationMinute,
	}, nil
}

// getAliyunConfig 获取阿里云相关配置
func getAliyunConfig(config *ini.File) (*string, *string, *string, *string, error) {
	// 读取阿里云相关的配置
	aliyunSection, err := config.GetSection("aliyun")
	if err != nil {
		return nil, nil, nil, nil, err
	}
	// 读取accessKeyId
	accessKeyIdKey, err := aliyunSection.GetKey("accessKeyId")
	if err != nil {
		return nil, nil, nil, nil, err
	}
	accessKeyId := accessKeyIdKey.Value()
	// 读取accessKeySecret
	accessKeySecretKey, err := aliyunSection.GetKey("accessKeySecret")
	if err != nil {
		return nil, nil, nil, nil, err
	}
	accessKeySecret := accessKeySecretKey.Value()
	// 读取域名的Endpoint
	domainEndpointKey, err := aliyunSection.GetKey("domainEndpoint")
	if err != nil {
		return nil, nil, nil, nil, err
	}
	domainEndpoint := domainEndpointKey.Value()
	// 读取dns的Endpoint
	dnsEndpointKey, err := aliyunSection.GetKey("dnsEndpoint")
	if err != nil {
		return nil, nil, nil, nil, err
	}
	dnsEndpoint := dnsEndpointKey.Value()
	return &accessKeyId, &accessKeySecret, &domainEndpoint, &dnsEndpoint, nil
}

// getDomainList 获取域名的配置列表
func getDomainList(config *ini.File) (*[]string, *string, error) {
	domainSection, err := config.GetSection("domain")
	if err != nil {
		return nil, nil, err
	}
	// 获取域名列表
	domainListKey, err := domainSection.GetKey("domainList")
	if err != nil {
		return nil, nil, err
	}
	domainListStr := domainListKey.Value()
	// 以逗号隔开获取域名列表
	domainList := strings.Split(domainListStr, ",")
	// 获取解析类型，确认解析类型是ipv4或ipv6
	dnsTypeKey, err := domainSection.GetKey("dnsType")
	if err != nil {
		return nil, nil, err
	}
	dnsType := dnsTypeKey.Value()
	return &domainList, &dnsType, nil
}

// getDurationMinute 获取同步时间
func getDurationMinute(config *ini.File) (*time.Duration, error) {
	timeSection, err := config.GetSection("time")
	if err != nil {
		return nil, err
	}
	durationMinuteKey, err := timeSection.GetKey("durationMinute")
	if err != nil {
		return nil, err
	}
	durationMinute, err := durationMinuteKey.Int64()
	if err != nil {
		return nil, err
	}
	duration := time.Duration(durationMinute)
	return &duration, err
}
