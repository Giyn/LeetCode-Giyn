/*
-------------------------------------
# @Time    : 2022/5/10 1:46:27
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1728_猫和老鼠 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	grid := []string{".M...", "..#..", "#..#.", "C#.#.", "...#F"}
	catJump := 3
	mouseJump := 1
	fmt.Println(canMouseWin(grid, catJump, mouseJump))
}

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	const (
		MouseTurn = 0
		CatTurn   = 1
		UNKNOWN   = 0
		MouseWin  = 1
		CatWin    = 2
		MaxMoves  = 1000
	)
	var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	rows, cols := len(grid), len(grid[0])
	getPos := func(row, col int) int { return row*cols + col }
	var startMouse, startCat, food int
	for i, row := range grid {
		for j, ch := range row {
			if ch == 'M' {
				startMouse = getPos(i, j)
			} else if ch == 'C' {
				startCat = getPos(i, j)
			} else if ch == 'F' {
				food = getPos(i, j)
			}
		}
	}

	// 计算每个状态的度
	total := rows * cols
	degrees := [64][64][2]int{}
	for mouse := 0; mouse < total; mouse++ {
		mouseRow := mouse / cols
		mouseCol := mouse % cols
		if grid[mouseRow][mouseCol] == '#' {
			continue
		}
		for cat := 0; cat < total; cat++ {
			catRow := cat / cols
			catCol := cat % cols
			if grid[catRow][catCol] == '#' {
				continue
			}
			degrees[mouse][cat][MouseTurn]++
			degrees[mouse][cat][CatTurn]++
			for _, dir := range dirs {
				for row, col, jump := mouseRow+dir.x, mouseCol+dir.y, 1; row >= 0 && row < rows && col >= 0 && col < cols && grid[row][col] != '#' && jump <= mouseJump; jump++ {
					nextMouse := getPos(row, col)
					nextCat := getPos(catRow, catCol)
					degrees[nextMouse][nextCat][MouseTurn]++
					row += dir.x
					col += dir.y
				}
				for row, col, jump := catRow+dir.x, catCol+dir.y, 1; row >= 0 && row < rows && col >= 0 && col < cols && grid[row][col] != '#' && jump <= catJump; jump++ {
					nextMouse := getPos(mouseRow, mouseCol)
					nextCat := getPos(row, col)
					degrees[nextMouse][nextCat][CatTurn]++
					row += dir.x
					col += dir.y
				}
			}
		}
	}

	results := [64][64][2][2]int{}
	type state struct{ mouse, cat, turn int }
	var q []state

	// 猫和老鼠在同一个单元格，猫获胜
	for pos := 0; pos < total; pos++ {
		row := pos / cols
		col := pos % cols
		if grid[row][col] == '#' {
			continue
		}
		results[pos][pos][MouseTurn][0] = CatWin
		results[pos][pos][MouseTurn][1] = 0
		results[pos][pos][CatTurn][0] = CatWin
		results[pos][pos][CatTurn][1] = 0
		q = append(q, state{pos, pos, MouseTurn}, state{pos, pos, CatTurn})
	}

	// 猫和食物在同一个单元格，猫获胜
	for mouse := 0; mouse < total; mouse++ {
		mouseRow := mouse / cols
		mouseCol := mouse % cols
		if grid[mouseRow][mouseCol] == '#' || mouse == food {
			continue
		}
		results[mouse][food][MouseTurn][0] = CatWin
		results[mouse][food][MouseTurn][1] = 0
		results[mouse][food][CatTurn][0] = CatWin
		results[mouse][food][CatTurn][1] = 0
		q = append(q, state{mouse, food, MouseTurn}, state{mouse, food, CatTurn})
	}

	// 老鼠和食物在同一个单元格且猫和食物不在同一个单元格，老鼠获胜
	for cat := 0; cat < total; cat++ {
		catRow := cat / cols
		catCol := cat % cols
		if grid[catRow][catCol] == '#' || cat == food {
			continue
		}
		results[food][cat][MouseTurn][0] = MouseWin
		results[food][cat][MouseTurn][1] = 0
		results[food][cat][CatTurn][0] = MouseWin
		results[food][cat][CatTurn][1] = 0
		q = append(q, state{food, cat, MouseTurn}, state{food, cat, CatTurn})
	}

	getPrevStates := func(mouse, cat, turn int) []state {
		mouseRow := mouse / cols
		mouseCol := mouse % cols
		catRow := cat / cols
		catCol := cat % cols
		prevTurn := MouseTurn
		if turn == MouseTurn {
			prevTurn = CatTurn
		}
		maxJump, startRow, startCol := catJump, catRow, catCol
		if prevTurn == MouseTurn {
			maxJump, startRow, startCol = mouseJump, mouseRow, mouseCol
		}
		prevStates := []state{{mouse, cat, prevTurn}}
		for _, dir := range dirs {
			for i, j, jump := startRow+dir.x, startCol+dir.y, 1; i >= 0 && i < rows && j >= 0 && j < cols && grid[i][j] != '#' && jump <= maxJump; jump++ {
				prevMouseRow := mouseRow
				prevMouseCol := mouseCol
				prevCatRow := i
				prevCatCol := j
				if prevTurn == MouseTurn {
					prevMouseRow = i
					prevMouseCol = j
					prevCatRow = catRow
					prevCatCol = catCol
				}
				prevMouse := getPos(prevMouseRow, prevMouseCol)
				prevCat := getPos(prevCatRow, prevCatCol)
				prevStates = append(prevStates, state{prevMouse, prevCat, prevTurn})
				i += dir.x
				j += dir.y
			}
		}
		return prevStates
	}

	// 拓扑排序
	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		mouse, cat, turn := s.mouse, s.cat, s.turn
		result := results[mouse][cat][turn][0]
		moves := results[mouse][cat][turn][1]
		for _, s := range getPrevStates(mouse, cat, turn) {
			prevMouse, prevCat, prevTurn := s.mouse, s.cat, s.turn
			if results[prevMouse][prevCat][prevTurn][0] == UNKNOWN {
				canWin := result == MouseWin && prevTurn == MouseTurn || result == CatWin && prevTurn == CatTurn
				if canWin {
					results[prevMouse][prevCat][prevTurn][0] = result
					results[prevMouse][prevCat][prevTurn][1] = moves + 1
					q = append(q, state{prevMouse, prevCat, prevTurn})
				} else {
					degrees[prevMouse][prevCat][prevTurn]--
					if degrees[prevMouse][prevCat][prevTurn] == 0 {
						loseResult := MouseWin
						if prevTurn == MouseTurn {
							loseResult = CatWin
						}
						results[prevMouse][prevCat][prevTurn][0] = loseResult
						results[prevMouse][prevCat][prevTurn][1] = moves + 1
						q = append(q, state{prevMouse, prevCat, prevTurn})
					}
				}
			}
		}
	}
	return results[startMouse][startCat][MouseTurn][0] == MouseWin && results[startMouse][startCat][MouseTurn][1] <= MaxMoves
}
