package utils

import (
	"fmt"
	"strings"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

func GetIpInfo(ip string) (ipInfo string, err error) {
	searcher, err := xdb.NewWithFileOnly("utils/ip.xdb")
	if err != nil {
		fmt.Printf("failed to create searcher: %s\n", err.Error())
		return "", err
	}
	defer searcher.Close()
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
		return "", err
	}
	arr := strings.Split(region, "|")
	ipInfo = fmt.Sprintf("%s-%s-%s-%s", arr[0], arr[2], arr[3], arr[4])
	return ipInfo, err
}
