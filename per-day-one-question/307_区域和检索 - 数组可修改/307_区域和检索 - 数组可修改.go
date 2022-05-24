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
	na := Constructor([]int{9, -8})
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

func Constructor(nums []int) NumArray {
	tree := make([]int, len(nums)+1)
	na := NumArray{nums, tree}
	for i, num := range nums {
		na.add(i+1, num)
	}
	return na
}

func (na *NumArray) add(index, val int) {
	for ; index < len(na.tree); index += index & -index {
		na.tree[index] += val
	}
}

func (na *NumArray) prefixSum(index int) (sum int) {
	for ; index > 0; index &= index - 1 {
		sum += na.tree[index]
	}
	return
}

func (na *NumArray) Update(index, val int) {
	na.add(index+1, val-na.nums[index])
	na.nums[index] = val
}

func (na *NumArray) SumRange(left, right int) int {
	return na.prefixSum(right+1) - na.prefixSum(left)
}
