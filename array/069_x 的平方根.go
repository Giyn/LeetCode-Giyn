/*
-------------------------------------
# @Time    : 2022/3/16 20:17:31
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 069_x 的平方根.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	x := 8
	fmt.Println(mySqrt(x))
}

func mySqrt(x int) (ans int) {
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return
}
