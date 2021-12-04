/*
-------------------------------------
# @Time    : 2021/12/4 9:31:42
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 131_分割回文串.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "aab"
	fmt.Println(partition(s))
}

func partition(s string) (ans [][]string) {
	var path []string
	// 判断回文串
	var isPalindrome func(str string, begin, end int) bool
	isPalindrome = func(str string, begin, end int) bool {
		for begin < end {
			if str[begin] == str[end] {
				begin++
				end--
			} else {
				return false
			}
		}
		return true
	}

	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(s) {
			ans = append(ans, append([]string{}, path...))
		}
		for i := index; i < len(s); i++ {
			if isPalindrome(s, index, i) {
				path = append(path, s[index:i+1])
			} else {
				continue
			}
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return
}
