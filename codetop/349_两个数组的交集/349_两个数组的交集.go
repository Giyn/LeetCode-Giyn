/*
-------------------------------------
# @Time    : 2021/11/1 22:15:11
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 349_两个数组的交集.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) (ans []int) {
	mp := make(map[int]bool, len(nums1))
	for _, num := range nums1 {
		mp[num] = true
	}
	for _, num := range nums2 {
		if mp[num] {
			ans = append(ans, num)
			mp[num] = false
		}
	}
	return
}
