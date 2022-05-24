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
	var mQueue []int
	Push := func(i int) {
		for len(mQueue) > 0 && nums[i] > mQueue[len(mQueue)-1] {
			mQueue = mQueue[:len(mQueue)-1]
		}
		mQueue = append(mQueue, nums[i])
	}
	for i := 0; i < k; i++ {
		Push(i)
	}
	ans = append(ans, mQueue[0])
	for i := k; i < len(nums); i++ {
		if len(mQueue) != 0 && nums[i-k] == mQueue[0] {
			mQueue = mQueue[1:]
		}
		Push(i)
		ans = append(ans, mQueue[0])
	}
	return
}
