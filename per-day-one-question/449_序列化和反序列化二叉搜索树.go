/*
-------------------------------------
# @Time    : 2022/5/12 0:07:33
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 449_序列化和反序列化二叉搜索树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ser := Constructor449()
	root := NewTreeNode([]int{2, 1, 3})
	data := "2,1,3,#,#,#,#"
	fmt.Println(ser.serialize(root))
	PrintBinaryTree(ser.deserialize(data))
}

type Codec struct {
}

func Constructor449() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var res []string
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		if queue[0] != nil {
			res = append(res, strconv.Itoa(queue[0].Val))
			queue = append(queue, queue[0].Left, queue[0].Right)
		} else {
			res = append(res, "#")
		}
		queue = queue[1:]
	}
	return strings.Join(res, ",")
}

func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	var res = strings.Split(data, ",")
	var root = &TreeNode{}
	root.Val, _ = strconv.Atoi(res[0])
	var queue = []*TreeNode{root}
	res = res[1:]
	for len(queue) > 0 {
		if res[0] != "#" {
			l, _ := strconv.Atoi(res[0])
			queue[0].Left = &TreeNode{Val: l}
			queue = append(queue, queue[0].Left)
		}
		if res[1] != "#" {
			r, _ := strconv.Atoi(res[1])
			queue[0].Right = &TreeNode{Val: r}
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
		res = res[2:]
	}
	return root
}
