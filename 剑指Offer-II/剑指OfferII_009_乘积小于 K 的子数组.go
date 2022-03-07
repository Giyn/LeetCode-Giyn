/*
-------------------------------------
# @Time    : 2022/3/7 1:01:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_009_乘积小于 K 的子数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{10, 9, 10, 4, 3, 8, 3, 3, 6, 2, 10, 10, 9, 3}
	k := 19
	fmt.Println(numSubarrayProductLessThanK(nums, k))
}

func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	if k <= 1 {
		return 0
	}
	left, right := 0, 0
	sum := 1
	for right < len(nums) {
		sum *= nums[right]
		for sum >= k && left < len(nums) {
			sum /= nums[left]
			left++
		}
		ans += right - left + 1 // 从右往左延伸,左在前进,不会重复,如:[4] [4, 3] [4, 3, 2] [4, 3, 2, 1],已经处理完[3] [3, 2] [3、2、1]
		right++
	}
	return
}
