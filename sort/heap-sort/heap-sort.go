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
	length := len(nums)
	for i := 0; i < length; i++ {
		lastLen := length - i
		Heapify(nums, lastLen)
		if i < length {
			nums[0], nums[lastLen-1] = nums[lastLen-1], nums[0]
		}
	}
}

func Heapify(nums []int, length int) {
	if length <= 1 {
		return
	}
	depth := length/2 - 1 // 二叉树深度
	for i := depth; i >= 0; i-- {
		top := i
		left := 2*i + 1
		right := 2*i + 2
		if left <= length-1 && nums[left] > nums[top] {
			top = left
		}
		if right <= length-1 && nums[right] > nums[top] {
			top = right
		}
		if top != i {
			nums[i], nums[top] = nums[top], nums[i]
		}
	}
}
