/*
-------------------------------------
# @Time    : 2022/3/2 1:20:01
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_002_二进制加法.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
	"strconv"
)

func main() {
	a := "1010"
	b := "1011"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) (ans string) {
	carry := 0
	lenA, lenB := len(a), len(b)
	n := Max(lenA, lenB)

	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-1-i] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-1-i] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return
}
