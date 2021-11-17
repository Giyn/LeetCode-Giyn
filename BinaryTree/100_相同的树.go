/*
-------------------------------------
# @Time    : 2021/11/17 20:29:47
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 100_相同的树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	p := &TreeNode100{1, &TreeNode100{1, nil, nil}, nil}
	q := &TreeNode100{1, nil, &TreeNode100{1, nil, nil}}
	fmt.Println(isSameTree(p, q))
}

func isSameTree(p *TreeNode100, q *TreeNode100) bool {
	// 迭代
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	queue1 := list.New()
	queue2 := list.New()
	queue1.PushBack(p)
	queue2.PushBack(q)

	for queue1.Len() > 0 && queue2.Len() > 0 {
		node1, node2 := queue1.Remove(queue1.Front()).(*TreeNode100), queue2.Remove(queue2.Front()).(*TreeNode100)
		if node1.Val != node2.Val {
			return false
		}
		left1, right1 := node1.Left, node1.Right
		left2, right2 := node2.Left, node2.Right
		if (left1 == nil && left2 != nil) || (left1 != nil && left2 == nil) {
			return false
		}
		if (right1 == nil && right2 != nil) || (right1 != nil && right2 == nil) {
			return false
		}
		if left1 != nil {
			queue1.PushBack(left1)
		}
		if right1 != nil {
			queue1.PushBack(right1)
		}
		if left2 != nil {
			queue2.PushBack(left2)
		}
		if right2 != nil {
			queue2.PushBack(right2)
		}
	}
	return queue1.Len() == 0 && queue2.Len() == 0

	// 递归
	//if p == nil && q == nil {
	//	return true
	//}
	//if p == nil || q == nil {
	//	return false
	//}
	//if p.Val != q.Val {
	//	return false
	//}
	//return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

type TreeNode100 struct {
	Val   int
	Left  *TreeNode100
	Right *TreeNode100
}
