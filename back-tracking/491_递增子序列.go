/*
-------------------------------------
# @Time    : 2021/12/6 16:31:14
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 491_递增子序列.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{4, 6, 7, 7}
	fmt.Println(findSubsequences(nums))
}

func findSubsequences(nums []int) (ans [][]int) {
	var path []int
	var backtrack func(int)
	backtrack = func(index int) {
		if len(path) >= 2 {
			ans = append(ans, append([]int{}, path...))
		}
		// 记录本层使用过的数
		//mp := map[int]bool{}
		record := [201]int{}
		for i := index; i < len(nums); i++ {
			if (len(path) > 0 && nums[i] < path[len(path)-1]) || record[nums[i]+100] == 1 {
				continue
			}
			//mp[nums[i]] = true
			record[nums[i]+100] = 1
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return
}
