/*
-------------------------------------
# @Time    : 2022/1/20 16:12:32
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2029_石子游戏 IX.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	stones := []int{5, 1, 2, 4, 3}
	fmt.Println(stoneGameIX(stones))
}

func stoneGameIX(stones []int) bool {
	cnt0, cnt1, cnt2 := 0, 0, 0
	for _, val := range stones {
		val %= 3
		if val == 0 {
			cnt0++
		} else if val == 1 {
			cnt1++
		} else {
			cnt2++
		}
	}
	if cnt0%2 == 0 {
		return cnt1 >= 1 && cnt2 >= 1
	}
	return cnt1-cnt2 > 2 || cnt2-cnt1 > 2
}
