/*
-------------------------------------
# @Time    : 2022/5/2 17:58:25
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 556_下一个更大元素 III.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	n := 12
	fmt.Println(nextGreaterElement(n))
}

func nextGreaterElement(n int) int {
	s := []byte(strconv.Itoa(n))
	i := len(s) - 2
	j := -1
	for i >= 0 {
		if s[i+1] > s[i] {
			j = i + 1
			break
		}
		i--
	}
	if j == -1 {
		return -1
	}
	for k := len(s) - 1; k >= j; k-- {
		if s[k] > s[i] {
			s[i], s[k] = s[k], s[i]
			break
		}
	}
	l, r := j, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	ans, _ := strconv.Atoi(string(s))
	if ans > math.MaxInt32 {
		return -1
	}
	return ans
}
