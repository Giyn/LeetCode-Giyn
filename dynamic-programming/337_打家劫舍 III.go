/*
-------------------------------------
# @Time    : 2022/1/28 21:42:11
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 337_打家劫舍 III.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

type TreeNode337 struct {
	Val   int
	Left  *TreeNode337
	Right *TreeNode337
}

func main() {
	root := &TreeNode337{3, &TreeNode337{2, nil, &TreeNode337{3, nil, nil}}, &TreeNode337{3, nil, &TreeNode337{1, nil, nil}}}
	fmt.Println(rob3(root))
}

func rob3(root *TreeNode337) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var robTree func(cur *TreeNode337) [2]int
	robTree = func(cur *TreeNode337) [2]int {
		if cur == nil {
			return [2]int{0, 0}
		}
		left := robTree(cur.Left)
		right := robTree(cur.Right)
		yesRob := cur.Val + left[0] + right[0]
		noRob := max(left[0], left[1]) + max(right[0], right[1])
		return [2]int{noRob, yesRob}
	}
	ans := robTree(root)
	return max(ans[0], ans[1])
}
