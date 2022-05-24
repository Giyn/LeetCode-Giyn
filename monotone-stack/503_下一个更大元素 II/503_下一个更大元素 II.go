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
	// 模拟循环遍历
	n := len(nums)
	ans = make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	var stack []int
	for i := 0; i < 2*n; i++ {
		for len(stack) > 0 && nums[i%n] > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top] = nums[i%n]
		}
		stack = append(stack, i%n)
	}
	return

	// 修改数组
	//n := len(nums)
	//var stack []int
	//ans = make([]int, 2*n)
	//for i := 0; i < n-1; i++ {
	//	nums = append(nums, nums[i])
	//}
	//for i := range ans {
	//	ans[i] = -1
	//}
	//for i := 0; i < len(nums); i++ {
	//	for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
	//		top := stack[len(stack)-1]
	//		stack = stack[:len(stack)-1]
	//		ans[top] = nums[i]
	//	}
	//	stack = append(stack, i)
	//}
	//return ans[:n]
}
