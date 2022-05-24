/*
-------------------------------------
# @Time    : 2022/5/18 18:02:13
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 528_按权重随机选择.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	solution := Constructor([]int{1, 3})
	for i := 0; i < 5; i++ {
		fmt.Println(solution.PickIndex())
	}
}

type Solution struct {
	pre []int
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{w}
}

func (s *Solution) PickIndex() int {
	// pre[i]−w[i]+1≤x≤pre[i]
	x := rand.Intn(s.pre[len(s.pre)-1]) + 1
	return sort.SearchInts(s.pre, x)
}
