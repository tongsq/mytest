package common

import (
	"fmt"
	"net"
	"testing"
)

func TestDns(t *testing.T) {
	ipRecords, _ := net.LookupIP("baidu.com")
	for _, ip := range ipRecords {
		fmt.Println(ip)
	}
	cname, _ := net.LookupCNAME("www.baidu.com")
	fmt.Println(cname)

	ptr, err := net.LookupAddr("114.114.114.114")
	if err != nil {
		fmt.Println(err)
	}
	for _, ptrvalue := range ptr {
		fmt.Println(ptrvalue)
	}
}
