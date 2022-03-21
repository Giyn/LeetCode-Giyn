/*
-------------------------------------
# @Time    : 2022/3/21 17:37:49
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : insertion-sort.go
# @Software: GoLand
-------------------------------------
*/

package insertion_sort

func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		prev := i - 1
		for prev >= 0 && nums[prev] > cur {
			nums[prev+1] = nums[prev] // 前移
			prev--
		}
		nums[prev+1] = cur // 插入
	}
}
