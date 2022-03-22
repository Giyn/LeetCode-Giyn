/*
-------------------------------------
# @Time    : 2022/3/22 10:55:33
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : bucket-sort.go
# @Software: GoLand
-------------------------------------
*/

package bucket_sort

import (
	. "LeetCodeGiyn/utils/math"
	"sort"
)

func BucketSort(nums []int) {
	bucketNum := len(nums) / 2 // 桶数
	maxVal := Max(nums)
	buckets := make([][]int, bucketNum)
	// 分配入桶
	index := 0
	for i := 0; i < bucketNum; i++ {
		// 分配桶 index = value * (n-1) /k
		index = nums[i] * (bucketNum - 1) / maxVal
		buckets[index] = append(buckets[index], nums[i])
	}
	// 桶内排序
	tmpPos := 0
	for i := 0; i < bucketNum; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			sort.Ints(buckets[i])
			copy(nums[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
}
