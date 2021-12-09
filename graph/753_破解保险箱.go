/*
-------------------------------------
# @Time    : 2021/12/9 12:55:50
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 753_破解保险箱.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	n := 2
	k := 2
	fmt.Println(crackSafe(n, k))
}

func crackSafe(n int, k int) string {
	kn := int(math.Pow(float64(k), float64(n)))
	kn1 := int(math.Pow(float64(k), float64(n-1)))
	num := make([]int, kn1)
	for i := range num {
		num[i] = k - 1
	}
	s := make([]rune, kn+(n-1))
	for i := range s {
		s[i] = '0'
	}
	for i, node := n-1, 0; i < len(s); i++ {
		s[i] = rune(num[node]) + '0'
		num[node]--
		node = node*k - int(s[i-(n-1)]-'0')*kn1 + num[node] + 1
	}
	return string(s)
}
