/*
-------------------------------------
# @Time    : 2021/12/22 19:42:55
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 763_划分字母区间.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	s := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(s))
}

func partitionLabels(s string) (ans []int) {
	lastPosition := [26]int{}
	for i, c := range s {
		lastPosition[c-'a'] = i
	}
	start, end := 0, 0
	for i, c := range s {
		if lastPosition[c-'a'] > end {
			end = lastPosition[c-'a']
		}
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}
	return
}
