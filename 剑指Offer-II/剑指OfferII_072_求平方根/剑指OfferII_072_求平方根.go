/*
-------------------------------------
# @Time    : 2022/3/16 20:01:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_072_求平方根.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	x := 25
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
