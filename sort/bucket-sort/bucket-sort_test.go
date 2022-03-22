/*
-------------------------------------
# @Time    : 2022/3/22 12:20:22
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : bucket-sort_test.go
# @Software: GoLand
-------------------------------------
*/

package bucket_sort

import (
	"fmt"
	"testing"
)

func TestBucketSort(t *testing.T) {
	nums := []int{6, 2, 7, 5, 9, 27, 12, 52, 57, 42}
	BucketSort(nums)
	fmt.Println(nums)
}
