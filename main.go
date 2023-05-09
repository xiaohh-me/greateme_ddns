package main

import (
	"fmt"
	"github.com/xiaohh-me/greateme_ddns/utils/alibaba"
	"log"
	"os"
	"strings"
)

const (
	// AccessKeyId 阿里云的AccessKey
	AccessKeyId = "AccessKeyId"
	// AccessKeySecret 阿里云的AccessKeySecret
	AccessKeySecret = "AccessKeySecret"
	// SyncDomainList 需要同步的域名列表，用逗号隔开
	SyncDomainList = "home.example.com,company.example.com"
)

func main() {
	// 初始化阿里云域名客户端
	err := alibaba.InitClient(AccessKeyId, AccessKeySecret)
	if err != nil {
		log.Fatalf("初始化阿里云域名客户端的时候发生了错误：%v\n", err)
		os.Exit(1)
	}
	log.Println("域名和DNS解析客户端初始化成功")
	// 获取所有的域名列表
	domainList, err := alibaba.GetAllDomainList()
	if err != nil {
		log.Fatalf("查询域名列表时候发生错误")
	}
	for i, domain := range *domainList {
		fmt.Printf("row: %v, domain name: %v\n", i, domain)
	}
	domainName := "yueyang.city"
	dnsList, err := alibaba.GetAllDNSList(&domainName)
	if err != nil {
		log.Fatalf("查询DNS解析列表时候发生错误")
	}
	for i, record := range *dnsList {
		fmt.Printf("row: %v, domain name: %v, line: %v, rr: %v, type: %v, value: %v\n", i+1, *record.DomainName, *record.Line, *record.RR, *record.Type, *record.Value)
	}
	// 拆分并获取需要同步的域名列表
	syncDomainList := strings.Split(SyncDomainList, ",")
	for i, domain := range syncDomainList {
		fmt.Printf("row: %v, domain name: %v\n", i, domain)
	}
}
