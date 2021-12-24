/*
-------------------------------------
# @Time    : 2021/12/25 1:32:02
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1609_奇偶树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/list"
	"fmt"
	"math"
)

type TreeNode1609 struct {
	Val   int
	Left  *TreeNode1609
	Right *TreeNode1609
}

func main() {
	//root := &TreeNode1609{1, &TreeNode1609{10, &TreeNode1609{3, &TreeNode1609{12, nil, nil}, &TreeNode1609{8, nil, nil}}, nil}, &TreeNode1609{4, &TreeNode1609{7, &TreeNode1609{6, nil, nil}, nil}, &TreeNode1609{9, nil, &TreeNode1609{2, nil, nil}}}}
	root := &TreeNode1609{5, &TreeNode1609{4, &TreeNode1609{3, nil, nil}, &TreeNode1609{3, nil, nil}}, &TreeNode1609{2, &TreeNode1609{7, nil, nil}, nil}}
	fmt.Println(isEvenOddTree(root))
}

func isEvenOddTree(root *TreeNode1609) bool {
	queue := list.New()
	queue.PushBack(root)
	var even = true
	for queue.Len() > 0 {
		length := queue.Len()
		var prev int
		if even {
			prev = math.MinInt64
		} else {
			prev = math.MaxInt64
		}
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode1609)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if even && (node.Val%2 != 1 || node.Val <= prev) {
				return false
			} else if !even && (node.Val%2 != 0 || node.Val >= prev) {
				return false
			}
			prev = node.Val
		}
		even = !even
	}
	return true
}
