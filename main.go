package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	board := [9][9]int{}
	rows := [9][10]bool{}
	cols := [9][10]bool{}
	boxes := [9][10]bool{}

	for i := 0; i < 9; i++ {
		line := os.Args[i+1]
		if len(line) != 9 {
			fmt.Println("Error")
			return
		}
		for j, ch := range line {
			if ch == '.' {
				board[i][j] = 0
				continue
			}
			if ch < '1' || ch > '9' {
				fmt.Println("Error")
				return
			}
			d := int(ch - '0')
			boxIdx := (i/3)*3 + (j / 3)
			if rows[i][d] || cols[j][d] || boxes[boxIdx][d] {
				fmt.Println("Error")
				return
			}
			board[i][j] = d
			rows[i][d] = true
			cols[j][d] = true
			boxes[boxIdx][d] = true
		}
	}

	solutionsFound := 0
	var firstSolution [9][9]int

	var solve func(r, c int) bool
	solve = func(r, c int) bool {
		for i := r; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if i == r && j < c {
					continue
				}
				if board[i][j] == 0 {
					boxIdx := (i/3)*3 + (j / 3)
					for d := 1; d <= 9; d++ {
						if !rows[i][d] && !cols[j][d] && !boxes[boxIdx][d] {
							board[i][j] = d
							rows[i][d] = true
							cols[j][d] = true
							boxes[boxIdx][d] = true

							if solve(i, j+1) {
							}

							rows[i][d] = false
							cols[j][d] = false
							boxes[boxIdx][d] = false
							board[i][j] = 0

							if solutionsFound >= 2 {
								return true
							}
						}
					}
					return false
				}
			}
		}
		solutionsFound++
		if solutionsFound == 1 {
			firstSolution = board
		}
		return false
	}

	_ = solve(0, 0)

	if solutionsFound != 1 {
		fmt.Println("Error")
		return
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(firstSolution[i][j])
		}
		fmt.Println()
	}
}
