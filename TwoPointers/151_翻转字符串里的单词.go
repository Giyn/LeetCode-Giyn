/*
-------------------------------------
# @Time    : 2021/11/4 21:45:50
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 151_翻转字符串里的单词.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "  Bob    Loves  Alice!   "
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	b := []byte(s)
	slowIndex, fastIndex := 0, 0

	// 去除前面的空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	// 去除单词间的多余空格
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex] == b[fastIndex-1] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	// 去除后面的空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}
	reverse151(&b, 0, len(b)-1)
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse151(&b, i, j-1)
		i = j
		i++
	}
	return string(b)
}

func reverse151(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}
