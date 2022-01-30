/*
-------------------------------------
# @Time    : 2022/1/30 18:36:28
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 004_寻找两个正序数组的中位数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	n := len(nums1)
	if n%2 == 1 {
		return float64(nums1[n/2])
	} else {
		return float64(nums1[n/2]+nums1[n/2-1]) / 2
	}
}
