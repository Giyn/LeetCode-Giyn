/*
-------------------------------------
# @Time    : 2022/1/24 20:44:48
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 416_分割等和子集.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 5}
	fmt.Println(canPartition(nums))
}

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	var max func(int, int) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	target := sum / 2
	dp := make([]int, target+1)
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}
