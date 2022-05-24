/*
-------------------------------------
# @Time    : 2021/12/4 11:15:17
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 093_复原 IP 地址.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}

func restoreIpAddresses(s string) (ans []string) {
	if len(s) > 4*3 {
		return
	}
	var path []string
	isIpAddresses := func(s string) bool {
		if s[0] == '0' && len(s) > 1 {
			return false
		}
		if tmp, _ := strconv.Atoi(s); tmp > 255 {
			return false
		}
		return true
	}
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(s) && len(path) == 4 {
			ans = append(ans, strings.Join(path, "."))
		}
		for i := index; i < len(s); i++ {
			if !isIpAddresses(s[index : i+1]) {
				return
			}
			path = append(path, s[index:i+1])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}
