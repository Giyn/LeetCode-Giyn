/*
-------------------------------------
# @Time    : 2022/3/21 16:33:40
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : bubble-sort.go
# @Software: GoLand
-------------------------------------
*/

package bubble_sort

func BubbleSort(nums []int) {
	// 排序的趟数
	for i := 0; i < len(nums)-1; i++ {
		var isChange bool
		// 当前需要排序的趟数
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isChange = true
			}
		}
		if !isChange {
			break
		}
	}
}
