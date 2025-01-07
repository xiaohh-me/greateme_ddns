package alibaba

import (
	domain20180129 "github.com/alibabacloud-go/domain-20180129/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
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
