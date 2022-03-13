/*
-------------------------------------
# @Time    : 2022/3/12 17:00:14
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 146_LRU 缓存.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	lru := Constructor146(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))
}

type LRUCache struct {
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
	cap        int
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func Constructor146(capacity int) LRUCache {
	l := LRUCache{
		map[int]*DLinkedNode{},
		&DLinkedNode{0, 0, nil, nil},
		&DLinkedNode{0, 0, nil, nil},
		capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		if len(this.cache) >= this.cap {
			tail := this.removeTail()
			delete(this.cache, tail.key)
		}
		node := &DLinkedNode{key, value, nil, nil}
		this.cache[key] = node
		this.addToHead(node)
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() (node *DLinkedNode) {
	node = this.tail.prev
	this.removeNode(node)
	return node
}
