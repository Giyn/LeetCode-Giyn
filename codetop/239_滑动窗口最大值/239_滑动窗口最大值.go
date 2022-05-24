/*
-------------------------------------
# @Time    : 2021/11/10 9:28:36
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 239_滑动窗口最大值.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}

func maxSlidingWindow(nums []int, k int) (ans []int) {
	var queue []int
	push := func(i int) {
		for len(queue) > 0 && nums[i] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	ans = append(ans, queue[0])
	for i := k; i < len(nums); i++ {
		if len(queue) > 0 && nums[i-k] == queue[0] {
			queue = queue[1:]
		}
		push(i)
		ans = append(ans, queue[0])
	}
	return
}
