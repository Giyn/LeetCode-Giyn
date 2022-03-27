/*
-------------------------------------
# @Time    : 2022/3/27 15:33:15
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2028_找出缺失的观测数据.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	rolls := []int{6, 1, 5, 2}
	mean := 4
	n := 4
	fmt.Println(missingRolls(rolls, mean, n))
}

func missingRolls(rolls []int, mean int, n int) (ans []int) {
	remain := (len(rolls)+n)*mean - Sum(rolls)
	if remain < n || remain > n*6 {
		return
	}
	base := remain / n
	carry := remain % n
	for n > 0 {
		if carry > 0 {
			ans = append(ans, base+1)
			carry--
		} else {
			ans = append(ans, base)
		}
		n--
	}
	return
}
