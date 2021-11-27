/*
-------------------------------------
# @Time    : 2021/10/26 8:43:33
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 503_下一个更大元素 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(nextGreaterElements(nums))
}

func nextGreaterElements(nums []int) (ans []int) {
	n := len(nums)
	var stack []int

	for i := 0; i < n-1; i++ {
		nums = append(nums, nums[i])
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, stack[len(stack)-1])
		}
		stack = append(stack, nums[i])
	}
	return reverse(ans)[:n]
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
