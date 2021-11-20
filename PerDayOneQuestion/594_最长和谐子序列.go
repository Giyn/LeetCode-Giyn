/*
-------------------------------------
# @Time    : 2021/11/20 9:05:56
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 594_最长和谐子序列.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	fmt.Println(findLHS(nums))
}

func findLHS(nums []int) (ans int) {
	// 滑窗
	sort.Ints(nums)
	for i, j := 0, 0; j < len(nums); j++ {
		for i < j && nums[j]-nums[i] > 1 {
			i++
		}
		if nums[j]-nums[i] == 1 {
			if ans < j-i+1 {
				ans = j - i + 1
			}
		}
	}
	return

	// 哈希
	//sort.Ints(nums)
	//mp := make(map[int]int)
	//for _, num := range nums {
	//	mp[num]++
	//}
	//for k, v := range mp {
	//	if mp[k+1] != 0 && mp[k+1]+v > ans {
	//		ans = mp[k+1] + v
	//	}
	//}
	//return
}
