/*
-------------------------------------
# @Time    : 2022/8/27 11:42:11
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 152_乘积最大子数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}
	maxRes, curMax, curMin := nums[0], nums[0], nums[0]
	for i := 1; i < size; i++ {
		if nums[i] < 0 {
			curMax, curMin = curMin, curMax
		}

		curMax = max(curMax*nums[i], nums[i])
		curMin = min(curMin*nums[i], nums[i])
		maxRes = max(maxRes, curMax)
	}
	return maxRes
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
