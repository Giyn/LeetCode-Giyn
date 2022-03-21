/*
-------------------------------------
# @Time    : 2022/3/21 16:06:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 128_最长连续序列.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) (ans int) {
	mp := map[int]bool{}
	for _, num := range nums {
		mp[num] = true
	}
	for v := range mp {
		if !mp[v-1] {
			tmp := v
			length := 1
			for mp[tmp+1] {
				tmp++
				length++
			}
			if length > ans {
				ans = length
			}
		}
	}
	return
}
