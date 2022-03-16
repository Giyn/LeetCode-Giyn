/*
-------------------------------------
# @Time    : 2022/3/16 0:38:04
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 432_全 O(1) 的数据结构.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	ao := Constructor432()
	ao.Inc("hello")
	ao.Inc("hello")
	fmt.Println(ao.GetMaxKey())
	fmt.Println(ao.GetMinKey())
	ao.Inc("leet")
	fmt.Println(ao.GetMaxKey())
	fmt.Println(ao.GetMinKey())
}

type node struct {
	keys  map[string]struct{}
	count int
}

type AllOne struct {
	*list.List
	nodes map[string]*list.Element
}

func Constructor432() AllOne {
	return AllOne{list.New(), map[string]*list.Element{}}
}

func (this *AllOne) Inc(key string) {
	if cur := this.nodes[key]; cur != nil {
		curNode := cur.Value.(node)
		if nxt := cur.Next(); nxt == nil || nxt.Value.(node).count > curNode.count+1 {
			this.nodes[key] = this.InsertAfter(node{map[string]struct{}{key: {}}, curNode.count + 1}, cur)
		} else {
			nxt.Value.(node).keys[key] = struct{}{}
			this.nodes[key] = nxt
		}
		delete(curNode.keys, key)
		if len(curNode.keys) == 0 {
			this.Remove(cur)
		}
	} else { // key 不在链表中
		if this.Front() == nil || this.Front().Value.(node).count > 1 {
			this.nodes[key] = this.PushFront(node{map[string]struct{}{key: {}}, 1})
		} else {
			this.Front().Value.(node).keys[key] = struct{}{}
			this.nodes[key] = this.Front()
		}
	}
}

func (this *AllOne) Dec(key string) {
	cur := this.nodes[key]
	curNode := cur.Value.(node)
	if curNode.count > 1 {
		if pre := cur.Prev(); pre == nil || pre.Value.(node).count < curNode.count-1 {
			this.nodes[key] = this.InsertBefore(node{map[string]struct{}{key: {}}, curNode.count - 1}, cur)
		} else {
			pre.Value.(node).keys[key] = struct{}{}
			this.nodes[key] = pre
		}
	} else { // key 仅出现一次，将其移出 nodes
		delete(this.nodes, key)
	}
	delete(curNode.keys, key)
	if len(curNode.keys) == 0 {
		this.Remove(cur)
	}
}

func (this *AllOne) GetMaxKey() string {
	if b := this.Back(); b != nil {
		for key := range b.Value.(node).keys {
			return key
		}
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	if f := this.Front(); f != nil {
		for key := range f.Value.(node).keys {
			return key
		}
	}
	return ""
}
