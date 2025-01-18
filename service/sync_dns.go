package service

import (
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	"github.com/xiaohh-me/greateme_ddns/entity"
	"github.com/xiaohh-me/greateme_ddns/utils"
	"github.com/xiaohh-me/greateme_ddns/utils/alibaba"
	"log"
	"strings"
)

// 上一次同步时候的公网IP地址
var preWanIp *string = nil

// SyncAllDomain 同步所有指定的域名到目前的公网IP上
func SyncAllDomain(domainList *[]entity.DomainAnalyze, syncWithNoChange *bool, dnsType, ipType, interfaceName *string) error {
	// 获取公网IP地址
	wanIp, err := utils.GetIpAddress(dnsType, ipType, interfaceName)
	if err != nil {
		return err
	}
	log.Printf("成功获取当前的公网IP地址：%v\n", *wanIp)
	// 判断这一次的IP地址与上一次同步的IP地址是否有变更，如果没有变更则无需同步
	if !*syncWithNoChange && preWanIp != nil && strings.Compare(*preWanIp, *wanIp) == 0 {
		log.Printf("IP地址%s没有变化，无需进行同步\n", *wanIp)
		return nil
	}
	// 遍历所有需要同步DNS的域名
	for _, domain := range *domainList {
		log.Printf("开始尝试同步域名：%v，二级域名：%v，rr值：%v\n", *domain.Domain, *domain.DomainL2, *domain.Rr)
		dnsList, err := alibaba.GetAllDNSListByDomainNameAndRR(domain.DomainL2, domain.Rr)
		if err != nil {
			log.Printf("查询%v域名解析记录时候发生错误，错误信息：%v，将继续同步下一个域名\n", *domain.Domain, err)
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
			err = alibaba.AddDNSRecord(domain.DomainL2, domain.Rr, wanIp, dnsType)
			if err != nil {
				log.Printf("新增%v解析的时候发生错误，错误信息：%v\n", *domain.Domain, err)
			} else {
				log.Printf("新增%v解析记录成功，解析到IP地址为%v\n", *domain.Domain, *wanIp)
			}
		} else if strings.Compare(*targetRecord.Type, *alibaba.GetDNSType(dnsType)) != 0 ||
			strings.Compare(*targetRecord.Value, *wanIp) != 0 ||
			strings.Compare(*targetRecord.Line, "default") != 0 ||
			*targetRecord.TTL != 600 {
			// 需要修改
			err = alibaba.UpdateDNSRecord(targetRecord.RecordId, domain.Rr, wanIp, dnsType)
			if err != nil {
				log.Printf("修改%v解析的时候发生错误，错误信息：%v\n", *domain.Domain, err)
			} else {
				log.Printf("修改%v解析记录成功，解析到IP地址：%v，原类型：%v，原记录值：%v\n", *domain.Domain, *wanIp, *targetRecord.Type, *targetRecord.Value)
			}
		} else {
			log.Printf("无需修改%v的解析记录，记录值为：%v\n", *domain.Domain, *wanIp)
		}
	}
	// 将本次同步的IP地址赋值给上一次的IP
	if !*syncWithNoChange {
		preWanIp = wanIp
	}
	return nil
}
