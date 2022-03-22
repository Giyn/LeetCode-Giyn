/*
-------------------------------------
# @Time    : 2022/3/22 10:56:47
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : counting-sort.go
# @Software: GoLand
-------------------------------------
*/

package counting_sort

import . "LeetCodeGiyn/utils/math"

func CountingSort(nums []int) {
	maxVal := Max(nums)
	minVal := Min(nums)
	bucket := make([]int, maxVal-minVal+1)
	for _, num := range nums {
		bucket[num-minVal]++
	}
	sortedIndex := 0
	for i := 0; i < maxVal-minVal+1; i++ {
		for bucket[i] > 0 {
			nums[sortedIndex] = i + minVal
			sortedIndex++
			bucket[i]--
		}
	}
}
