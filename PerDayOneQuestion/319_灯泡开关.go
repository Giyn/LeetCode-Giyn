/*
-------------------------------------
# @Time    : 2021/11/15 10:46:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 319_灯泡开关.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	n := 15
	fmt.Println(bulbSwitch(n))
}

func bulbSwitch(n int) int {
	return int(math.Floor(math.Sqrt(float64(n))))
}
