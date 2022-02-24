/*
-------------------------------------
# @Time    : 2022/2/24 13:06:38
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 496_下一个更大元素 I.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums1 := []int{2, 4}
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
		if v, ok := mp[num]; ok {
			ans = append(ans, v)
		} else {
			ans = append(ans, -1)
		}
	}
	return
}
