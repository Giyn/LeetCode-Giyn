/*
-------------------------------------
# @Time    : 2022/3/21 20:25:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : merge-sort.go
# @Software: GoLand
-------------------------------------
*/

package merge_sort

func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	return Merge(MergeSort(nums[:mid]), MergeSort(nums[mid:]))
}

func Merge(left, right []int) (ans []int) {
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			ans = append(ans, left[0])
			left = left[1:]
		} else {
			ans = append(ans, right[0])
			right = right[1:]
		}
	}
	for len(left) > 0 {
		ans = append(ans, left[0])
		left = left[1:]
	}
	for len(right) > 0 {
		ans = append(ans, right[0])
		right = right[1:]
	}
	return
}
