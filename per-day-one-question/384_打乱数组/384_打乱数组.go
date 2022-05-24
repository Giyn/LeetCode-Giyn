/*
-------------------------------------
# @Time    : 2021/11/22 9:13:43
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 384_打乱数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	solution := Constructor([]int{1, 2, 3})
	fmt.Println(solution.Shuffle())
	fmt.Println(solution.Reset())
	fmt.Println(solution.Shuffle())
}

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (s *Solution) Reset() []int {
	return s.nums
}

func (s *Solution) Shuffle() []int {
	tmp := make([]int, len(s.nums))
	copy(tmp, s.nums)
	n := len(tmp)
	for i := 0; i < n; i++ {
		j := i + rand.Intn(n-i)
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	return tmp
}
