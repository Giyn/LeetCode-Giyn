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

func majorityElement(nums []int) []int {
	n := len(nums)
	m := make(map[int]int)
	var ans []int
	t := n / 3

	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if v > t {
			ans = append(ans, k)
		}
	}

	return ans
}
