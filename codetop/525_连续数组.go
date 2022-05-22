/*
-------------------------------------
# @Time    : 2022/5/22 14:20:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 525_连续数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

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
		if v, ok := mp[pre]; ok {
			if i-v > ans {
				ans = i - v
			}
		} else {
			mp[pre] = i
		}
	}
	return
}
