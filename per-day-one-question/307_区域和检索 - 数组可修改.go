/*
-------------------------------------
# @Time    : 2022/4/4 0:03:46
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 307_区域和检索 - 数组可修改.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	na := Constructor307([]int{9, -8})
	na.Update(0, 3)
	fmt.Println(na.SumRange(1, 1))
	fmt.Println(na.SumRange(0, 1))
	na.Update(1, -3)
	fmt.Println(na.SumRange(0, 1))
}

type NumArray struct {
	nums []int
	tree []int
}

func Constructor307(nums []int) NumArray {
	tree := make([]int, len(nums)+1)
	na := NumArray{nums, tree}
	for i, num := range nums {
		na.add(i+1, num)
	}
	return na
}

func (this *NumArray) add(index, val int) {
	for ; index < len(this.tree); index += index & -index {
		this.tree[index] += val
	}
}

func (this *NumArray) prefixSum(index int) (sum int) {
	for ; index > 0; index &= index - 1 {
		sum += this.tree[index]
	}
	return
}

func (this *NumArray) Update(index, val int) {
	this.add(index+1, val-this.nums[index])
	this.nums[index] = val
}

func (this *NumArray) SumRange(left, right int) int {
	return this.prefixSum(right+1) - this.prefixSum(left)
}
