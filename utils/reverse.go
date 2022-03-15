/*
-------------------------------------
# @Time    : 2022/3/15 22:12:10
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : reverse.go
# @Software: GoLand
-------------------------------------
*/

package utils

func Reverse(s interface{}, left, right int) {
	switch s.(type) {
	case []int:
		for left < right {
			s.([]int)[left], s.([]int)[right] = s.([]int)[right], s.([]int)[left]
			left++
			right--
		}
	case []byte:
		for left < right {
			s.([]byte)[left], s.([]byte)[right] = s.([]byte)[right], s.([]byte)[left]
			left++
			right--
		}
	}
}
