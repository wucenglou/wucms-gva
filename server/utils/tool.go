package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

func GetIpInfo(ip string) (ipInfo string, err error) {
	searcher, err := xdb.NewWithFileOnly("resource/ip.xdb")
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

func ParseAndFormatTime(strTime string) (time.Time, error) {
	// 定义输入时间字符串和格式
	// input := "2023-4-3 21:24:22"
	inputFormat := "2006-1-2 15:04:05"
	outputFormat := "2006-01-02T15:04:05.000Z"

	// 解析输入时间字符串为time.Time类型的值
	t, err := time.Parse(inputFormat, strTime)
	if err != nil {
		t, err = time.Parse(inputFormat, strTime+" 00:00:00")
		if err != nil {
			return time.Time{}, err
		}
	}

	// 将时间转换为UTC时间
	utc := t.UTC().Add(-8 * time.Hour)

	// 将时间格式化为目标字符串格式
	output := utc.Format(outputFormat)

	// 将格式化后的字符串转换为time.Time类型的值
	t2, err := time.Parse(outputFormat, output)
	if err != nil {
		return time.Time{}, err
	}

	// 输出结果
	fmt.Println(t2)
	return t2, nil
}
