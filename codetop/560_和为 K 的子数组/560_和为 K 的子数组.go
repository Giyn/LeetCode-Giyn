/*
-------------------------------------
# @Time    : 2022/3/7 11:51:01
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 560_和为 K 的子数组.go
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
	mp := make(map[int]int, len(nums))
	mp[0] = 1 // 从0到最后的情况
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if v, ok := mp[pre-k]; ok {
			ans += v
		}
		mp[pre]++
	}
	return
}
