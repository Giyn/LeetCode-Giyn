/*
-------------------------------------
# @Time    : 2022/3/21 16:46:55
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : selection-sort.go
# @Software: GoLand
-------------------------------------
*/

package selection_sort

func SelectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		var minIndex = i
		// 扫描到底
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}
