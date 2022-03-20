/*
-------------------------------------
# @Time    : 2022/2/24 9:55:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : min.go
# @Software: GoLand
-------------------------------------
*/

package math

import "math"

func Min(args ...interface{}) (ans int) {
	ans = math.MaxInt64
	for _, arg := range args {
		switch arg.(type) {
		case []int:
			for i := 0; i < len(arg.([]int)); i++ {
				if arg.([]int)[i] < ans {
					ans = arg.([]int)[i]
				}
			}
			return
		case int:
			if arg.(int) < ans {
				ans = arg.(int)
			}
		}
	}
	return
}
