/*
-------------------------------------
# @Time    : 2022/3/22 10:24:47
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_006_从尾到头打印链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	head := NewListNode([]int{1, 3, 2})
	fmt.Println(reversePrint(head))
}

func reversePrint(head *ListNode) (ans []int) {
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	reverse(ans, 0, len(ans)-1)
	return

	// 递归
	//if head == nil {
	//	return nil
	//}
	//return append(reversePrint(head.Next), head.Val)
}

func reverse(s []int, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
