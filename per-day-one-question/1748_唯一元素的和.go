/*
-------------------------------------
# @Time    : 2022/2/6 1:04:39
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1748_唯一元素的和.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sumOfUnique(nums))
}

func sumOfUnique(nums []int) (ans int) {
	mp := make(map[int]int, len(nums))
	for _, num := range nums {
		mp[num]++
		if mp[num] == 1 {
			ans += num
		} else if mp[num] == 2 {
			ans -= num
		}
	}
	return
}
