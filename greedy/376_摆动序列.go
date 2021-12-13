/*
-------------------------------------
# @Time    : 2021/12/12 22:52:18
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 376_摆动序列.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	fmt.Println(wiggleMaxLength(nums))
}

func wiggleMaxLength(nums []int) (ans int) {
	if len(nums) <= 1 {
		return len(nums)
	}
	preDiff := 0
	curDiff := 0
	ans = 1
	for i := 0; i < len(nums)-1; i++ {
		curDiff = nums[i+1] - nums[i]
		if (curDiff < 0 && preDiff >= 0) || (curDiff > 0 && preDiff <= 0) {
			ans++
			preDiff = curDiff
		}
	}
	return
}
