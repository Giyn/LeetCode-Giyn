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
	solution := Constructor384([]int{1, 2, 3})
	fmt.Println(solution.Shuffle())
	fmt.Println(solution.Reset())
	fmt.Println(solution.Shuffle())
}

type Solution struct {
	nums []int
}

func Constructor384(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	tmp := make([]int, len(this.nums))
	copy(tmp, this.nums)
	n := len(tmp)
	for i := 0; i < n; i++ {
		j := i + rand.Intn(n-i)
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	return tmp
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
