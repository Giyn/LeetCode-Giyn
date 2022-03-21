/*
-------------------------------------
# @Time    : 2022/3/21 16:38:21
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : bubble-sort_test.go
# @Software: GoLand
-------------------------------------
*/

package bubble_sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{6, 2, 7, 5, 9, 27, 12, 52, 57, 42}
	BubbleSort(nums)
	fmt.Println(nums)
}
