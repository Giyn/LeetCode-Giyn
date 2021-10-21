/*
-------------------------------------
# @Time    : 2021/10/22 0:02:47
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 229_求众数 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) (res []int) {
	cand1, count1 := nums[0], 0
	cand2, count2 := nums[0], 0

	for _, v := range nums {
		if cand1 == v {
			count1++
			continue
		}
		if cand2 == v {
			count2++
			continue
		}
		if count1 == 0 {
			cand1 = v
			count1++
			continue
		}
		if count2 == 0 {
			cand2 = v
			count2++
			continue
		}
		count1--
		count2--
	}

	count1 = 0
	count2 = 0

	for _, v := range nums {
		if cand1 == v {
			count1++
		} else if cand2 == v {
			count2++
		}
	}

	if count1 > len(nums)/3 {
		res = append(res, cand1)
	}
	if count2 > len(nums)/3 {
		res = append(res, cand2)
	}

	return
}

//func majorityElement(nums []int) []int {
//	n := len(nums)
//	m := make(map[int]int)
//	var ans []int
//	t := n / 3
//
//	for _, v := range nums {
//		m[v]++
//	}
//
//	for k, v := range m {
//		if v > t {
//			ans = append(ans, k)
//		}
//	}
//
//	return ans
//}
