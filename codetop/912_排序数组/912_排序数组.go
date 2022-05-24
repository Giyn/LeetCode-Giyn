/*
-------------------------------------
# @Time    : 2022/4/2 18:20:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 912_排序数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{5, 1, 1, 2, 0, 0}
	fmt.Println(sortArray(nums))
}

func sortArray(nums []int) []int {
	//var quickSort func(s []int, left, right int)
	//quickSort = func(s []int, left, right int) {
	//	i, j := left, right
	//	pivot := s[(left+right)/2]
	//	for i <= j {
	//		for s[i] < pivot {
	//			i++
	//		}
	//		for s[j] > pivot {
	//			j--
	//		}
	//		if i <= j {
	//			s[i], s[j] = s[j], s[i]
	//			i++
	//			j--
	//		}
	//	}
	//	if j > left {
	//		quickSort(s, left, j)
	//	}
	//	if i < right {
	//		quickSort(s, i, right)
	//	}
	//}
	//quickSort(nums, 0, len(nums)-1)
	//return nums

	//var mergeSort func(left, right []int) (ans []int)
	//mergeSort = func(left, right []int) (ans []int) {
	//	for len(left) > 0 && len(right) > 0 {
	//		if left[0] <= right[0] {
	//			ans = append(ans, left[0])
	//			left = left[1:]
	//		} else {
	//			ans = append(ans, right[0])
	//			right = right[1:]
	//		}
	//	}
	//	for len(left) > 0 {
	//		ans = append(ans, left[0])
	//		left = left[1:]
	//	}
	//	for len(right) > 0 {
	//		ans = append(ans, right[0])
	//		right = right[1:]
	//	}
	//	return
	//}
	//if len(nums) < 2 {
	//	return nums
	//}
	//mid := len(nums) / 2
	//return mergeSort(sortArray(nums[:mid]), sortArray(nums[mid:]))

	var heapify func(i, n int)
	heapify = func(i, n int) {
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
			heapify(max, n)
		}
	}
	n := len(nums)
	// 建堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(i, n)
	}
	// 堆顶交换到尾部
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(0, i) // 继续调整
	}
	return nums
}
