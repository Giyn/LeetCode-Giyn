/*
-------------------------------------
# @Time    : 2022/3/11 18:27:29
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : node.go
# @Software: GoLand
-------------------------------------
*/

package doubly_linked_list

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}
