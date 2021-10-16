/*
-------------------------------------
# @Time    : 2021/10/15 19:34:15
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 038_外观数列.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 3
	fmt.Println(countAndSay(n))
}

func countAndSay(n int) string {
	prev := "1"
	
	for i := 2; i <= n; i++ {
		cur := &strings.Builder{}
		for j, start := 0, 0; j < len(prev); start = j {
			for j < len(prev) && prev[j] == prev[start] {
				j++
			}
			cur.WriteString(strconv.Itoa(j-start))
			cur.WriteByte(prev[start])
		}
		prev = cur.String()
	}
	return prev
}
