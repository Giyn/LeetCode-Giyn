/*
-------------------------------------
# @Time    : 2022/5/17 4:42:57
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_039_数组中出现次数超过一半的数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) (ans int) {
	score := 0
	for _, num := range nums {
		if score == 0 {
			ans = num
			score++
		} else if num == ans {
			score++
		} else {
			score--
		}
	}
	return
}
