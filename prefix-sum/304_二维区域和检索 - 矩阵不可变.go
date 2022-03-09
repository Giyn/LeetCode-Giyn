/*
-------------------------------------
# @Time    : 2022/3/9 10:50:46
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 304_二维区域和检索 - 矩阵不可变.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	matrix := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}
	mat := Constructor304(matrix)
	fmt.Println(mat.SumRegion(2, 1, 4, 3))
	fmt.Println(mat.SumRegion(1, 1, 2, 2))
	fmt.Println(mat.SumRegion(1, 2, 2, 4))
}

type NumMatrix struct {
	sum [][]int
}

func Constructor304(matrix [][]int) NumMatrix {
	mat := NumMatrix{}
	m, n := len(matrix), len(matrix[0])
	mat.sum = make([][]int, m)
	for i := 0; i < m; i++ {
		mat.sum[i] = make([]int, n+1)
		for j := 0; j < n; j++ {
			mat.sum[i][j+1] += mat.sum[i][j] + matrix[i][j]
		}
	}
	return mat
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) (ans int) {
	for i := row1; i <= row2; i++ {
		ans += this.sum[i][col2+1] - this.sum[i][col1]
	}
	return
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
