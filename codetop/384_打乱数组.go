/*
-------------------------------------
# @Time    : 2022/5/3 13:44:12
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

type Solution384 struct {
	nums []int
}

func Constructor384(nums []int) Solution384 {
	return Solution384{nums}
}

func (this *Solution384) Reset() []int {
	return this.nums
}

func (this *Solution384) Shuffle() []int {
	n := len(this.nums)
	tmp := make([]int, n)
	copy(tmp, this.nums)
	for i := 0; i < n; i++ {
		j := i + rand.Intn(n-i) // 每个元素都等概率地交换到其他位置
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	return tmp
}
