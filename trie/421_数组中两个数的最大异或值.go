/*
-------------------------------------
# @Time    : 2022/3/16 18:56:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 421_数组中两个数的最大异或值.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	nums := []int{3, 10, 5, 25, 2, 8}
	fmt.Println(findMaximumXOR(nums))
}

type Trie421 struct {
	left, right *Trie421
}

func (t *Trie421) Add(num int) {
	node := t
	for i := 30; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if node.left == nil {
				node.left = &Trie421{}
			}
			node = node.left
		} else {
			if node.right == nil {
				node.right = &Trie421{}
			}
			node = node.right
		}
	}
}

func (t *Trie421) Check(num int) (x int) {
	node := t
	for i := 30; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			// a_i 的第 k 个二进制位为 0，应当往表示 1 的子节点 right 走
			if node.right != nil {
				node = node.right
				x = x*2 + 1
			} else {
				node = node.left
				x = x * 2
			}
		} else {
			// a_i 的第 k 个二进制位为 1，应当往表示 0 的子节点 left 走
			if node.left != nil {
				node = node.left
				x = x*2 + 1
			} else {
				node = node.right
				x = x * 2
			}
		}
	}
	return
}

func findMaximumXOR(nums []int) (ans int) {
	root := &Trie421{}
	for i := 1; i < len(nums); i++ {
		root.Add(nums[i-1])
		ans = Max(ans, root.Check(nums[i]))
	}
	return
}
