/*
-------------------------------------
# @Time    : 2021/10/28 0:18:01
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 869_重新排序得到 2 的幂.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(reorderedPowerOf2(46))
}

func reorderedPowerOf2(n int) bool {
	length := len(strconv.Itoa(n))
	mp := make(map[[10]int]bool)
	for i := 0; ; i++ {
		val := int(math.Pow(2, float64(i)))
		if len(strconv.Itoa(val)) > length {
			break
		} else if len(strconv.Itoa(val)) < length {
			continue
		}
		mp[countBit(val)] = true
	}
	if _, ok := mp[countBit(n)]; ok {
		return true
	} else {
		return false
	}
}

func countBit(n int) (res [10]int) {
	for n > 0 {
		res[n%10]++
		n /= 10
	}
	return
}
