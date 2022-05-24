/*
-------------------------------------
# @Time    : 2022/1/27 15:48:46
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 377_组合总和 Ⅳ.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	target := 4
	fmt.Println(combinationSum(nums, target))
}

func combinationSum(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i >= nums[j] {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}
