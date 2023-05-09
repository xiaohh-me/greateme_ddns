package main

import (
	"github.com/xiaohh-me/greateme_ddns/service"
	"github.com/xiaohh-me/greateme_ddns/utils/alibaba"
	"log"
	"os"
	"strings"
	"time"
)

const (
	// AccessKeyId 阿里云的AccessKey
	AccessKeyId = "AccessKeyId"
	// AccessKeySecret 阿里云的AccessKeySecret
	AccessKeySecret = "AccessKeySecret"
	// SyncDomainList 需要同步的域名列表，用逗号隔开
	SyncDomainList = "home.example.com,company.example.com"
	// DurationMinute 每多少分钟同步一次
	DurationMinute time.Duration = 10
)

// 域名服务文档地址：https://next.api.aliyun.com/api/Domain/2018-01-29/SaveSingleTaskForCreatingOrderActivate?lang=GO
// 云解析文档地址：https://next.api.aliyun.com/api/Alidns/2015-01-09/AddCustomLine?lang=GO

func main() {
	// 初始化阿里云域名客户端
	err := alibaba.InitClient(AccessKeyId, AccessKeySecret)
	if err != nil {
		log.Fatalf("初始化阿里云域名客户端的时候发生了错误：%v\n", err)
		os.Exit(1)
	}
	log.Println("域名和DNS解析客户端初始化成功")
	// 拆分并获取需要同步的域名列表
	syncDomainList := strings.Split(SyncDomainList, ",")
	for {
		go _main(&syncDomainList)
		time.Sleep(DurationMinute * time.Minute)
	}
}

func _main(domainNameList *[]string) {
	// 捕捉所有异常，兜底的方法
	defer func() {
		if err := recover(); err != nil {
			log.Printf("系统发生了异常：%v\n", err)
		}
	}()

	// 开始同步
	err := service.SyncAllDomain(domainNameList)
	if err != nil {
		log.Printf("同步域名信息的时候发生了异常：%v\n", err)
	}
}
