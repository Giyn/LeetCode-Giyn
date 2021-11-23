/*
-------------------------------------
# @Time    : 2021/11/23 14:27:12
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 859_亲密字符串.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "abcaa"
	goal := "abcbb"
	fmt.Println(buddyStrings(s, goal))
}

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		record := [26]int{}
		for _, ch := range s {
			record[ch-'a']++
			if record[ch-'a'] >= 2 {
				return true
			}
		}
		return false
	}
	// 只能有两个字符不同
	if s != goal {
		first, second := -1, -1
		for i := 0; i < len(s); i++ {
			if s[i] != goal[i] {
				if first == -1 {
					first = i
				} else if second == -1 {
					second = i
				} else {
					return false
				}
			}
		}
		return first != -1 && second != -1 && s[first] == goal[second] && s[second] == goal[first]
	}
	return true
}
