/*
-------------------------------------
# @Time    : 2022/5/9 0:45:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 942_增减字符串匹配.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "IDID"
	fmt.Println(diStringMatch(s))
}

func diStringMatch(s string) []int {
	n := len(s)
	perm := make([]int, n+1)
	left, right := 0, n
	for i, ch := range s {
		if ch == 'I' {
			perm[i] = left
			left++
		} else {
			perm[i] = right
			right--
		}
	}
	perm[n] = left
	return perm
}
