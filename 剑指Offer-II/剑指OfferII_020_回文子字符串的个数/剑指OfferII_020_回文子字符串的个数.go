/*
-------------------------------------
# @Time    : 2022/3/10 10:39:38
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_020_回文子字符串的个数.go
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
	// 中心扩展法
	n := len(s)
	for center := 0; center < 2*n-1; center++ {
		// 例如aba中心点: 0.(0, 0), 1.(0, 1), 2.(1, 1), 3.(1, 2), 4.(2, 2)
		left := center / 2
		right := left + center%2
		for left >= 0 && right < n && s[left] == s[right] {
			ans++
			left--
			right++
		}
	}
	return

	// dp
	//dp := make([][]bool, len(s))
	//for i := range dp {
	//	dp[i] = make([]bool, len(s))
	//}
	//for i := len(s) - 1; i >= 0; i-- {
	//	for j := i; j < len(s); j++ {
	//		if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
	//			dp[i][j] = true
	//			ans++
	//		}
	//	}
	//}
	//return
}
