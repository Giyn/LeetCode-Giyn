/*
-------------------------------------
# @Time    : 2022/2/24 9:56:26
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : max.go
# @Software: GoLand
-------------------------------------
*/

package utils

func Max(args ...int) (ans int) {
	ans = args[0]
	for _, arg := range args[1:] {
		if arg > ans {
			ans = arg
		}
	}
	return
}
