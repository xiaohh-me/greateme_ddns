package conf

import (
	"github.com/go-ini/ini"
	"strings"
	"time"
)

// GetConfig 获取配置文件内容
func GetConfig(path *string) (*string, *string, *[]string, *time.Duration, error) {
	config, err := ini.Load(*path)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	// 读取阿里云相关的配置
	accessKeyId, accessKeySecret, err := getAliyunConfig(config)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	// 读取域名相关配置
	domainList, err := getDomainList(config)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	// 获取同步时间
	durationMinute, err := getDurationMinute(config)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	// 读取域名配置
	return accessKeyId, accessKeySecret, domainList, durationMinute, nil
}

// getAliyunConfig 获取阿里云相关配置
func getAliyunConfig(config *ini.File) (*string, *string, error) {
	// 读取阿里云相关的配置
	aliyunSection, err := config.GetSection("aliyun")
	if err != nil {
		return nil, nil, err
	}
	// 读取accessKeyId
	accessKeyIdKey, err := aliyunSection.GetKey("accessKeyId")
	if err != nil {
		return nil, nil, err
	}
	accessKeyId := accessKeyIdKey.Value()
	// 读取accessKeySecret
	accessKeySecretKey, err := aliyunSection.GetKey("accessKeySecret")
	if err != nil {
		return nil, nil, err
	}
	accessKeySecret := accessKeySecretKey.Value()
	return &accessKeyId, &accessKeySecret, nil
}

// getDomainList 获取域名的配置列表
func getDomainList(config *ini.File) (*[]string, error) {
	domainSection, err := config.GetSection("domain")
	if err != nil {
		return nil, err
	}
	// 获取域名列表
	domainListKey, err := domainSection.GetKey("domainList")
	if err != nil {
		return nil, err
	}
	domainListStr := domainListKey.Value()
	// 以逗号隔开获取域名列表
	domainList := strings.Split(domainListStr, ",")
	return &domainList, nil
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
