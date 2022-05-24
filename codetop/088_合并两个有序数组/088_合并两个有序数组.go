/*
-------------------------------------
# @Time    : 2022/3/30 22:05:52
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 088_合并两个有序数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for last := len(nums1) - 1; m > 0 || n > 0; last-- {
		if m == 0 {
			nums1[last] = nums2[n-1]
			n--
		} else if n == 0 {
			nums1[last] = nums1[m-1]
			m--
		} else if nums1[m-1] >= nums2[n-1] {
			nums1[last] = nums1[m-1]
			m--
		} else {
			nums1[last] = nums2[n-1]
			n--
		}
	}
}
