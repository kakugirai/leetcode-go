package main

import "fmt"

func countBattleships(board [][]byte) int {
	numBattleships := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' {
				numBattleships++
				sink(board, i, j)
			}
		}
	}
	return numBattleships
}

func sink(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || board[i][j] != 'X' {
		return
	}
	board[i][j] = '.'
	sink(board, i+1, j)
	sink(board, i-1, j)
	sink(board, i, j+1)
	sink(board, i, j-1)
}

func main() {
	board := [][]byte{
		{'X', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
	}
	fmt.Println(countBattleships(board))
}
