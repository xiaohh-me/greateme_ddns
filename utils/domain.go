package utils

import (
	"errors"
	"fmt"
	"strings"
)

// CheckDomainRepeat 检查域名列表是否重复，如果重复则返回异常，异常内容为重复的域名名字
func CheckDomainRepeat(domainList *[]string) error {
	domainListLength := len(*domainList) // 域名列表长度
	if domainListLength > 1 {
		for i := 0; i < domainListLength-1; i++ {
			for j := i + 1; j < domainListLength; j++ {
				if strings.Compare((*domainList)[i], (*domainList)[j]) == 0 {
					return errors.New(fmt.Sprintf("%s域名重复了\n", (*domainList)[i]))
				}
			}
		}
	}
	return nil
}
