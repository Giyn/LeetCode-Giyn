/*
-------------------------------------
# @Time    : 2022/3/21 17:44:11
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : insertion-sort_test.go
# @Software: GoLand
-------------------------------------
*/

package insertion_sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	nums := []int{6, 2, 7, 5, 9, 27, 12, 52, 57, 42}
	InsertionSort(nums)
	fmt.Println(nums)
}
