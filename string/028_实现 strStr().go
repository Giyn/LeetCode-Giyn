/*
-------------------------------------
# @Time    : 2021/11/6 18:09:58
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 028_实现 strStr().go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
}

func getNext028(next []int, s string) {
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		// 回退
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}

func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	j := 0
	next := make([]int, n)
	getNext028(next, needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

//// 暴力
//func strStr(haystack string, needle string) int {
//	if needle == "" {
//		return 0
//	}
//	if len(needle) > len(haystack) {
//		return -1
//	}
//	for i := 0; i < len(haystack); i++ {
//		if haystack[i] == needle[0] && len(needle) == 1 {
//			return i
//		}
//		if haystack[i] == needle[0] {
//		outer:
//			for j, k := i+1, 1; j < len(haystack) && k < len(needle); j, k = j+1, k+1 {
//				if haystack[j] != needle[k] {
//					break outer
//				}
//				if k == len(needle)-1 && haystack[j] == needle[k] {
//					return i
//				}
//			}
//		}
//	}
//	return -1
//}
