/*
-------------------------------------
# @Time    : 2022/4/1 17:38:16
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 954_二倍数对数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-5, -2}
	fmt.Println(canReorderDoubled(arr))
}

func canReorderDoubled(arr []int) bool {
	mp := make(map[int]int, len(arr))
	for _, v := range arr {
		mp[v]++
	}
	keys := make([]int, len(arr))
	for k := range mp {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		if key > 0 {
			if mp[key] > mp[2*key] {
				return false
			}
			mp[2*key] -= mp[key]
		} else if key < 0 {
			if mp[key] > 0 && (key%2 != 0 || mp[key] > mp[key/2]) {
				return false
			}
			mp[key/2] -= mp[key]
		} else {
			if mp[key]%2 == 1 {
				return false
			}
		}
	}
	return true
}
