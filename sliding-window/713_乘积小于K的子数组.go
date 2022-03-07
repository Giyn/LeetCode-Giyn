/*
-------------------------------------
# @Time    : 2022/3/7 1:44:42
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 713_乘积小于K的子数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{10, 5, 2, 6}
	k := 100
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
