/*
-------------------------------------
# @Time    : 2022/4/25 21:09:19
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
	l, r := 0, x
	for l <= r {
		mid := l + (r-l)>>1
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return
}
