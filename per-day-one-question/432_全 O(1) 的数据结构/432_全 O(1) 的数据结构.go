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
	ao := Constructor()
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

func Constructor() AllOne {
	return AllOne{list.New(), map[string]*list.Element{}}
}

func (ao *AllOne) Inc(key string) {
	if cur := ao.nodes[key]; cur != nil {
		curNode := cur.Value.(node)
		if nxt := cur.Next(); nxt == nil || nxt.Value.(node).count > curNode.count+1 {
			ao.nodes[key] = ao.InsertAfter(node{map[string]struct{}{key: {}}, curNode.count + 1}, cur)
		} else {
			nxt.Value.(node).keys[key] = struct{}{}
			ao.nodes[key] = nxt
		}
		delete(curNode.keys, key)
		if len(curNode.keys) == 0 {
			ao.Remove(cur)
		}
	} else { // key 不在链表中
		if ao.Front() == nil || ao.Front().Value.(node).count > 1 {
			ao.nodes[key] = ao.PushFront(node{map[string]struct{}{key: {}}, 1})
		} else {
			ao.Front().Value.(node).keys[key] = struct{}{}
			ao.nodes[key] = ao.Front()
		}
	}
}

func (ao *AllOne) Dec(key string) {
	cur := ao.nodes[key]
	curNode := cur.Value.(node)
	if curNode.count > 1 {
		if pre := cur.Prev(); pre == nil || pre.Value.(node).count < curNode.count-1 {
			ao.nodes[key] = ao.InsertBefore(node{map[string]struct{}{key: {}}, curNode.count - 1}, cur)
		} else {
			pre.Value.(node).keys[key] = struct{}{}
			ao.nodes[key] = pre
		}
	} else { // key 仅出现一次，将其移出 nodes
		delete(ao.nodes, key)
	}
	delete(curNode.keys, key)
	if len(curNode.keys) == 0 {
		ao.Remove(cur)
	}
}

func (ao *AllOne) GetMaxKey() string {
	if b := ao.Back(); b != nil {
		for key := range b.Value.(node).keys {
			return key
		}
	}
	return ""
}

func (ao *AllOne) GetMinKey() string {
	if f := ao.Front(); f != nil {
		for key := range f.Value.(node).keys {
			return key
		}
	}
	return ""
}
