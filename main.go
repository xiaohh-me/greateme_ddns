package main

import (
	"github.com/alibabacloud-go/tea/tea"
	"github.com/xiaohh-me/greateme_ddns/conf"
	"github.com/xiaohh-me/greateme_ddns/service"
	"github.com/xiaohh-me/greateme_ddns/utils/alibaba"
	"log"
	"os"
	"strings"
	"time"
)

// 域名服务文档地址：https://next.api.aliyun.com/api/Domain/2018-01-29/SaveSingleTaskForCreatingOrderActivate?lang=GO
// 云解析文档地址：https://next.api.aliyun.com/api/Alidns/2015-01-09/AddCustomLine?lang=GO

func main() {
	// 配置文件的路径
	var configFilePath *string
	if len(os.Args) >= 2 {
		// 自定义配置文件路径，读取执行参数的第二个值，也就是下标为1的值
		configFilePath = tea.String(os.Args[1])
	} else {
		// 读取默认配置文件路径
		configFilePath = tea.String("./conf/config.ini")
	}
	// 初始化配置文件
	dnsConfig, err := conf.GetConfig(configFilePath)
	if err != nil {
		log.Fatalf("读取配置文件时候发生错误：%v\n", err)
	}
	// 初始化阿里云域名客户端
	err = alibaba.InitClient(dnsConfig.AccessKeyId, dnsConfig.AccessKeySecret, dnsConfig.DomainEndpoint, dnsConfig.DnsEndpoint)
	if err != nil {
		log.Fatalf("初始化阿里云域名客户端的时候发生了错误：%v\n", err)
	}
	log.Println("域名和DNS解析客户端初始化成功")
	if strings.Compare(*dnsConfig.ExecType, "repetition") == 0 {
		// 多次执行
		for {
			go _main(dnsConfig.DomainList, dnsConfig.DnsType)
			time.Sleep(*dnsConfig.DurationMinute * time.Minute)
		}
	} else if strings.Compare(*dnsConfig.ExecType, "single") == 0 {
		// 单次执行
		_main(dnsConfig.DomainList, dnsConfig.DnsType)
	} else {
		// 执行类型配置错误
		log.Fatalln("执行类型（time.type）配置错误，值只能为single（单次执行）和repetition（多次执行）")
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
