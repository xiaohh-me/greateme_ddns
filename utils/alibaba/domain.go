package alibaba

import (
	"fmt"
	domain20180129 "github.com/alibabacloud-go/domain-20180129/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/xiaohh-me/greateme_ddns/entity"
	"log"
	"strings"
)

// GetAllDomainList 获得所有的域名列表
func GetAllDomainList() (*[]string, error) {
	// 域名解表的结果
	var domainList []string

	// 当前页码
	var currentPageNum int32 = 1
	for {
		// 请求参数
		queryDomainListRequest := &domain20180129.QueryDomainListRequest{
			PageNum:  &currentPageNum,
			PageSize: tea.Int32(50),
		}
		runtime := &util.RuntimeOptions{}
		// 请求获取域名列表
		request, err := domainClient.QueryDomainListWithOptions(queryDomainListRequest, runtime)
		if err != nil {
			return nil, err
		}
		// 遍历并将所有域名放入切片
		for _, domain := range request.Body.Data.Domain {
			domainList = append(domainList, *domain.DomainName)
		}
		// 判断是否要继续分页
		if int(*request.Body.TotalItemNum) <= cap(domainList) {
			// 查询到现在的域名数量大于等于总域名数量，那么我们就跳出循环
			break
		}
		// 继续循环，页码+1
		currentPageNum++
	}
	return &domainList, nil
}

// GetAllDomainAnalyzeByDomainList 通过待解析的域名列表获取所有域名解析记录详情
func GetAllDomainAnalyzeByDomainList(domainList *[]string) (*[]entity.DomainAnalyze, error) {
	// 获取所有域名列表
	accountDomainList, err := GetAllDomainList()
	if err != nil {
		return nil, err
	}
	// 域名解析结果
	domainAnalyzesList := make([]entity.DomainAnalyze, 0)
	for _, domain := range *domainList {
		if strings.Compare(domain, "") == 0 {
			continue
		}
		// 匹配的二级域名
		var level2Domain string
		// 匹配的RR记录
		var rr string
		// 遍历账户域名列表，找到匹配的二级域名
		for _, accountDomain := range *accountDomainList {
			// 如果二级域名比三级域名还要长，说明不是这个域名
			if len(accountDomain) > len(domain) {
				continue
			}
			// 判断后缀是否相等，如果相等那么就找到了这个二级域名
			if strings.HasSuffix(domain, fmt.Sprintf(".%v", accountDomain)) || strings.Compare(accountDomain, domain) == 0 {
				level2Domain = accountDomain
				// 判断RR值
				if strings.Compare(accountDomain, domain) == 0 {
					rr = "@"
				} else {
					rr = strings.TrimSuffix(domain, fmt.Sprintf(".%v", accountDomain))
				}
				break
			}
		}
		if strings.Compare(level2Domain, "") == 0 || strings.Compare(rr, "") == 0 {
			log.Printf("非常抱歉域名%v可能不属于您，请您确认你的阿里云账户的域名信息！\n", domain)
			continue
		}
		log.Printf("成功查询到%v域名信息信息，二级域名：%v，rr值：%v\n", domain, level2Domain, rr)
		// 拼接查询到的结果
		domainAnalyzesList = append(domainAnalyzesList, entity.DomainAnalyze{
			Domain:   &domain,
			DomainL2: &level2Domain,
			Rr:       &rr,
		})
	}
	return &domainAnalyzesList, nil
}
