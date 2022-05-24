/*
-------------------------------------
# @Time    : 2021/11/10 9:28:36
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 239_滑动窗口最大值.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}

func maxSlidingWindow(nums []int, k int) (ans []int) {
	mQueue := NewMonotoneQueue()
	for i := 0; i < k; i++ {
		mQueue.Push(nums[i])
	}
	ans = append(ans, mQueue.Front())
	for i := k; i < len(nums); i++ {
		mQueue.Pop(nums[i-k])
		mQueue.Push(nums[i])
		ans = append(ans, mQueue.Front())
	}
	return
}

type MonotoneQueue struct {
	queue []int
}

func NewMonotoneQueue() *MonotoneQueue {
	return &MonotoneQueue{
		queue: make([]int, 0),
	}
}

func (mq *MonotoneQueue) Front() int {
	return mq.queue[0]
}

func (mq *MonotoneQueue) Back() int {
	return mq.queue[len(mq.queue)-1]
}

func (mq *MonotoneQueue) Empty() bool {
	return len(mq.queue) == 0
}

func (mq *MonotoneQueue) Pop(val int) {
	if !mq.Empty() && val == mq.Front() {
		mq.queue = mq.queue[1:]
	}
}

func (mq *MonotoneQueue) Push(val int) {
	for !mq.Empty() && val > mq.Back() {
		mq.queue = mq.queue[:len(mq.queue)-1]
	}
	mq.queue = append(mq.queue, val)
}
