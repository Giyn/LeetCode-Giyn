/*
-------------------------------------
# @Time    : 2022/4/26 21:33:32
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 868_二进制间距.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
	"math"
)

func main() {
	n := 22
	fmt.Println(binaryGap(n))
}

func binaryGap(n int) (ans int) {
	last := math.MaxInt64
	for i := 0; n > 0; i++ {
		if n&1 == 1 {
			if last != -1 {
				ans = Max(ans, i-last)
			}
			last = i
		}
		n >>= 1
	}
	return
}
