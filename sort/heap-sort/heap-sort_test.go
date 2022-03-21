/*
-------------------------------------
# @Time    : 2022/3/21 21:16:46
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : heap-sort_test.go
# @Software: GoLand
-------------------------------------
*/

package heap_sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	nums := []int{6, 2, 7, 5, 9, 27, 12, 52, 57, 42}
	HeapSort(nums)
	fmt.Println(nums)
}
