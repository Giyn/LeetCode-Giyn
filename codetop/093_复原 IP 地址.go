/*
-------------------------------------
# @Time    : 2022/5/22 19:39:46
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
	isIpAddresses := func(s string) bool {
		if s[0] == '0' && len(s) > 1 {
			return false
		}
		if res, _ := strconv.Atoi(s); res > 255 {
			return false
		}
		return true
	}

	var path []string
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(s) && len(path) == 4 {
			ans = append(ans, strings.Join(path, "."))
			return
		}
		for i := idx; i < len(s); i++ {
			if !isIpAddresses(s[idx : i+1]) {
				return
			}
			path = append(path, s[idx:i+1])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}
