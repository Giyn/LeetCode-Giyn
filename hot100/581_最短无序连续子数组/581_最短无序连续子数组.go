/*
-------------------------------------
# @Time    : 2022/3/31 23:44:24
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 581_最短无序连续子数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarray(nums))
}

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	min := nums[n-1]
	max := nums[0]
	begin, end := 0, -1
	for i := 0; i < n; i++ {
		// 从左到右维持最大值，寻找右边界end
		if nums[i] < max {
			end = i
		} else {
			max = nums[i]
		}
		// 从右到左维持最小值，寻找左边界begin
		if nums[n-i-1] > min {
			begin = n - i - 1
		} else {
			min = nums[n-i-1]
		}
	}
	return end - begin + 1
}
