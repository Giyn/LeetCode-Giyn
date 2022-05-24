/*
-------------------------------------
# @Time    : 2022/1/26 9:51:15
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2013_检测正方形.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

type DetectSquares struct {
	cnts map[int]map[int]int
}

func main() {
	detectSquares := Constructor()
	detectSquares.Add([]int{3, 10})
	detectSquares.Add([]int{11, 2})
	detectSquares.Add([]int{3, 2})
	fmt.Println(detectSquares.Count([]int{11, 10}))
	fmt.Println(detectSquares.Count([]int{14, 8}))
	detectSquares.Add([]int{11, 2})
	fmt.Println(detectSquares.Count([]int{11, 10}))
}

func Constructor() DetectSquares {
	return DetectSquares{map[int]map[int]int{}}
}

func (ds *DetectSquares) Add(point []int) {
	if ds.cnts[point[1]] == nil {
		ds.cnts[point[1]] = map[int]int{}
	}
	ds.cnts[point[1]][point[0]]++
}

func (ds *DetectSquares) Count(point []int) (ans int) {
	if ds.cnts[point[1]] != nil {
		for x, v := range ds.cnts[point[1]] {
			diff := x - point[0]
			if diff != 0 {
				if uy := point[1] + diff; ds.cnts[uy] != nil {
					ans += v * ds.cnts[uy][x] * ds.cnts[uy][point[0]]
				}
				if by := point[1] - diff; ds.cnts[by] != nil {
					ans += v * ds.cnts[by][x] * ds.cnts[by][point[0]]
				}
			}
		}
	}
	return
}
