/*
-------------------------------------
# @Time    : 2022/4/7 12:12:49
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_036_二叉搜索树与双向链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{4, 2, 5, 1, 3})
	ans := treeToDoublyList(root)
	tmp := ans
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Right
		if ans == tmp {
			break
		}
	}
}

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var dfs func(node *TreeNode) (head, tail *TreeNode)
	dfs = func(node *TreeNode) (head, tail *TreeNode) {
		if node == nil {
			return
		}
		// 递归,左子树
		lHead, lTail := dfs(node.Left)
		if lHead != nil {
			// 如果左子树不为空,头结点为左子树的头节点.并拼接当前节点到左子树的尾节点
			head = lHead
			lTail.Right = node
			node.Left = lTail
		} else {
			// 左子树为空,头结点为当前节点
			head = node
		}
		// 递归,右子树
		rHead, rTail := dfs(node.Right)
		if rTail != nil {
			// 如果右子树不为空,尾节点为右子树的尾节点.并拼接当前节点到右子树的头结点
			tail = rTail
			node.Right = rHead
			rHead.Left = node
		} else {
			// 右子树为空,尾节点为当前节点
			tail = node
		}
		return
	}
	head, tail := dfs(root)
	// 最后将返回的头尾节点拼接成环
	tail.Right = head
	head.Left = tail
	return head
}
