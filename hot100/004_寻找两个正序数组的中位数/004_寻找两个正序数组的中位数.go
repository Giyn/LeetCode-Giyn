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
)

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// o(m+n)
	//nums1 = append(nums1, nums2...)
	//sort.Ints(nums1)
	//n := len(nums1)
	//if n%2 == 1 {
	//	return float64(nums1[n/2])
	//} else {
	//	return float64(nums1[n/2]+nums1[n/2-1]) / 2
	//}

	// o(log(m+n))
	var getKth func(nums1 []int, start1 int, end1 int, nums2 []int, start2 int, end2 int, k int) int
	getKth = func(nums1 []int, start1 int, end1 int, nums2 []int, start2 int, end2 int, k int) int {
		len1 := end1 - start1 + 1
		len2 := end2 - start2 + 1
		// 让len1的长度小于len2 保证若有数组空了 就是len1
		if len1 > len2 {
			return getKth(nums2, start2, end2, nums1, start1, end1, k)
		}
		if len1 == 0 {
			return nums2[start2+k-1]
		}
		if k == 1 {
			return min(nums1[start1], nums2[start2])
		}
		i := start1 + min(len1, k/2) - 1
		j := start2 + min(len2, k/2) - 1
		if nums1[i] > nums2[j] {
			return getKth(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
		} else {
			return getKth(nums1, i+1, end1, nums2, start2, end2, k-(i-start1+1))
		}
	}
	n, m := len(nums1), len(nums2)
	left, right := (n+m+1)/2, (n+m+2)/2 // 合并偶数和奇数的情况
	return float64(getKth(nums1, 0, n-1, nums2, 0, m-1, left)+getKth(nums1, 0, n-1, nums2, 0, m-1, right)) / 2
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
