/*
-------------------------------------
# @Time    : 2022/3/21 20:40:45
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : merge-sort_test.go
# @Software: GoLand
-------------------------------------
*/

package merge_sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{6, 2, 7, 5, 9, 27, 12, 52, 57, 42}
	fmt.Println(MergeSort(nums))
}
