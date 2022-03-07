/*
-------------------------------------
# @Time    : 2022/3/7 9:54:29
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_010_和为 k 的子数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, -1, 0}
	k := 0
	fmt.Println(subarraySum(nums, k))
}

func subarraySum(nums []int, k int) (ans int) {
	pre := 0
	mp := make(map[int]int, len(nums)) // 前缀和的出现次数
	mp[0] = 1                          // 下标从0到i的情况
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if times, ok := mp[pre-k]; ok {
			ans += times
		}
		mp[pre] += 1
	}
	return
}
