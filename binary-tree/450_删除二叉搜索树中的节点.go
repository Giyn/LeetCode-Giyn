/*
-------------------------------------
# @Time    : 2021/11/26 15:56:30
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 450_删除二叉搜索树中的节点.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	root := &TreeNode450{5, &TreeNode450{3, &TreeNode450{2, nil, nil}, &TreeNode450{4, nil, nil}}, &TreeNode450{6, nil, &TreeNode450{7, nil, nil}}}
	key := 3
	res := deleteNode(root, key)
	fmt.Println(res.Left.Val)
}

func deleteNode(root *TreeNode450, key int) *TreeNode450 {
	// 递归
	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		}
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

type TreeNode450 struct {
	Val   int
	Left  *TreeNode450
	Right *TreeNode450
}
