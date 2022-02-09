/*
-------------------------------------
# @Time    : 2022/2/9 23:03:46
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2006_差的绝对值为 K 的数对数目.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 1}
	k := 1
	fmt.Println(countKDifference(nums, k))
}

func countKDifference(nums []int, k int) (ans int) {
	//abs := func(n int) int {
	//	y := n >> 63
	//	return (n ^ y) - y
	//}
	//for i := 0; i < len(nums); i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		if i < j && abs(nums[i]-nums[j]) == k {
	//			ans++
	//		}
	//	}
	//}
	//return

	// 哈希
	mp := make(map[int]int, len(nums))
	for _, num := range nums {
		mp[num]++
	}
	for _, num := range nums {
		if mp[num+k] > 0 {
			ans += mp[num+k]
		}
	}
	return
}
