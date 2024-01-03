package main

import (
	"github.com/alibabacloud-go/tea/tea"
	"github.com/xiaohh-me/greateme_ddns/conf"
	"github.com/xiaohh-me/greateme_ddns/service"
	"github.com/xiaohh-me/greateme_ddns/utils/alibaba"
	"log"
	"time"
)

// 域名服务文档地址：https://next.api.aliyun.com/api/Domain/2018-01-29/SaveSingleTaskForCreatingOrderActivate?lang=GO
// 云解析文档地址：https://next.api.aliyun.com/api/Alidns/2015-01-09/AddCustomLine?lang=GO

func main() {
	dnsConfig, err := conf.GetConfig(tea.String("./conf/config.ini"))
	if err != nil {
		log.Fatalf("读取配置文件时候发生错误：%v\n", err)
	}
	// 初始化阿里云域名客户端
	err = alibaba.InitClient(dnsConfig.AccessKeyId, dnsConfig.AccessKeySecret, dnsConfig.DomainEndpoint, dnsConfig.DnsEndpoint)
	if err != nil {
		log.Fatalf("初始化阿里云域名客户端的时候发生了错误：%v\n", err)
	}
	log.Println("域名和DNS解析客户端初始化成功")
	for {
		go _main(dnsConfig.DomainList, dnsConfig.DnsType)
		time.Sleep(*dnsConfig.DurationMinute * time.Minute)
	}
}

func _main(domainNameList *[]string, dnsType *string) {
	// 捕捉所有异常，兜底的方法
	defer func() {
		if err := recover(); err != nil {
			log.Printf("系统发生了异常：%v\n", err)
		}
	}()

	// 开始同步
	err := service.SyncAllDomain(domainNameList, dnsType)
	if err != nil {
		log.Printf("同步域名信息的时候发生了异常：%v\n", err)
	}
}
