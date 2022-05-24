/*
-------------------------------------
# @Time    : 2022/1/25 16:09:18
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1049_最后一块石头的重量 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	stones := []int{31, 26, 33, 21, 40}
	fmt.Println(lastStoneWeightII(stones))
}

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	target := sum / 2
	dp := make([]int, target+1)
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return (sum - dp[target]) - dp[target]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
