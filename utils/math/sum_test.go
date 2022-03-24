/*
-------------------------------------
# @Time    : 2022/3/24 9:49:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : sum_test.go
# @Software: GoLand
-------------------------------------
*/

package math

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	fmt.Println(Sum([]int{1, 2, 3, 4, 6}))
	fmt.Println(Sum(1, 2, 3, 4, 6))
}
