/*
-------------------------------------
# @Time    : 2022/1/8 21:26:04
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 089_格雷编码.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 2
	fmt.Println(grayCode(n))
}

func grayCode(n int) []int {
	ans := make([]int, 1<<n)
	for i := range ans {
		ans[i] = i>>1 ^ i
	}
	return ans
}
