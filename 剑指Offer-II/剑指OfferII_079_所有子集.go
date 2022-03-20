/*
-------------------------------------
# @Time    : 2022/3/19 17:40:58
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_079_所有子集.go
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
	var dfs func(index int)
	dfs = func(index int) {
		ans = append(ans, append([]int{}, path...))
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
		return
	}
	dfs(0)
	return
}
