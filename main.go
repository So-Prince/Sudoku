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
	var b [9][9]int
	for r, s := range os.Args[1:] {
		if len(s) != 9 {
			fmt.Println("Error")
			return
		}
		for c, char := range s {
			if char == '.' {
				b[r][c] = 0
			} else if char >= '1' && char <= '9' {
				num := int(char - '0')
				if !isSafe(b, r, c, num) {
					fmt.Println("Error")
					return
				}
				b[r][c] = num
			} else {
				fmt.Println("Error")
				return
			}
		}
	}
	var count int
	var res [9][9]int
	solve(b, &count, &res)
	if count == 1 {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Print(res[i][j])
				if j < 8 {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Error")
	}
}
func isSafe(b [9][9]int, r, c, n int) bool {
	for i := 0; i < 9; i++ {
		if b[r][i] == n || b[i][c] == n {
			return false
		}
	}
	sr, sc := (r/3)*3, (c/3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[sr+i][sc+j] == n {
				return false
			}
		}
	}
	return true
}
func solve(b [9][9]int, count *int, res *[9][9]int) {
	if *count > 1 {
		return
	}
	r, c := -1, -1
	for i := 0; i < 81; i++ {
		if b[i/9][i%9] == 0 {
			r, c = i/9, i%9
			break
		}
	}
	if r == -1 {
		*count++
		*res = b
		return
	}
	for n := 1; n <= 9; n++ {
		if isSafe(b, r, c, n) {
			b[r][c] = n
			solve(b, count, res)
		}
	}
}
