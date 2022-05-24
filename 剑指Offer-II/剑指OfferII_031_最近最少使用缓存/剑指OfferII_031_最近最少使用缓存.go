/*
-------------------------------------
# @Time    : 2022/3/12 14:47:49
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_031_最近最少使用缓存.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	lru := Constructor(2)
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

func Constructor(capacity int) LRUCache {
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

func (lru *LRUCache) Get(key int) int {
	if _, ok := lru.cache[key]; !ok {
		return -1
	}
	node := lru.cache[key]
	lru.moveToHead(node)
	return node.value
}

func (lru *LRUCache) Put(key int, value int) {
	if _, ok := lru.cache[key]; !ok {
		if len(lru.cache) >= lru.cap {
			tail := lru.removeTail()
			delete(lru.cache, tail.key)
		}
		node := &DLinkedNode{key, value, nil, nil}
		lru.cache[key] = node
		lru.addToHead(node)
	} else {
		node := lru.cache[key]
		node.value = value
		lru.moveToHead(node)
	}
}

func (lru *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache) moveToHead(node *DLinkedNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LRUCache) removeTail() (node *DLinkedNode) {
	node = lru.tail.prev
	lru.removeNode(node)
	return node
}
