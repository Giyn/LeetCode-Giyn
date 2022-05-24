/*
-------------------------------------
# @Time    : 2022/1/22 11:35:44
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1332_删除回文子序列.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "baabb"
	fmt.Println(removePalindromeSub(s))
}

func removePalindromeSub(s string) int {
	for i, n := 0, len(s)-1; i < n/2; i++ {
		if s[i] != s[n-i] {
			return 2
		}
	}
	return 1
}
