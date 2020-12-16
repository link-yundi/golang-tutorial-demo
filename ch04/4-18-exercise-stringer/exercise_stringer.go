package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
通过让 IPAddr 类型实现fmt.Stringer来打印点好分隔的地址
 */

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var ipParts []string
	for _, item := range ip {
		ipParts = append(ipParts, strconv.Itoa(int(item)))
	}
	return strings.Join(ipParts, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}