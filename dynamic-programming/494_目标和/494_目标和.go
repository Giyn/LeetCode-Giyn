/*
-------------------------------------
# @Time    : 2022/1/26 16:52:25
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 494_目标和.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	fmt.Println(findTargetSumWays(nums, target))
}

func findTargetSumWays(nums []int, target int) int {
	var abs func(int) int
	abs = func(n int) int {
		y := n >> 63
		return (n ^ y) - y
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if (target+sum)%2 == 1 || abs(target) > sum {
		return 0
	}
	bagSize := (target + sum) / 2
	dp := make([]int, bagSize+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := bagSize; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bagSize]
}
