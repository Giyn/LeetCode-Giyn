/*
-------------------------------------
# @Time    : 2022/1/18 16:40:18
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1220_统计元音字母序列的数目.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 100
	fmt.Println(countVowelPermutation(n))
}

func countVowelPermutation(n int) (ans int) {
	const mod int = 1e9 + 7
	dp := [5]int{1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		dp = [5]int{
			(dp[1] + dp[2] + dp[4]) % mod,
			(dp[0] + dp[2]) % mod,
			(dp[1] + dp[3]) % mod,
			dp[2],
			(dp[2] + dp[3]) % mod,
		}
	}
	for _, v := range dp {
		ans = (ans + v) % mod
	}
	return
}
