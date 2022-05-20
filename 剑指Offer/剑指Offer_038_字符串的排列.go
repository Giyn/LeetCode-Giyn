/*
-------------------------------------
# @Time    : 2022/5/20 17:41:19
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_038_字符串的排列.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "abc"
	fmt.Println(permutation(s))
}

func permutation(s string) (ans []string) {
	sb := []byte(s)
	sort.Slice(sb, func(i, j int) bool {
		return sb[i] < sb[j]
	})
	var path []byte
	vis := make([]bool, len(s))
	var dfs func()
	dfs = func() {
		if len(path) == len(s) {
			ans = append(ans, string(path))
			return
		}
		for i := 0; i < len(s); i++ {
			// !vis[i-1] 对树层去重, 递归已回退到该层
			if vis[i] || (i > 0 && sb[i] == sb[i-1] && !vis[i-1]) {
				continue
			}
			path = append(path, sb[i])
			vis[i] = true
			dfs()
			vis[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs()
	return
}
