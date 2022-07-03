/*
-------------------------------------
# @Time    : 2022/7/3 18:13:38
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_057_II_和为s的连续正数序列.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	target := 15
	fmt.Println(findContinuousSequence(target))
}

func findContinuousSequence(target int) (ans [][]int) {
	left, right := 1, 1
	sum := 0
	for left <= target>>1 {
		if sum < target {
			sum += right
			right++
		} else if sum > target {
			sum -= left
			left++
		} else {
			tmp := []int{}
			for i := left; i < right; i++ {
				tmp = append(tmp, i)
			}
			ans = append(ans, tmp)
			sum -= left
			left++
		}
	}
	return
}
