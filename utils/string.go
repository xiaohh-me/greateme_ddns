package utils

import (
	"net"
)

// IsIPv4Address 判断字符串是IPv4或者是IPv6
func IsIPv4Address(ip *string) bool {
	return net.ParseIP(*ip).To4() != nil
}
