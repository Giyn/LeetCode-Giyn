/*
-------------------------------------
# @Time    : 2022/3/12 19:52:39
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_037_小行星碰撞.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	asteroids := []int{1, -2, -2, -2}
	fmt.Println(asteroidCollision(asteroids))
}

func asteroidCollision(asteroids []int) (ans []int) {
	for _, asteroid := range asteroids {
		if len(ans) == 0 || ans[len(ans)-1] < 0 || asteroid > 0 {
			ans = append(ans, asteroid)
		} else {
			// ans[len(ans)-1] > 0 && asteroid < 0
			for len(ans) > 0 && ans[len(ans)-1] > 0 {
				if -asteroid > ans[len(ans)-1] {
					ans = ans[:len(ans)-1]
					if len(ans) == 0 || (len(ans) > 0 && ans[len(ans)-1] < 0) {
						ans = append(ans, asteroid)
						break
					}
				} else if -asteroid == ans[len(ans)-1] {
					ans = ans[:len(ans)-1]
					break
				} else {
					break
				}
			}
		}
	}
	return
}
