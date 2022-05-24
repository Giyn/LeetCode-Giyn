/*
-------------------------------------
# @Time    : 2022/3/16 15:47:03
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_067_最大的异或.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 10, 5, 25, 2, 8}
	fmt.Println(findMaximumXOR(nums))
}

type Trie struct {
	left, right *Trie
}

func (t *Trie) Add(num int) {
	node := t
	for i := 30; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if node.left == nil {
				node.left = &Trie{}
			}
			node = node.left
		} else {
			if node.right == nil {
				node.right = &Trie{}
			}
			node = node.right
		}
	}
}

func (t *Trie) Check(num int) (x int) {
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
	root := &Trie{}
	for i := 1; i < len(nums); i++ {
		root.Add(nums[i-1])
		ans = max(ans, root.Check(nums[i]))
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
