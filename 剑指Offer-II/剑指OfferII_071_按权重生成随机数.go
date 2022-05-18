/*
-------------------------------------
# @Time    : 2022/5/18 16:14:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_071_按权重生成随机数.go
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
	solution := Constructor071([]int{1, 3})
	for i := 0; i < 5; i++ {
		fmt.Println(solution.PickIndex())
	}
}

type Solution071 struct {
	pre []int
}

func Constructor071(w []int) Solution071 {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution071{w}
}

func (this *Solution071) PickIndex() int {
	// pre[i]−w[i]+1≤x≤pre[i]
	x := rand.Intn(this.pre[len(this.pre)-1]) + 1
	return sort.SearchInts(this.pre, x)
}
