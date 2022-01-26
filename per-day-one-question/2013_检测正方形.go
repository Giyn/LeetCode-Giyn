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
	detectSquares := Constructor2013()
	detectSquares.Add([]int{3, 10})
	detectSquares.Add([]int{11, 2})
	detectSquares.Add([]int{3, 2})
	fmt.Println(detectSquares.Count([]int{11, 10}))
	fmt.Println(detectSquares.Count([]int{14, 8}))
	detectSquares.Add([]int{11, 2})
	fmt.Println(detectSquares.Count([]int{11, 10}))
}

func Constructor2013() DetectSquares {
	return DetectSquares{map[int]map[int]int{}}
}

func (this *DetectSquares) Add(point []int) {
	if this.cnts[point[1]] == nil {
		this.cnts[point[1]] = map[int]int{}
	}
	this.cnts[point[1]][point[0]]++
}

func (this *DetectSquares) Count(point []int) (ans int) {
	if this.cnts[point[1]] != nil {
		for x, v := range this.cnts[point[1]] {
			diff := x - point[0]
			if diff != 0 {
				if uy := point[1] + diff; this.cnts[uy] != nil {
					ans += v * this.cnts[uy][x] * this.cnts[uy][point[0]]
				}
				if by := point[1] - diff; this.cnts[by] != nil {
					ans += v * this.cnts[by][x] * this.cnts[by][point[0]]
				}
			}
		}
	}
	return
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
