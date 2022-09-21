/*
-------------------------------------
# @Time    : 2022/9/21 23:12:30
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 460_LFU 缓存.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	fmt.Println(lfu.Get(1))
	lfu.Put(3, 3)
	fmt.Println(lfu.Get(2))
	fmt.Println(lfu.Get(3))
	lfu.Put(4, 4)
	fmt.Println(lfu.Get(1))
	fmt.Println(lfu.Get(3))
	fmt.Println(lfu.Get(4))
}

type listNode struct {
	key, value, frequency int
	prev, next            *listNode
}

type doubleList struct {
	head *listNode
	tail *listNode
}

type LFUCache struct {
	recent   map[int]*listNode
	count    map[int]int
	cache    map[int]*listNode
	list     *doubleList
	capacity int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		recent:   make(map[int]*listNode, capacity),
		count:    make(map[int]int),
		cache:    make(map[int]*listNode, capacity),
		list:     newDoubleList(),
		capacity: capacity,
	}

}

func (this *LFUCache) Get(key int) int {
	if this.capacity == 0 {
		return -1
	}

	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	// key存在
	next := node.next
	if this.count[node.frequency+1] > 0 {
		// 存在其他使用次数n+1的缓存，指定缓存移动到所有使用次数为n+1的节点之前
		removeNode(node)
		addFirst(this.recent[node.frequency+1], node)
	} else if this.count[node.frequency] > 1 && this.recent[node.frequency] != node {
		// 不存在其他使用次数为n+1的缓存，但存在其他使用次数为n的缓存，且当前节点不是最近的节点
		// 将指定缓存移动到所有使用次数为n的节点之前
		removeNode(node)
		addFirst(this.recent[node.frequency], node)
	}
	// 更新recent
	this.recent[node.frequency+1] = node
	if this.count[node.frequency] <= 1 {
		// 不存在freq=n的节点，recent置空
		this.recent[node.frequency] = nil
	} else if this.recent[node.frequency] == node {
		// 存在，且recent=node，将recent后移
		this.recent[node.frequency] = next
	}
	// 更新使用次数对应的节点数
	this.count[node.frequency+1]++
	this.count[node.frequency]--
	// 更新缓存使用次数
	node.frequency++
	return node.value
}

func (this *LFUCache) Put(key, value int) {
	if this.capacity == 0 {
		return
	}
	node, ok := this.cache[key]
	if ok {
		this.Get(key)
		node.value = value
		return
	}
	// key不存在
	if len(this.cache) >= this.capacity {
		// cache 满了
		tailNode := this.list.tail.prev
		this.list.removeTail()
		if this.count[tailNode.frequency] <= 1 {
			this.recent[tailNode.frequency] = nil
		}
		this.count[tailNode.frequency]--
		delete(this.cache, tailNode.key)
	}

	newNode := &listNode{
		key:       key,
		value:     value,
		frequency: 1,
	}
	// insert new cache node
	if this.count[1] > 0 {
		addFirst(this.recent[1], newNode)
	} else {
		addFirst(this.list.tail, newNode)
	}
	// 更新recent， count， cache
	this.recent[1] = newNode
	this.count[1]++
	this.cache[key] = newNode
}

func newDoubleList() *doubleList {
	head := &listNode{}
	tail := &listNode{}
	head.next = tail
	tail.prev = head
	return &doubleList{
		head: head,
		tail: tail,
	}
}

func removeNode(node *listNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func addFirst(curNode, newNode *listNode) {
	prev := curNode.prev
	prev.next = newNode
	newNode.next = curNode
	curNode.prev = newNode
	newNode.prev = prev
}

func (this *doubleList) removeTail() {
	prev := this.tail.prev.prev
	prev.next = this.tail
	this.tail.prev = prev
}

func (this *doubleList) isEmpty() bool {
	return this.head.next == this.tail
}
