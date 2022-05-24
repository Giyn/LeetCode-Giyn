/*
-------------------------------------
# @Time    : 2022/1/26 20:49:19
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 474_一和零.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m := 5
	n := 3
	fmt.Println(findMaxForm(strs, m, n))
}

func findMaxForm(strs []string, m int, n int) int {
	var max func(int, int) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		zeroNum, oneNum := 0, 0
		for _, ch := range str {
			if ch == '0' {
				zeroNum++
			} else {
				oneNum++
			}
		}
		for i := m; i >= zeroNum; i-- {
			for j := n; j >= oneNum; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeroNum][j-oneNum]+1)
			}
		}
	}
	return dp[m][n]
}
