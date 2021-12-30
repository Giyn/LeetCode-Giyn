/*
-------------------------------------
# @Time    : 2021/12/30 10:48:26
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1296_划分数组为连续数字的集合.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 3, 4, 4, 5, 6}
	k := 4
	fmt.Println(isPossibleDivide(nums, k))
}

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}
	sort.Ints(nums)
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}
	for _, x := range nums {
		if v, ok := mp[x]; ok && v > 0 {
			for i := x; i < x+k; i++ {
				if mp[i] == 0 {
					return false
				}
				mp[i]--
			}
		}
	}
	return true
}
