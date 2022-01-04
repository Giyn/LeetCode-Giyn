/*
-------------------------------------
# @Time    : 2021/12/27 1:37:29
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 825_适龄的朋友.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []int{20, 30, 100, 110, 120}
	fmt.Println(numFriendRequests(ages))
}

func numFriendRequests(ages []int) (ans int) {
	sort.Ints(ages)
	left, right := 0, 0
	for _, age := range ages {
		if age < 15 {
			continue
		}
		for ages[left]*2 <= age+14 {
			left++
		}
		for right+1 < len(ages) && ages[right+1] <= age {
			right++
		}
		ans += right - left
	}
	return
}
