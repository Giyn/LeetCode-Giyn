/*
-------------------------------------
# @Time    : 2021/11/17 9:04:54
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 318_最大单词长度乘积.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	words := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	fmt.Println(maxProduct(words))
}

func maxProduct(words []string) (ans int) {
	masks := map[int]int{}
	for _, word := range words {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		if len(word) > masks[mask] {
			masks[mask] = len(word)
		}
	}
	for x, lenX := range masks {
		for y, lenY := range masks {
			if x&y == 0 && lenX*lenY > ans {
				ans = lenX * lenY
			}
		}
	}
	return
}
