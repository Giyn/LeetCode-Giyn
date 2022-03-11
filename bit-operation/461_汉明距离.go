/*
-------------------------------------
# @Time    : 2022/3/11 19:30:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 461_汉明距离.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	x, y := 1, 4
	fmt.Println(hammingDistance(x, y))
}

func hammingDistance(x int, y int) (ans int) {
	for s := x ^ y; s > 0; s >>= 1 {
		ans += s & 1
	}
	return
}
