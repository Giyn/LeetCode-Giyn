/*
-------------------------------------
# @Time    : 2022/2/24 9:56:26
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : max.go
# @Software: GoLand
-------------------------------------
*/

package math

import "math"

func Max(args ...interface{}) (ans int) {
	ans = math.MinInt64
	for _, arg := range args {
		switch arg.(type) {
		case []int:
			for i := 0; i < len(arg.([]int)); i++ {
				if arg.([]int)[i] > ans {
					ans = arg.([]int)[i]
				}
			}
			return
		case int:
			if arg.(int) > ans {
				ans = arg.(int)
			}
		}
	}
	return
}
