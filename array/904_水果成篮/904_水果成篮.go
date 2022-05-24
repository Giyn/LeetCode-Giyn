/*
-------------------------------------
# @Time    : 2021/10/14 20:33:13
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 904_水果成篮.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	fruits := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	fmt.Println(totalFruit(fruits))
}

func totalFruit(fruits []int) int {
	hashMap := map[int]int{}
	i, j, ans, n := 0, 0, 0, len(fruits)
	for ; j < n; j++ {
		hashMap[fruits[j]]++
		for i <= j && len(hashMap) == 3 {
			hashMap[fruits[i]]--
			if hashMap[fruits[i]] == 0 {
				delete(hashMap, fruits[i])
			}
			i++
		}
		if ans < j-i+1 {
			ans = j - i + 1
		}
	}
	return ans
}
