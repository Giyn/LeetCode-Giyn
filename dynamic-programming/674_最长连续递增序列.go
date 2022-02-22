/*
-------------------------------------
# @Time    : 2022/2/22 10:42:44
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 674_最长连续递增序列.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(nums))
}

func findLengthOfLCIS(nums []int) (ans int) {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	ans = dp[0]
	for i := 1; i < n; i++ {
		if nums[i-1] < nums[i] {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return
}
