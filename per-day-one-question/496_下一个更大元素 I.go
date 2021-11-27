/*
-------------------------------------
# @Time    : 2021/10/26 0:00:43
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 496_下一个更大元素 I.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	//nums1 := []int{4, 1, 2}
	nums1 := []int{2, 4}
	//nums2 := []int{1, 3, 4, 2}
	nums2 := []int{1, 2, 3, 4}
	fmt.Println(nextGreaterElement(nums1, nums2))
}

func nextGreaterElement(nums1 []int, nums2 []int) (ans []int) {
	mp := make(map[int]int, len(nums2))
	var stack []int

	for _, num := range nums2 {
		for len(stack) > 0 && num > stack[len(stack)-1] {
			mp[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	for _, num := range nums1 {
		if v, ok := mp[num]; !ok {
			ans = append(ans, -1)
		} else {
			ans = append(ans, v)
		}
	}
	return ans
}

//func nextGreaterElement(nums1 []int, nums2 []int) (ans []int) {
//	for _, v := range nums1 {
//		pos := position(nums2, v)
//
//		for i := pos; i < len(nums2); i++ {
//			if nums2[i] > v {
//				ans = append(ans, nums2[i])
//				break
//			}
//			if i == len(nums2)-1 {
//				ans = append(ans, -1)
//			}
//		}
//	}
//	return
//}

//func position(slice []int, value int) int {
//	for p, v := range slice {
//		if v == value {
//			return p
//		}
//	}
//	return -1
//}
