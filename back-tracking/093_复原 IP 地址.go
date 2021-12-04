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
	var path []string

	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(s) && len(path) == 4 {
			ans = append(ans, strings.Join(path, "."))
			return
		}
		for i := index; i < len(s); i++ {
			if isIpAddresses(s[index : i+1]) {
				path = append(path, s[index:i+1])
				backtrack(i + 1)
			} else {
				return
			}
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return
}

func isIpAddresses(s string) bool {
	if s[0] == '0' && len(s) > 1 {
		return false
	}
	tmp, _ := strconv.Atoi(s)
	if tmp > 255 {
		return false
	}
	return true
}
