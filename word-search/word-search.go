package main

import "fmt"

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] && dfs(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i int, j int, count int, word string) bool {
	if count == len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || board[i][j] != word[count] {
		return false
	}

	temp := board[i][j]
	board[i][j] = ' '
	found := dfs(board, i+1, j, count+1, word) ||
		dfs(board, i-1, j, count+1, word) ||
		dfs(board, i, j+1, count+1, word) ||
		dfs(board, i, j-1, count+1, word)
	board[i][j] = temp
	return found
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}
