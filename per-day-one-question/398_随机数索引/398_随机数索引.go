/*
-------------------------------------
# @Time    : 2022/4/26 21:24:57
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 398_随机数索引.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	solution := Constructor([]int{1, 2, 3, 3, 3})
	fmt.Println(solution.Pick(3))
	fmt.Println(solution.Pick(1))
	fmt.Println(solution.Pick(3))
}

type Solution struct {
	mp map[int][]int
}

func Constructor(nums []int) Solution {
	mp := map[int][]int{}
	for i, num := range nums {
		mp[num] = append(mp[num], i)
	}
	return Solution{mp}
}

func (s *Solution) Pick(target int) int {
	nums := s.mp[target]
	return nums[rand.Intn(len(nums))]
}
