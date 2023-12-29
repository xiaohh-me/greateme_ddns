package service

import (
	"fmt"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	"github.com/xiaohh-me/greateme_ddns/utils"
	"github.com/xiaohh-me/greateme_ddns/utils/alibaba"
	"log"
	"strings"
)

// SyncAllDomain 同步所有指定的域名到目前的公网IP上
func SyncAllDomain(domainNameList *[]string, dnsType *string) error {
	// 获取公网IP地址
	wanIp, err := utils.GetWanIpAddress(dnsType)
	if err != nil {
		return err
	}
	log.Printf("成功获取当前的公网IP地址：%v\n", *wanIp)
	// 获取所有的域名列表
	domainList, err := alibaba.GetAllDomainList()
	if err != nil {
		return err
	}
	// 遍历所有需要同步DNS的域名，查询到他的二级域名
	for _, domainName := range *domainNameList {
		log.Printf("开始尝试同步域名：%v\n", domainName)
		// 匹配的二级域名
		var level2Domain string
		// 匹配的RR记录
		var rr string
		// 遍历所有域名，找到需要同步的二级域名和rr记录值
		for _, domain := range *domainList {
			// 如果二级域名比三级域名还要长，说明不是这个域名
			if len(domain) > len(domainName) {
				continue
			}
			// 判断后缀是否相等，如果相等那么就找到了这个二级域名
			if strings.HasSuffix(domainName, fmt.Sprintf(".%v", domain)) || strings.Compare(domainName, domain) == 0 {
				level2Domain = domain
				// 判断RR值
				if strings.Compare(domainName, domain) == 0 {
					rr = "@"
				} else {
					rr = strings.TrimSuffix(domainName, fmt.Sprintf(".%v", domain))
				}
				break
			}
		}
		if strings.Compare(level2Domain, "") == 0 || strings.Compare(rr, "") == 0 {
			log.Printf("非常抱歉域名%v可能不属于您，请您确认你的阿里云账户的域名信息！\n", domainName)
			continue
		}
		log.Printf("成功查询到%v域名信息信息，二级域名：%v，rr值：%v\n", domainName, level2Domain, rr)
		dnsList, err := alibaba.GetAllDNSListByDomainNameAndRR(&level2Domain, &rr)
		if err != nil {
			log.Printf("查询%v域名解析记录时候放生错误，错误信息：%v，将继续同步下一个域名\n", domainName, err)
			continue
		}
		var targetRecord *alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord = nil
		// 判断记录类型是否存在
		for _, record := range *dnsList {
			if strings.Compare(*record.Type, "A") == 0 || // IPv4记录类型
				strings.Compare(*record.Type, "AAAA") == 0 || // IPv6记录类型
				strings.Compare(*record.Type, "CNAME") == 0 || // CNAME记录类型
				strings.Compare(*record.Type, "TXT") == 0 { // TXT记录类型
				targetRecord = record
				break
			}
		}
		if targetRecord == nil {
			// 需要新增
			err = alibaba.AddDNSRecord(&level2Domain, &rr, wanIp, dnsType)
			if err != nil {
				log.Printf("新增%v解析的时候发生错误，错误信息：%v\n", domainName, err)
			} else {
				log.Printf("新增%v解析记录成功，解析到IP地址为%v\n", domainName, *wanIp)
			}
		} else if strings.Compare(*targetRecord.Type, *alibaba.GetDNSType(dnsType)) != 0 ||
			strings.Compare(*targetRecord.Value, *wanIp) != 0 ||
			strings.Compare(*targetRecord.Line, "default") != 0 ||
			*targetRecord.TTL != 600 {
			// 需要修改
			err = alibaba.UpdateDNSRecord(targetRecord.RecordId, &rr, wanIp, dnsType)
			if err != nil {
				log.Printf("修改%v解析的时候发生错误，错误信息：%v\n", domainName, err)
			} else {
				log.Printf("修改%v解析记录成功，解析到IP地址：%v，原类型：%v，原记录值：%v\n", domainName, *wanIp, *targetRecord.Type, *targetRecord.Value)
			}
		} else {
			log.Printf("无需修改%v的解析记录，记录值为：%v\n", domainName, *wanIp)
		}
	}
	return nil
}
