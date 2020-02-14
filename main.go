package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	hexTable  = "0123456789abcdef"
	prefix    = flag.String("pre", "", "MAC地址前缀，如：00:11:22")
	separator = flag.String("sep", ":", "分隔符：冒号:，2中横线-，3下划线_，4空格 ")
)

func main() {
	flag.Parse()

	// 前缀处理，最大6位
	mac := *prefix
	prefixLen := len(mac)
	if prefixLen > 0 {
		mac = strings.Replace(*prefix, ":", "", -1)
		mac = strings.Replace(mac, "-", "", -1)
		mac = strings.Replace(mac, " ", "", -1)
		if prefixLen > 6 {
			mac = mac[0:6]
		}
	}

	// 补全位数
	rand.Seed(int64(time.Now().UnixNano()))
	for i := 12 - len(mac); i > 0; i-- {
		mac += hex(rand.Intn(16))
	}

	// 按指定的符合分割
	format := ""
	for i := 0; i < 12; i += 2 {
		format = format + mac[i:i+2] + *separator
	}

	fmt.Println(strings.ToUpper(format[0:17]))
}

func hex(n int) string {
	if n < 10 {
		return strconv.Itoa(n)
	}

	return hexTable[n-1 : n]
}
