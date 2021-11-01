/*
-------------------------------------
# @Time    : 2021/11/1 22:35:15
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 350_两个数组的交集 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersect(nums1, nums2))
}

func intersect(nums1 []int, nums2 []int) (ans []int) {
	mp := make(map[int]int)
	for _, num := range nums1 {
		mp[num]++
	}
	for _, num := range nums2 {
		if mp[num] != 0 {
			ans = append(ans, num)
			mp[num]--
		}
	}
	return
}
