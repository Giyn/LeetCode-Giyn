/*
-------------------------------------
# @Time    : 2021/11/24 8:43:30
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 423_从英文中重建数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "fviefuro"
	fmt.Println(originalDigits(s))
}

func originalDigits(s string) string {
	mp := map[rune]int{}
	for _, ch := range s {
		mp[ch]++
	}

	cnt := [10]int{}
	cnt[0] = mp['z']
	cnt[2] = mp['w']
	cnt[4] = mp['u']
	cnt[6] = mp['x']
	cnt[8] = mp['g']

	cnt[3] = mp['h'] - cnt[8]
	cnt[5] = mp['f'] - cnt[4]
	cnt[7] = mp['s'] - cnt[6]

	cnt[1] = mp['o'] - cnt[0] - cnt[2] - cnt[4]

	cnt[9] = mp['i'] - cnt[5] - cnt[6] - cnt[8]

	var ans []byte
	for i, c := range cnt {
		ans = append(ans, bytes.Repeat([]byte{byte('0' + i)}, c)...)
	}
	return string(ans)
}
