/*
-------------------------------------
# @Time    : 2022/2/24 10:37:20
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : abs.go
# @Software: GoLand
-------------------------------------
*/

package math

func Abs(n int) int {
	x := n >> 63
	return (n ^ x) - x
}
