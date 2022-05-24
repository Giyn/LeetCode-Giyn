/*
-------------------------------------
# @Time    : 2022/5/2 14:02:26
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 396_旋转函数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 3, 2, 6}
	fmt.Println(maxRotateFunction(nums))
}

func maxRotateFunction(nums []int) (ans int) {
	// F(k)=F(k−1)+Sum−n×nums[n−k]
	sum := 0
	n := len(nums)
	f := 0
	for _, num := range nums {
		sum += num
	}
	for i, num := range nums {
		f += i * num
	}
	ans = f
	for i := 1; i < n; i++ {
		f += sum - n*nums[n-i]
		ans = max(ans, f)
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
