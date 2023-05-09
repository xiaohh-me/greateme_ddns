package alibaba

import (
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// GetAllDNSList 根据域名获取所有的DNS解析列表
func GetAllDNSList(domainName *string) (*[]*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord, error) {
	// 最终结果
	var dnsList []*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord

	// 当前页码
	var currentPageNum int64 = 1

	for {
		// 初始化查询参数
		describeDomainRecordsRequest := &alidns20150109.DescribeDomainRecordsRequest{
			DomainName: domainName,
			PageNumber: &currentPageNum,
			PageSize:   tea.Int64(10),
		}
		runtime := &util.RuntimeOptions{}
		// 查询域名列表
		dnsResult, err := dnsClient.DescribeDomainRecordsWithOptions(describeDomainRecordsRequest, runtime)
		if err != nil {
			return nil, err
		}
		dnsList = append(dnsList, dnsResult.Body.DomainRecords.Record...)

		// 判断是否要继续分页
		if int(*dnsResult.Body.TotalCount) <= len(dnsList)+1 {
			// 查询到现在的域名数量大于等于总域名数量，那么我们就跳出循环
			break
		}
		// 继续循环，页码+1
		currentPageNum++
	}
	return &dnsList, nil
}
