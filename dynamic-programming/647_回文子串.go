/*
-------------------------------------
# @Time    : 2022/1/31 15:55:38
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 647_回文子串.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "abc"
	fmt.Println(countSubstrings(s))
}

func countSubstrings(s string) (ans int) {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			// 只有一个字符 || 上一层是回文串
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
				ans++
			}
		}
	}
	return

	// 中心扩展法
	//n := len(s)
	//for center := 0; center < 2*n-1; center++ {
	//	// 例如 aba: 0.(0, 0), 1.(0, 1), 2.(1, 1), 3.(1, 2), 4.(2, 2)
	//	left := center / 2
	//	right := left + center%2
	//	for left >= 0 && right < n && s[left] == s[right] {
	//		ans++
	//		left--
	//		right++
	//	}
	//}
	//return
}
