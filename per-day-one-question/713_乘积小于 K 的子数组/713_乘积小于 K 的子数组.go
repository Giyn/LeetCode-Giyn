/*
-------------------------------------
# @Time    : 2022/5/5 19:58:47
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 713_乘积小于 K 的子数组.go
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
	product := 1
	for left, right := 0, 0; right < len(nums); right++ {
		product *= nums[right]
		for left <= right && product >= k {
			product /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return
}
