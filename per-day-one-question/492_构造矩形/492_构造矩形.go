/*
-------------------------------------
# @Time    : 2021/10/24 20:24:45
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 492_构造矩形.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(constructRectangle(19))
}

func constructRectangle(area int) []int {
	tmp := int(math.Sqrt(float64(area)))
	for {
		if area%tmp == 0 {
			return []int{area / tmp, tmp}
		}
		tmp--
	}
}
