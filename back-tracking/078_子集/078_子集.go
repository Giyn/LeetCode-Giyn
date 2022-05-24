/*
-------------------------------------
# @Time    : 2021/12/4 15:49:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 078_子集.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}

func subsets(nums []int) (ans [][]int) {
	var path []int
	var backtrack func(index int)
	backtrack = func(index int) {
		ans = append(ans, append([]int{}, path...))
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
		return
	}
	backtrack(0)
	return
}
