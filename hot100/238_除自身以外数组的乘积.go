/*
-------------------------------------
# @Time    : 2022/3/30 12:59:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 238_除自身以外数组的乘积.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) (ans []int) {
	n := len(nums)
	ans = make([]int, n)
	for i := range ans {
		ans[i] = 1
	}
	left, right := 1, 1 // 左边乘积和右边乘积
	for i := 0; i < n; i++ {
		ans[i] *= left
		left *= nums[i]
		ans[n-1-i] *= right
		right *= nums[n-1-i]
	}
	return
}
