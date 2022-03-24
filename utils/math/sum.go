/*
-------------------------------------
# @Time    : 2022/3/24 9:48:45
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : sum.go
# @Software: GoLand
-------------------------------------
*/

package math

func Sum(args ...interface{}) (ans int) {
	for _, arg := range args {
		switch arg.(type) {
		case []int:
			for i := 0; i < len(arg.([]int)); i++ {
				ans += arg.([]int)[i]
			}
			return
		case int:
			ans += arg.(int)
		}
	}
	return
}
