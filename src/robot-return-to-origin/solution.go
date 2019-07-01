package main

import "fmt"

func judgeCircle(moves string) bool {
	m := map[byte]int{
		'L': 1,
		'R': -1,
		'U': 1,
		'D': -1,
	}

	row := 0
	col := 0
	for i := range moves {
		move := moves[i]
		if move == 'L' || move == 'R' {
			row += m[move]
		}
		if move == 'U' || move == 'D' {
			col += m[move]
		}
	}
	if row == 0 && col == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(judgeCircle("LDRRLRUULR"))
}
