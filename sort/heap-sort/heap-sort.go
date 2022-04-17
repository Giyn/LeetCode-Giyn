/*
-------------------------------------
# @Time    : 2022/3/21 21:06:49
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : heap-sort.go
# @Software: GoLand
-------------------------------------
*/

package heap_sort

func HeapSort(nums []int) {
	n := len(nums)
	// 建堆
	for i := n/2 - 1; i >= 0; i-- {
		Heapify(nums, i, n)
	}
	// 堆顶交换到尾部
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		Heapify(nums, 0, i) // 继续调整
	}
}

func Heapify(nums []int, i, n int) {
	max := i
	left := i<<1 + 1
	right := i<<1 + 2
	if left < n && nums[left] > nums[max] {
		max = left
	}
	if right < n && nums[right] > nums[max] {
		max = right
	}
	// 把父节点换下去调整
	if max != i {
		nums[max], nums[i] = nums[i], nums[max]
		Heapify(nums, max, n)
	}
}
