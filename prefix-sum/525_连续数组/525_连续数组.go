/*
-------------------------------------
# @Time    : 2022/3/7 14:58:18
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 525_连续数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 1}
	fmt.Println(findMaxLength(nums))
}

func findMaxLength(nums []int) (ans int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}
	pre := 0
	mp := make(map[int]int, len(nums))
	mp[0] = -1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		// 从第0位开始数,到第i位时,0比1多3个;到第j位时,0也比1多3个,这中间的个数必然相等
		if v, ok := mp[pre]; ok {
			ans = max(ans, i-v)
		} else {
			mp[pre] = i
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
