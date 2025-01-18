package entity

// DomainAnalyze 域名解析实体
type DomainAnalyze struct {
	// Domain 被解析的域名
	Domain *string
	// DomainL2 二级域名
	DomainL2 *string
	// 解析的RR值
	Rr *string
}
