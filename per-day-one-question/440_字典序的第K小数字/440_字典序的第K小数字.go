/*
-------------------------------------
# @Time    : 2022/3/23 0:03:24
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 440_字典序的第K小数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	n := 13
	k := 2
	fmt.Println(findKthNumber(n, k))
}

func findKthNumber(n int, k int) int {
	getSteps := func(cur, n int) (steps int) {
		first, last := cur, cur
		for first <= n {
			steps += min(last, n) - first + 1
			first *= 10
			last = last*10 + 9
		}
		return
	}
	cur := 1
	k--
	for k > 0 {
		steps := getSteps(cur, n)
		if steps <= k {
			k -= steps
			cur++
		} else {
			cur *= 10
			k--
		}
	}
	return cur
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
