/*
-------------------------------------
# @Time    : 2022/6/28 5:31:20
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 007_整数反转.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	x := -123
	fmt.Println(reverse(x))
}

func reverse(x int) (ans int) {
	for x != 0 {
		if ans < math.MinInt32/10 || ans > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		ans = ans*10 + digit
	}
	return
}
