/*
-------------------------------------
# @Time    : 2021/12/6 16:04:18
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 090_子集 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}

func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	var path []int
	var backtrack func(index int)
	backtrack = func(index int) {
		ans = append(ans, append([]int{}, path...))
		for i := index; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return
}
