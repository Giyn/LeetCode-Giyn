/*
-------------------------------------
# @Time    : 2021/12/1 0:53:17
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1446_连续字符.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "abbcccddddeeeeedcba"
	fmt.Println(maxPower(s))
}

func maxPower(s string) (ans int) {
	ans = 1
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
			if count > ans {
				ans = count
			}
		} else {
			count = 1
		}
	}
	return
}
