/*
-------------------------------------
# @Time    : 2021/11/9 0:07:14
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 488_祖玛游戏.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	board := "RRWWRRBBRR"
	hand := "WB"
	fmt.Println(findMinStep(board, hand))
}

func findMinStep(board string, hand string) int {
	// BFS
	queue := make([][]string, 0)
	queue = append(queue, []string{board, hand})
	depth := 0
	mp := map[string]bool{}
	for len(queue) > 0 {
		depth++
		k := len(queue)
		for k > 0 {
			cur := queue[0]
			queue = queue[1:]
			// 把 h 中的每个元素暴力插入 b 中的每个位置, 然后递归碰撞
			b, h := cur[0], cur[1]
			for i := 0; i < len(b); i++ {
				for j := 0; j < len(h); j++ {
					b2 := removeDup(b[:i] + string(h[j]) + b[i:])
					h2 := h[:j] + h[j+1:]
					if b2 == "" {
						return depth
					}
					// 去重剪枝
					key := fmt.Sprintf("%v_%v", b2, h2)
					if mp[key] {
						continue
					}
					mp[key] = true
					queue = append(queue, []string{b2, h2})
				}
			}
			k--
		}
	}
	return -1
}

func removeDup(s string) string {
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			if count >= 3 {
				return removeDup(s[:i-count] + s[i:])
			}
			count = 1
		}
	}
	if count >= 3 {
		return removeDup(s[:len(s)-count])
	}
	return s
}
