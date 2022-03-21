/*
-------------------------------------
# @Time    : 2022/3/21 17:45:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : quick-sort.go
# @Software: GoLand
-------------------------------------
*/

package quick_sort

func QuickSort(nums []int, left, right int) {
	i, j := left, right
	pivot := nums[(left+right)/2] // 基准
	for i <= j {
		// 找到比基准大的数
		for nums[i] < pivot {
			i++
		}
		// 找到比基准小的数
		for nums[j] > pivot {
			j--
		}
		// 此时已经找到了比基准小(右)和大(左)的数,交换它们
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	if left < j {
		QuickSort(nums, left, j)
	}
	if right > i {
		QuickSort(nums, i, right)
	}
}
