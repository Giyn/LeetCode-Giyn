/*
-------------------------------------
# @Time    : 2021/10/12 11:53:38
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 844_比较含退格的字符串.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s, t := "ab#c", "ad#c"
	fmt.Println(backspaceCompare(s, t))
}

func backspaceCompare(s string, t string) bool {
	skipS, skipT := 0, 0
	i, j := len(s)-1, len(t)-1

	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}
