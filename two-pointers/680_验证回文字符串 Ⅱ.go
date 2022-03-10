/*
-------------------------------------
# @Time    : 2022/3/10 10:38:46
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 680_验证回文字符串 Ⅱ.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "abca"
	fmt.Println(validPalindrome(s))
}

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			flag1, flag2 := true, true
			for i, j := left+1, right; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}
			for i, j := left, right-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}
