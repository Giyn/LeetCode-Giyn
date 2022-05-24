/*
-------------------------------------
# @Time    : 2022/5/20 11:27:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_033_二叉搜索树的后序遍历序列.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	postorder := []int{1, 6, 3, 2, 5}
	fmt.Println(verifyPostorder(postorder))
}

func verifyPostorder(postorder []int) bool {
	n := len(postorder)
	if n <= 1 {
		return true
	}
	rootVal := postorder[n-1]
	split := 0
	for i := 0; i < n; i++ {
		if postorder[i] >= rootVal {
			split = i
			break
		}
	}
	for i := split; i < n; i++ {
		if postorder[i] < rootVal {
			return false
		}
	}
	return verifyPostorder(postorder[:split]) && verifyPostorder(postorder[split:n-1])
}
