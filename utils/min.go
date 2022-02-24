/*
-------------------------------------
# @Time    : 2022/2/24 9:55:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : min.go
# @Software: GoLand
-------------------------------------
*/

package utils

func Min(args ...int) (ans int) {
	ans = args[0]
	for _, arg := range args[1:] {
		if arg < ans {
			ans = arg
		}
	}
	return
}
