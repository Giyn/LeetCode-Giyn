/*
-------------------------------------
# @Time    : 2022/2/13 10:52:23
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1189_“气球” 的最大数量.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	text := "lloo"
	fmt.Println(maxNumberOfBalloons(text))
}

func maxNumberOfBalloons(text string) (ans int) {
	// 模拟
	min := func(values ...int) (ans int) {
		ans = values[0]
		for _, v := range values {
			if v < ans {
				ans = v
			}
		}
		return
	}
	b, a, l, o, n := 0, 0, 0, 0, 0
	for _, ch := range text {
		if ch == 'b' {
			b++
		} else if ch == 'a' {
			a++
		} else if ch == 'l' {
			l++
		} else if ch == 'o' {
			o++
		} else if ch == 'n' {
			n++
		}
	}
	return min(b, a, l/2, o/2, n)

	// 哈希
	//n := len(text)
	//mp := make(map[byte]int, n)
	//for i := 0; i < n; i++ {
	//	mp[text[i]]++
	//}
	//for {
	//	if mp['b'] < 1 || mp['a'] < 1 || mp['l'] < 2 || mp['o'] < 2 || mp['n'] < 1 {
	//		break
	//	} else {
	//		ans++
	//		mp['b']--
	//		mp['a']--
	//		mp['l'] -= 2
	//		mp['o'] -= 2
	//		mp['n']--
	//	}
	//}
	//return
}
