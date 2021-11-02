/*
-------------------------------------
# @Time    : 2021/11/2 15:26:39
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 454_四数相加 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}
	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) (ans int) {
	mp := make(map[int]int)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			mp[v1+v2]++
		}
	}
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			ans += mp[-v3-v4]
		}
	}
	return
}
