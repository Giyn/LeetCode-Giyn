/*
-------------------------------------
# @Time    : 2021/11/25 9:07:15
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 458_可怜的小猪.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	buckets := 1000
	minutesToDie := 15
	minutesToTest := 60
	fmt.Println(poorPigs(buckets, minutesToDie, minutesToTest))
}

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	return int(math.Ceil(math.Log2(float64(buckets)) / math.Log2(float64(minutesToTest/minutesToDie+1))))
}
