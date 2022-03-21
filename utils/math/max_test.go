/*
-------------------------------------
# @Time    : 2022/3/13 13:20:16
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : max_test.go
# @Software: GoLand
-------------------------------------
*/

package math

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	fmt.Println(Max([]int{1, 2, 3, 4, 6}))
	fmt.Println(Max(1, 2, 3, 4, 6))
}
